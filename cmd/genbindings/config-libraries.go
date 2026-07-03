package main

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	ErrNoVersionSpecified = errors.New("Qt version must be specified (use -version=5.15 etc)")
)

// getQtVersion returns the Qt version string for the given major version.
// If minor is provided (e.g. "15"), it is used directly as "major.minor".
// Otherwise it falls back to QT{major}_VERSION env var, then defaults.
func getQtVersion(major int, minor string) string {
	if minor != "" {
		return strconv.Itoa(major) + "." + minor
	}
	envKey := "QT"
	if major == 5 {
		envKey = "QT5"
	}
	if major == 6 {
		envKey = "QT6"
	}
	if v := os.Getenv(envKey + "_VERSION"); v != "" {
		parts := strings.Split(v, ".")
		if len(parts) >= 2 {
			return parts[0] + "." + parts[1]
		}
		return v
	}
	if major == 5 {
		return "5.15"
	}
	return "6.4"
}

func ProcessLibraries(clangBin, outDir, extraLibsDir string, buildVersion string) {
	if buildVersion == "" {
		panic(ErrNoVersionSpecified)
	}

	parts := strings.SplitN(buildVersion, ".", 2)
	major, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("invalid version: " + buildVersion)
	}
	var minor string
	if len(parts) > 1 {
		minor = parts[1]
	}
	ver := getQtVersion(major, minor)
	CacheSubdir = "qt" + ver

	if major == 5 {
		buildQt5(clangBin, outDir, extraLibsDir, ver)
	} else if major == 6 {
		buildQt6(clangBin, outDir, extraLibsDir, ver)
	} else {
		panic("invalid version: " + buildVersion)
	}
}

func buildQt5(clangBin, outDir, extraLibsDir string, version string) {
	flushKnownTypes()
	InsertTypedefs(false)

	qt5Include := os.Getenv("QT5_INCLUDE")
	if qt5Include == "" {
		qt5Include = "/usr/include/x86_64-linux-gnu/qt5"
	}

	qtDir := filepath.Join(outDir, "seaqt-"+version)
	seaqtDir := filepath.Join(qtDir, "seaqt")
	nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+version, "seaqt")
	nimbleDir := filepath.Join(outDir, "nim-seaqt-"+version)

	os.RemoveAll(seaqtDir)
	os.RemoveAll(nimSeaqtDir)

	writeNimbleFile(nimbleDir, version)

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
			"Qt"+lib, "Qt5"+lib,
			[]string{filepath.Join(qt5Include, "Qt"+lib)},
			clangBin, "", seaqtDir, nimSeaqtDir,
		)
	}
}

func buildQt6(clangBin, outDir, extraLibsDir string, version string) {
	flushKnownTypes()
	InsertTypedefs(true)

	qt6Include := os.Getenv("QT6_INCLUDE")
	if qt6Include == "" {
		qt6Include = "/usr/include/x86_64-linux-gnu/qt6"
	}

	qtDir := filepath.Join(outDir, "seaqt-"+version)
	seaqtDir := filepath.Join(qtDir, "seaqt")
	nimSeaqtDir := filepath.Join(outDir, "nim-seaqt-"+version, "seaqt")
	nimbleDir := filepath.Join(outDir, "nim-seaqt-"+version)

	os.RemoveAll(seaqtDir)
	os.RemoveAll(nimSeaqtDir)

	writeNimbleFile(nimbleDir, version)

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

	for _, lib := range modules {
		generate(
			"Qt"+lib, "Qt6"+lib,
			[]string{filepath.Join(qt6Include, "Qt"+lib)},
			clangBin, "--std=c++17", seaqtDir, nimSeaqtDir,
		)
	}
}
