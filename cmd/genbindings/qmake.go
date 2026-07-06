package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	ErrNoQtFound = errors.New("could not find a working qmake for Qt5 or Qt6")
)

// qmakeBin is the resolved qmake binary path, determined at init time.
var qmakeBin string

// initQmakeBin discovers the qmake binary by checking:
// 1. QMAKE environment variable
// 2. Plain qmake in PATH
// 3. Falls back to version-specific qmake (qmake-qt5, qmake6)
func initQmakeBin() {
	if path := os.Getenv("QMAKE"); path != "" {
		err := exec.Command(path, "-query", "QT_VERSION").Run()
		if err == nil {
			qmakeBin = path
			return
		}
		log.Printf("qmake from QMAKE env (%q) failed: %v — trying fallback", path, err)
	}

	if _, err := exec.LookPath("qmake"); err == nil {
		if out, err := exec.Command("qmake", "-query", "QT_VERSION").Output(); err == nil {
			_ = strings.TrimSpace(string(out))
			qmakeBin = "qmake"
			return
		}
		log.Printf("qmake in PATH failed: %v — trying version-specific fallback", err)
	}

	for _, name := range []string{"qmake-qt5", "qmake6"} {
		if _, err := exec.LookPath(name); err == nil {
			if out, err := exec.Command(name, "-query", "QT_VERSION").Output(); err == nil {
				_ = strings.TrimSpace(string(out))
				qmakeBin = name
				return
			}
		}
	}

	panic(ErrNoQtFound)
}

// qmakeQuery runs qmake -query <key> using the globally resolved qmake binary
// and returns the result. Panics if qmake is not found.
func qmakeQuery(key string) string {
	if qmakeBin == "" {
		initQmakeBin()
	}

	stdout, err := exec.Command(qmakeBin, "-query", key).Output()
	if err != nil {
		log.Printf("qmake(-query %s): %v", key, err)
		panic(err)
	}
	return strings.TrimSpace(string(stdout))
}
