ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
SEAQT_DIR := $(ROOT_DIR)../../gen/seaqt-6.4/

QT_LIBS := Core Gui Widgets

OPT := -flto=auto -Os

CFLAGS := -g -I$(SEAQT_DIR) -fPIC
CXXFLAGS := $(CFLAGS) -I. $(QT_LIBS:%=-I$(SEAQT_DIR)/seaqt/Qt%) $(shell pkg-config --cflags $(QT_LIBS:%=Qt6%))
LDFLAGS := -g -fPIC $(shell pkg-config --libs $(QT_LIBS:%=Qt6%)) -L.

examples := helloworld
mains := $(patsubst %,%/main,$(examples))

.PHONY: all
all: $(mains)

# Magic implicit rule(s) to find seaqt/X/X.cpp for X.o
.SECONDEXPANSION:
%.o: $$(wildcard $(SEAQT_DIR)seaqt/*/%.cpp)
	$(CXX) $(CXXFLAGS) $(OPT) -c $< -o $@

libqt.a: $(addsuffix .o,$(QT_LIBS:%=Qt%))
	ar rcs $@ $^

$(mains): %: %.c libqt.a
	$(CXX) $(CFLAGS) $(OPT) -o $@ $^ $(LDFLAGS)

clean:
	rm -rf *.a *.o $(mains)

.SECONDARY:
