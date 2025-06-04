package main

import (
	"encoding/json"
	"flag"
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
		log.Printf("pkg-config(%q): %v", pcName, string(err.(*exec.ExitError).Stderr))
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

func generate(qtModuleName, pcName string, srcDirs []string, clangBin, cflagsExtras, outDir string) {

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

	atr := astTransformRedundant{
		preserve: make(map[string]*CppEnum),
	}

	//
	// PASS 1 (Parse headers and generate IL)
	//

	processHeaders := parseHeaders(includeFiles, clangBin, cflags)

	for _, parsed := range processHeaders {
		// AST transforms on our IL
		astTransformChildClasses(parsed)        // must be first
		astTransformApplyQuirks(pcName, parsed) // must be before optional/overload expansion
		astTransformOptional(parsed)
		astTransformOverloads(parsed)
		astTransformConstructorOrder(parsed)
		atr.Process(parsed)

		// Update global state tracker (AFTER astTransformChildClasses)
		addKnownTypes(qtModuleName, parsed)
	}

	//
	// PASS 2
	//

	cppFiles := make([]string, 0, len(processHeaders))

	for _, parsed := range processHeaders {

		log.Printf("Processing %q...", parsed.Filename)

		// More AST transforms on our IL
		astTransformTypedefs(parsed)
		astTransformBlocklist(parsed) // Must happen after typedef transformation

		{
			// Save the IL file for debug inspection
			file, err := os.Create(parsedPath(parsed.Filename))
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

		bindingCppSrc, err := emitBindingCpp(parsed, filepath.Base(parsed.Filename))
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".cpp", []byte(bindingCppSrc), 0644)
		if err != nil {
			panic(err)
		}
		cppFiles = append(cppFiles, genUnitName(parsed.Filename)+".cpp")

		bindingHSrc, err := emitBindingHeader(parsed, filepath.Base(parsed.Filename), qtModuleName)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".h", []byte(bindingHSrc), 0644)
		if err != nil {
			panic(err)
		}

		// Done

	}

	{
		amalgamation := emitAmalgamation(cppFiles)

		err := os.WriteFile(filepath.Join(outDir, qtModuleName+".cpp"), []byte(amalgamation), 0644)
		if err != nil {
			panic(err)
		}
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
