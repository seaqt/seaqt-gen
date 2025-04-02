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

			"Qml", "Quick",

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

	// TODO figure out a way to deal with the special cases
	// // Depends on QtCore/Gui/Widgets, QPrintSupport
	// generate(
	// 	"qt-restricted-extras/qscintilla",
	// 	[]string{
	// 		"/usr/include/x86_64-linux-gnu/qt5/Qsci",
	// 	},
	// 	AllowAllHeaders,
	// 	clangBin,
	// 	pkgConfigCflags("Qt5PrintSupport"),
	// 	outDir,
	// 	ClangMatchSameHeaderDefinitionOnly,
	// )

	// // Depends on QtCore/Gui/Widgets
	// generate(
	// 	"qt-extras/scintillaedit",
	// 	[]string{
	// 		filepath.Join(extraLibsDir, "scintilla/qt/ScintillaEdit/ScintillaEdit.h"),
	// 	},
	// 	AllowAllHeaders,
	// 	clangBin,
	// 	"--std=c++1z "+pkgConfigCflags("ScintillaEdit"),
	// 	outDir,
	// 	(&clangMatchUnderPath{filepath.Join(extraLibsDir, "scintilla")}).Match,
	// )

	// FLUSH all known typedefs / ...

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
			"WebEngineCore", "WebEngineQuick", "WebEngineWidgets",

			"Qml", "Quick",
		}

		for _, lib := range libs6 {
			generate(
				"Qt"+lib, "Qt6"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt6/Qt" + lib},
				clangBin, "--std=c++17", seaqtDir,
			)
		}
	}

	// Qt 6 Charts
	// Depends on QtCore/Gui/Widgets
	// generate(
	// 	"qt-restricted-extras/charts6",
	// 	[]string{
	// 		"/usr/include/x86_64-linux-gnu/qt6/QtCharts",
	// 	},
	// 	AllowAllHeaders,
	// 	clangBin,
	// 	"--std=c++17 "+pkgConfigCflags("Qt6Charts"),
	// 	outDir,
	// 	ClangMatchSameHeaderDefinitionOnly,
	// )

	// Qt 6 QScintilla
	// Depends on QtCore/Gui/Widgets, QPrintSupport
	// generate(
	// 	"qt-restricted-extras/qscintilla6",
	// 	[]string{
	// 		"/usr/include/x86_64-linux-gnu/qt6/Qsci",
	// 	},
	// 	AllowAllHeaders,
	// 	clangBin,
	// 	"--std=c++17 "+pkgConfigCflags("Qt6PrintSupport"),
	// 	outDir,
	// 	ClangMatchSameHeaderDefinitionOnly,
	// )
}
