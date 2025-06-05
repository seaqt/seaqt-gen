dirs := $(wildcard seaqt-*)
nimdirs := $(wildcard nim-seaqt-*)

.PHONY: all
all: $(dirs) $(nimdirs)

.PHONY: $(dirs)
$(dirs):
	echo $(dirs)
	$(MAKE) -C $@

.PHONY: $(nimdirs)
$(nimdirs):
	nimble examples_$$(echo $(@:nim-seaqt-%=%) | tr -- . _)
	QT_QPA_PLATFORM=offscreen nimble test_$$(echo $(@:nim-seaqt-%=%) | tr -- . _)

.PHONY: clean $(dirs:%=%-clean) $(nim-dirs:%=%-clean)
clean: $(dirs:%=%-clean) $(nim-dirs:%=%-clean)

$(dirs:%=%-clean):
	$(MAKE) -C $(@:%-clean=%) clean

$(nim-dirs:%=%-clean):
	find nim-* -name main -delete
