BUILDSTAMPS := docker/genbindings.docker-buildstamp
DOCKER := docker
SHELL := /bin/bash

# DOCKEREXEC runs the target command in the `genbindings` docker container.
# It mounts in the current GOCACHE and GOMODCACHE.
DOCKEREXEC = mkdir -p "$$(go env GOCACHE)" && \
	mkdir -p "$$(go env GOMODCACHE)" && \
	$(DOCKER) run \
	--user "$$(id -u):$$(id -g)" \
	-v "$$(go env GOCACHE):/.cache/go-build" \
	-v "$$(go env GOMODCACHE):/go/pkg/mod" \
	-v "$$PWD:/src" \
	-w /src \
	miqt/genbindings:latest \
	/bin/bash -c

.PHONY: all
all: genbindings copy-libseaqt

docker/genbindings.docker-buildstamp: docker/genbindings.Dockerfile
	$(DOCKER) build -t miqt/genbindings:latest -f docker/genbindings.Dockerfile .
	touch $@

.PHONY: clean
clean:
	$(DOCKER) image rm -f miqt/genbindings:latest
	rm -f $(BUILDSTAMPS)

.PHONY: genbindings
genbindings: $(BUILDSTAMPS)
	$(DOCKEREXEC) 'cd cmd/genbindings && go build && ./genbindings' 2>log.txt

copy-libseaqt: genbindings
	cd gen/ ;\
	for a in *seaqt-*; do cp -ar ../libseaqt $$a/seaqt; done ;

gencommits: copy-libseaqt
	cd gen/ ;\
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

.PHONY : all clean genbindings gencommits copy-libseaqt github-ssh

github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

# test C++ files
# TODO generate makefile in each version for these

rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# TODO Can we avoid repeating the qt module list?
COREPRIV515 := $(shell pkg-config --variable=includedir Qt5Core)/QtCore/$(shell pkg-config --modversion Qt5Core)

CXXFLAGS515 := $(shell pkg-config --cflags Qt5Core Qt5Gui Qt5Widgets Qt5Network Qt5Multimedia Qt5MultimediaWidgets Qt5PrintSupport Qt5Script Qt5Svg Qt5WebChannel Qt5WebEngineCore Qt5Qml Qt5Quick Qt5WebKit Qt5WebKitWidgets)
OBJS515 := $(patsubst %.cpp,%.o, $(call rwildcard,gen/seaqt-5.15,*.cpp))

# -fsyntax-only means no object files will be written which means everything will be recompiled every time
.PHONY: $(OBJS515)

$(OBJS515): %.o: %.cpp
	@echo Checking $<...
	@$(CXX) -fsyntax-only $< $(CXXFLAGS515) -I$(COREPRIV515) -I$(COREPRIV515)/QtCore -fPIC

test-gen-5.15: $(OBJS515)

# TODO Can we avoid repeating the qt module list?
COREPRIV64 := $(shell pkg-config --variable=includedir Qt6Core)/QtCore/$(shell pkg-config --modversion Qt6Core)

CXXFLAGS64 := $(shell pkg-config --cflags Qt6Core Qt6Gui Qt6Widgets Qt6Network Qt6Multimedia Qt6MultimediaWidgets Qt6PrintSupport Qt6Svg Qt6SvgWidgets Qt6SpatialAudio Qt6WebChannel Qt6WebEngineCore Qt6WebEngineQuick Qt6WebEngineWidgets Qt6Qml Qt6Quick)
OBJS64 := $(patsubst %.cpp,%.o, $(call rwildcard,gen/seaqt-6.4,*.cpp))

# -fsyntax-only means no object files will be written which means everything will be recompiled every time
.PHONY: $(OBJS64)

$(OBJS64): %.o: %.cpp
	@echo Checking $<...
	@$(CXX) -fsyntax-only $< $(CXXFLAGS64) -I$(COREPRIV64) -I$(COREPRIV64)/QtCore -fPIC

test-gen-6.4: $(OBJS64)

test: test-gen-5.15 test-gen-6.4
