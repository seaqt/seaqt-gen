package main

import (
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

func ProcessLibraries(clangBin, outDir, extraLibsDir string) {
	qtVersion := qmakeQuery("QT_VERSION")
	major := toMajor(qtVersion)
	majorMinor := toMajorMinor(qtVersion)
	CacheSubdir = "qt" + majorMinor

	if major == "5" {
		buildQt5(clangBin, outDir, extraLibsDir, qtVersion)
	} else if major == "6" {
		buildQt6(clangBin, outDir, extraLibsDir, qtVersion)
	} else {
		panic("unsupported qt version: " + major)
	}
}

func buildQt5(clangBin, outDir, extraLibsDir, qtVersion string) {
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
	for _, lib := range modules {
		generate(
			"Qt"+lib, "Qt5"+lib, qtHeaders, qtVersion,
			[]string{filepath.Join(qtHeaders, "Qt"+lib)},
			clangBin, "", seaqtDir, nimSeaqtDir,
		)
	}
}

func buildQt6(clangBin, outDir, extraLibsDir, qtVersion string) {
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

		"RemoteObjects",

		"StateMachine",

		"Designer", "UiPlugin", "UiTools",

		"Qml",

		"Quick", "WebEngineQuick",
	}

	for _, lib := range modules {
		generate(
			"Qt"+lib, "Qt6"+lib, qtHeaders, qtVersion,
			[]string{filepath.Join(qtHeaders, "Qt"+lib)},
			clangBin, "--std=c++17", seaqtDir, nimSeaqtDir,
		)
	}
}
