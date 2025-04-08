.PHONY: all
all:

rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# TODO Can we avoid repeating the qt module list?
COREPRIV := $(shell pkg-config --variable=includedir Qt6Core)/QtCore/$(shell pkg-config --modversion Qt6Core)

CXXFLAGS := $(shell pkg-config --cflags Qt6Core Qt6Gui Qt6Widgets Qt6Network Qt6Multimedia Qt6MultimediaWidgets Qt6PrintSupport Qt6Svg Qt6SvgWidgets Qt6SpatialAudio Qt6WebChannel Qt6WebEngineCore Qt6WebEngineQuick Qt6WebEngineWidgets Qt6Qml Qt6Quick)
OBJS := $(patsubst %.cpp,%.o, $(call rwildcard,.,*.cpp))

# -fsyntax-only means no object files will be written which means everything will be recompiled every time
.PHONY: $(OBJS)

$(OBJS): %.o: %.cpp
	@echo Checking $<...
	@$(CXX) -fsyntax-only $< $(CXXFLAGS) -I$(COREPRIV) -I$(COREPRIV)/QtCore -fPIC

test: $(OBJS)
