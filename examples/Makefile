dirs := qt5 qt6

.PHONY: all
all: qt5 qt6

.PHONY: qt5 qt6
qt5 qt6:
	$(MAKE) -C $@

clean: $(dirs:%=%-clean)

$(dirs:%=%-clean):
	$(MAKE) -C $(@:%-clean=%) clean
