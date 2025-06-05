dirs := $(wildcard seaqt-*)

.PHONY: all
all: $(dirs)

.PHONY: $(dirs)
$(dirs):
	$(MAKE) -C $@

clean: $(dirs:%=%-clean)

$(dirs:%=%-clean):
	$(MAKE) -C $(@:%-clean=%) clean
