package main

import (
	"os"
	"path/filepath"
)

func ProcessLibraries(clangBin, outDir, extraLibsDir string) {
	flushKnownTypes()
	InsertTypedefs(false)

	{
		version := "5.15"
		qtDir := filepath.Join(outDir, "seaqt-"+version)
		seaqtDir := filepath.Join(qtDir, "seaqt")
		nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+version, "seaqt")
		nimbleDir := filepath.Join(outDir, "nim-seaqt-"+version)

		os.RemoveAll(seaqtDir)
		os.RemoveAll(nimSeaqtDir)

		writeNimbleFile(nimbleDir, version)

		// TODO more modules, ideally via command-line support instead of this file
		// Must be processed in topological dependency order - loosely based on miqt
		// but without external modules
		libs5 := []string{
			// https://doc.qt.io/qt-5/qtmodules.html#qt-essentials
			"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",
			"Sql",

			// https://doc.qt.io/qt-5/qtmodules.html#qt-add-ons
			"PrintSupport", "Script", "Svg", "WebChannel",
			"WebEngineCore", "WebEngineWidgets",

			"Pdf", "PdfWidgets", "Positioning",

			"Qml", "Quick",

			// TODO Not sure where these are from :)
			"WebKit", "WebKitWidgets"}

		for _, lib := range libs5 {
			generate(
				"Qt"+lib, "Qt5"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt5/Qt" + lib},
				clangBin, "", seaqtDir, nimSeaqtDir,
			)
		}
	}

	flushKnownTypes()
	InsertTypedefs(true)

	// Qt 6
	{
		version := "6.4"
		qtDir := filepath.Join(outDir, "seaqt-"+version)
		seaqtDir := filepath.Join(qtDir, "seaqt")
		nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+version, "seaqt")
		nimbleDir := filepath.Join(outDir, "nim-seaqt-"+version)

		os.RemoveAll(seaqtDir)
		os.RemoveAll(nimSeaqtDir)

		writeNimbleFile(nimbleDir, version)

		libs6 := []string{
			"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",
			"Sql",

			"PrintSupport", "Svg", "SvgWidgets", "SpatialAudio", "WebChannel",
			"WebEngineCore", "WebEngineQuick", "WebEngineWidgets",

			"Pdf", "PdfWidgets", "Positioning",

			"Qml", "Quick",
		}

		for _, lib := range libs6 {
			generate(
				"Qt"+lib, "Qt6"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt6/Qt" + lib},
				clangBin, "--std=c++17", seaqtDir, nimSeaqtDir,
			)
		}
	}
}
