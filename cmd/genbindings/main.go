package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

const (
	MaxClangSubprocessCount = 16
	BaseModule              = "github.com/mappu/miqt"
)

func cacheFileRoot(inputHeader string) string {
	return filepath.Join("cachedir", strings.Replace(inputHeader, `/`, `__`, -1))
}

func parsedPath(inputHeader string) string {
	return cacheFileRoot(inputHeader) + ".ours.json"
}

func findHeadersInDir(srcDir string) []string {
	content, err := os.ReadDir(srcDir)
	if err != nil {
		panic(err)
	}

	var ret []string

	for _, includeFile := range content {
		if includeFile.IsDir() {
			continue
		}
		if !strings.HasSuffix(includeFile.Name(), `.h`) {
			continue
		}
		fullPath := filepath.Join(srcDir, includeFile.Name())
		if !AllowHeader(fullPath) {
			continue
		}
		ret = append(ret, fullPath)
	}

	return ret
}

func pkgConfig(pcName, what string) string {
	stdout, err := exec.Command(`pkg-config`, what, pcName).Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(stdout))
}
func genUnitName(header string) string {
	return "gen_" + strings.TrimSuffix(filepath.Base(header), `.h`)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parseHeaders(includeFiles []string, clangBin string, cflags []string) []*CppParsedHeader {
	result := make([]*CppParsedHeader, len(includeFiles))

	// Run clang / parsing in parallel but not too parallel
	var wg sync.WaitGroup
	ch := make(chan struct{}, min(runtime.NumCPU(), MaxClangSubprocessCount))

	for i, includeFile := range includeFiles {
		ch <- struct{}{}
		wg.Add(1)

		go func(i int, includeFile string) {
			defer func() {
				wg.Done()
				<-ch
			}()

			result[i] = &CppParsedHeader{Filename: includeFile}
			ast := getFilteredAst(includeFile, clangBin, cflags)
			// Convert it to our intermediate format
			parseHeader(ast, "", result[i])
		}(i, includeFile)
	}
	wg.Wait()
	return result
}

func generate(qtModuleName, pcName string, srcDirs []string, clangBin, cflagsExtras, outDir, nimOutdir string) {

	var includeFiles []string
	for _, srcDir := range srcDirs {
		if strings.HasSuffix(srcDir, `.h`) {
			includeFiles = append(includeFiles, srcDir) // single .h
		} else {
			includeFiles = append(includeFiles, findHeadersInDir(srcDir)...)
		}
	}

	log.Printf("Found %d header files to process.", len(includeFiles))

	cflags := append(strings.Fields(cflagsExtras), strings.Fields(pkgConfig(pcName, "--cflags"))...)
	outDir = filepath.Join(outDir, qtModuleName)
	os.MkdirAll(outDir, 0755)

	nimBaseDir := nimOutdir
	nimOutdir = filepath.Join(nimOutdir, qtModuleName)
	os.MkdirAll(nimOutdir, 0755)

	atr := astTransformRedundant{
		preserve: make(map[string]*CppEnum),
	}

	//
	// PASS 1 (Parse headers and generate IL)
	//

	processHeaders := parseHeaders(includeFiles, clangBin, cflags)

	for _, parsed := range processHeaders {
		// AST transforms on our IL
		astTransformChildClasses(parsed) // must be first
		astTransformOptional(parsed)
		astTransformOverloads(parsed)
		astTransformConstructorOrder(parsed)
		atr.Process(parsed)

		// Hack to add some overloads
		if strings.Contains(parsed.Filename, "qvariant.h") {
			for i, c := range parsed.Classes {
				if c.ClassName == "QVariant" {
					m := CppMethod{
						MethodName: "fromValue",
						ReturnType: CppParameter{ParameterType: "QVariant"},
						Parameters: []CppParameter{CppParameter{ParameterName: "value", ParameterType: "QObject", Const: true, Pointer: true}},
						IsStatic:   true,
					}
					parsed.Classes[i].Methods = append(parsed.Classes[i].Methods, m)
				}
			}
		}
		// Update global state tracker (AFTER astTransformChildClasses)
		addKnownTypes(qtModuleName, parsed)
	}

	//
	// PASS 2
	//
	{
		libsSrc := emitNimPkg(qtModuleName, pcName)
		os.WriteFile(filepath.Join(nimOutdir, nimModulePkgName(qtModuleName)+".nim"), []byte(libsSrc), 0644)
	}

	for _, parsed := range processHeaders {

		log.Printf("Processing %q...", parsed.Filename)

		// More AST transforms on our IL
		astTransformTypedefs(parsed)
		astTransformBlocklist(parsed) // Must happen after typedef transformation

		{
			// Save the IL file for debug inspection
			file, err := os.OpenFile(parsedPath(parsed.Filename), os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			enc := json.NewEncoder(file)
			enc.SetIndent("", "\t")
			enc.Encode(parsed)
		}

		// Breakout if there is nothing bindable
		if parsed.Empty() {
			log.Printf("Nothing in this header was bindable.")
			continue
		}

		// Emit 3 code files from the intermediate format
		outputName := filepath.Join(outDir, genUnitName(parsed.Filename))

		// For packages where we scan multiple directories, it's possible that
		// there are filename collisions (e.g. Qt 6 has QtWidgets/qaction.h include
		// QtGui/qaction.h as a compatibility measure).
		// If the path exists, disambiguate it
		var counter = 0
		for {
			testName := outputName
			if counter > 0 {
				testName += fmt.Sprintf(".%d", counter)
			}

			if _, err := os.Stat(testName + ".go"); err != nil && os.IsNotExist(err) {
				outputName = testName // Safe
				break
			}

			counter++
		}

		nimSrc, nim64Src, err := emitNim(parsed, filepath.Base(parsed.Filename), qtModuleName, pcName)
		if err != nil {
			log.Printf("Error in %s: %s", parsed.Filename, err)
			continue

			// panic(err)
		}

		nimOutputName := filepath.Join(nimOutdir, genUnitName(parsed.Filename))

		err = os.WriteFile(nimOutputName+".nim", []byte(nimSrc), 0644)
		if err != nil {
			panic(err)
		}

		if len(nim64Src) > 0 {
			err = os.WriteFile(nimOutputName+"_types.nim", []byte(nim64Src), 0644)
			if err != nil {
				panic(err)
			}
		}

		for _, c := range parsed.Classes {
			if strings.Contains(c.ClassName, "::") {
				// Inner classes are covered by their main declaration
				continue
			}

			os.WriteFile(
				filepath.Join(nimBaseDir, strings.ToLower(c.ClassName)+".nim"),
				[]byte("import ./"+qtModuleName+"/"+genUnitName(parsed.Filename)+"\nexport "+genUnitName(parsed.Filename)), 0644)
		}
		bindingCppSrc, err := emitBindingCpp(parsed, filepath.Base(parsed.Filename))
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".cpp", []byte(bindingCppSrc), 0644)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(nimOutputName+".cpp", []byte(bindingCppSrc), 0644)
		if err != nil {
			panic(err)
		}

		bindingHSrc, err := emitBindingHeader(parsed, filepath.Base(parsed.Filename), qtModuleName)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".h", []byte(bindingHSrc), 0644)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(nimOutputName+".h", []byte(bindingHSrc), 0644)
		if err != nil {
			panic(err)
		}

		// Done

	}

	log.Printf("Processing %d file(s) completed", len(includeFiles))
}

func main() {
	// data/time flags make logs hard to compare across runs
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	clang := flag.String("clang", "clang", "Custom path to clang")
	outDir := flag.String("outdir", "../../gen", "Output directory for generated gen_** files")
	extraLibsDir := flag.String("extralibs", "/usr/local/src/", "Base directory to find extra library checkouts")

	flag.Parse()

	ProcessLibraries(*clang, *outDir, *extraLibsDir)
}
