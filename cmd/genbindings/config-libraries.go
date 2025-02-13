package main

import (
	"os"
	"path/filepath"
)

func ProcessLibraries(clangBin, outDir, extraLibsDir string) {

	flushKnownTypes()
	InsertTypedefs(false)

	{
		seeqtDir := filepath.Join(outDir, "qt-5.15", "seeqt")
		os.RemoveAll(seeqtDir)

		// TODO more modules
		libs5 := []string{
			// https://doc.qt.io/qt-5/qtmodules.html#qt-essentials
			"Core", "Gui", "Widgets", "Multimedia", "MultimediaWidgets", "Network", "Qml", "Quick",

			// https://doc.qt.io/qt-5/qtmodules.html#qt-add-ons
			"PrintSupport", "Script", "Svg", "WebChannel",
			"WebEngineCore", "WebEngine", "WebEngineWidgets",

			// TODO Not sure where these are from :)
			"WebKit", "WebKitWidgets"}

		for _, lib := range libs5 {
			generate(
				"Qt"+lib, "Qt5"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt5/Qt" + lib},
				clangBin, pkgConfigCflags("Qt5"+lib), seeqtDir,
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
		seeqtDir := filepath.Join(outDir, "qt-6.4", "seeqt")
		os.RemoveAll(seeqtDir)

		libs6 := []string{
			"Core", "Gui", "Widgets", "Multimedia", "MultimediaWidgets", "Network", "Qml", "Quick",

			"PrintSupport", "Svg", "SvgWidgets", "SpatialAudio", "WebChannel",
			"WebEngineCore", "WebEngineQuick", "WebEngineWidgets"}

		for _, lib := range libs6 {
			generate(
				"Qt"+lib, "Qt6"+lib,
				[]string{"/usr/include/x86_64-linux-gnu/qt6/Qt" + lib},
				clangBin, "--std=c++17 "+pkgConfigCflags("Qt6"+lib), seeqtDir,
			)
		}
	}

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
