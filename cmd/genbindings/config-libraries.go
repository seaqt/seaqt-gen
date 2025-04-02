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
		os.RemoveAll(seaqtDir)

		// TODO more modules
		// Must be processed in topological dependency order
		libs5 := []string{
			// https://doc.qt.io/qt-5/qtmodules.html#qt-essentials
			"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",

			// https://doc.qt.io/qt-5/qtmodules.html#qt-add-ons
			"PrintSupport", "Script", "Svg", "WebChannel",
			"WebEngineCore",

			// TODO Not sure where these are from :)
			"WebKit", "WebKitWidgets"}

		for _, lib := range libs5 {
			generate(
				"Qt"+lib, "Qt5"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt5/Qt" + lib},
				clangBin, "", seaqtDir,
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
		os.RemoveAll(seaqtDir)

		libs6 := []string{
			"Core", "Gui", "Widgets", "Network", "Multimedia", "MultimediaWidgets",

			"PrintSupport", "Svg", "SvgWidgets", "SpatialAudio", "WebChannel",
			"WebEngineCore", "WebEngineWidgets",

			"Qml",
		}

		for _, lib := range libs6 {
			generate(
				"Qt"+lib, "Qt6"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt6/Qt" + lib},
				clangBin, "--std=c++17", seaqtDir,
			)
		}
	}
}
