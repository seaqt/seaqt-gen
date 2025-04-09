.PHONY: all
all:

# TODO Can we avoid repeating the qt module list?
COREPRIV := $(shell pkg-config --variable=includedir Qt5Core)/QtCore/$(shell pkg-config --modversion Qt5Core)

CXXFLAGS := $(shell pkg-config --cflags Qt5Core Qt5Gui Qt5Widgets Qt5Network Qt5Multimedia Qt5MultimediaWidgets Qt5PrintSupport Qt5Script Qt5Svg Qt5WebChannel Qt5WebEngineCore Qt5Qml Qt5Quick Qt5WebKit Qt5WebKitWidgets)

DIRS := $(wildcard seaqt/Qt*)

# -fsyntax-only means no object files will be written which means everything will be recompiled every time
.PHONY: $(DIRS)

$(DIRS):
	@echo Checking $@...
	@cd $@ && cat *.cpp | clang++ $(CXXFLAGS) -xc++ -fsyntax-only -
	@cd $@ && clang++ $(CXXFLAGS) -xc++-header -fsyntax-only *.h

test: $(DIRS)
