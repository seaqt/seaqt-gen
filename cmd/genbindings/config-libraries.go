package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func toMajor(qtVersion string) string {
	parts := strings.SplitN(qtVersion, ".", 2)
	return parts[0]
}
func toMajorMinor(qtVersion string) string {
	parts := strings.SplitN(qtVersion, ".", 3)
	return parts[0] + "." + parts[1]
}

// selectModules narrows modules to the comma-separated names in filter, keeping the
// order of modules. An empty filter selects every module.
func selectModules(modules []string, filter string) []string {
	if filter == "" {
		return modules
	}
	want := map[string]bool{}
	for _, m := range strings.Split(filter, ",") {
		if m = strings.TrimSpace(m); m != "" {
			want[m] = true
		}
	}
	var ret []string
	for _, m := range modules {
		if want[m] {
			ret = append(ret, m)
			delete(want, m)
		}
	}
	for m := range want {
		log.Panicf("-modules: unknown module %q", m)
	}
	log.Printf("-modules: generating %v", ret)
	return ret
}

func ProcessLibraries(clangBin, outDir, extraLibsDir, moduleFilter string) {
	qtVersion := qmakeQuery("QT_VERSION")
	major := toMajor(qtVersion)
	majorMinor := toMajorMinor(qtVersion)
	CacheSubdir = "qt" + majorMinor

	if major == "5" {
		buildQt5(clangBin, outDir, extraLibsDir, qtVersion, moduleFilter)
	} else if major == "6" {
		buildQt6(clangBin, outDir, extraLibsDir, qtVersion, moduleFilter)
	} else {
		panic("unsupported qt version: " + major)
	}
}

func buildQt5(clangBin, outDir, extraLibsDir, qtVersion, moduleFilter string) {
	flushKnownTypes()
	InsertTypedefs(false)

	qtHeaders := qmakeQuery("QT_INSTALL_HEADERS")
	majorMinor := toMajorMinor(qtVersion)
	qtDir := filepath.Join(outDir, "seaqt-"+majorMinor)
	seaqtDir := filepath.Join(qtDir, "seaqt")
	nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+majorMinor, "seaqt")
	nimbleDir := filepath.Join(outDir, "nim-seaqt-"+majorMinor)

	os.RemoveAll(seaqtDir)
	os.RemoveAll(nimSeaqtDir)

	writeNimbleFile(nimbleDir, qtVersion)

	modules := []string{
		// https://doc.qt.io/qt-5/qtmodules.html#qt-essentials
		"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",
		"Sql",

		// https://doc.qt.io/qt-5/qtmodules.html#qt-add-ons
		"PrintSupport", "Script", "Svg", "WebChannel", "WebSockets",
		"WebEngineCore", "WebEngineWidgets",

		// OpenGL is deprecated

		"Pdf", "PdfWidgets", "Positioning",

		"Qml", "Quick",
		// TODO Not sure where these are from :)
		"WebKit", "WebKitWidgets",

		// Tools
		// TODO(build fails) "Designer", "UiPlugin", "UiTools",
	}
	for _, lib := range selectModules(modules, moduleFilter) {
		generate(
			"Qt"+lib, "Qt5"+lib, qtHeaders, qtVersion,
			[]string{filepath.Join(qtHeaders, "Qt"+lib)},
			clangBin, "", seaqtDir, nimSeaqtDir,
		)
	}
}

func buildQt6(clangBin, outDir, extraLibsDir, qtVersion, moduleFilter string) {
	flushKnownTypes()
	InsertTypedefs(true)

	qtHeaders := qmakeQuery("QT_INSTALL_HEADERS")
	majorMinor := toMajorMinor(qtVersion)
	qtDir := filepath.Join(outDir, "seaqt-"+majorMinor)
	seaqtDir := filepath.Join(qtDir, "seaqt")
	nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+majorMinor, "seaqt")
	nimbleDir := filepath.Join(outDir, "nim-seaqt-"+majorMinor)

	os.RemoveAll(seaqtDir)
	os.RemoveAll(nimSeaqtDir)

	writeNimbleFile(nimbleDir, qtVersion)

	modules := []string{
		"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",
		"Sql",

		"PrintSupport", "Svg", "SvgWidgets", "SpatialAudio", "WebChannel", "WebSockets",
		"WebEngineCore", "WebEngineWidgets",

		"Pdf", "PdfWidgets", "Positioning",

		"StateMachine",

		"Designer", "UiPlugin", "UiTools",

		"Qml",

		"Quick", "WebEngineQuick",
	}

	for _, lib := range selectModules(modules, moduleFilter) {
		generate(
			"Qt"+lib, "Qt6"+lib, qtHeaders, qtVersion,
			[]string{filepath.Join(qtHeaders, "Qt"+lib)},
			clangBin, "--std=c++17", seaqtDir, nimSeaqtDir,
		)
	}
}
