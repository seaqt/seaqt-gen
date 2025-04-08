.PHONY: all
all:

rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# TODO Can we avoid repeating the qt module list?
COREPRIV := $(shell pkg-config --variable=includedir Qt5Core)/QtCore/$(shell pkg-config --modversion Qt5Core)

CXXFLAGS := $(shell pkg-config --cflags Qt5Core Qt5Gui Qt5Widgets Qt5Network Qt5Multimedia Qt5MultimediaWidgets Qt5PrintSupport Qt5Script Qt5Svg Qt5WebChannel Qt5WebEngineCore Qt5Qml Qt5Quick Qt5WebKit Qt5WebKitWidgets)
OBJS := $(patsubst %.cpp,%.o, $(call rwildcard, ,*.cpp))

# -fsyntax-only means no object files will be written which means everything will be recompiled every time
.PHONY: $(OBJS)

$(OBJS): %.o: %.cpp
	@echo Checking $<...
	@$(CXX) -fsyntax-only $< $(CXXFLAGS) -I$(COREPRIV) -I$(COREPRIV)/QtCore -fPIC

test: $(OBJS)
