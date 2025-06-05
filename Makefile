SHELL := /bin/bash
GO := go

VERSIONS := 5.15 6.4

.PHONY: all
all: genbindings copy-libseaqt

cmd/miqt-docker/miqt-docker: go.mod cmd/miqt-docker/*.go docker/*.Dockerfile
	$(GO) build -o cmd/miqt-docker/miqt-docker ./cmd/miqt-docker

.PHONY: clean-cache
clean-cache:
	rm -f cmd/genbindings/cachedir/*.json

cmd/genbindings/genbindings: go.mod cmd/genbindings/*.go
	$(GO) build -o cmd/genbindings/genbindings ./cmd/genbindings

.PHONY: genbindings
genbindings: cmd/miqt-docker/miqt-docker cmd/genbindings/genbindings
	cd cmd/genbindings && ../miqt-docker/miqt-docker genbindings ./genbindings

.PHONY: build-all
build-all: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings go build ./...

# reset-gen is used when rebasing the submodules, to reset the "qt-<version>" branches
reset-gen:
	git submodule update --init --recursive
	for name in gen/seaqt-*; do \
		git -C $$name switch $$(git config -f .gitmodules submodule.$$name.branch); \
		git -C $$name reset --hard $$(git -C $$name rev-list --max-parents=0 HEAD); \
	done

.PHONY: copy-libseaqt
copy-libseaqt: | genbindings
	cd gen/ ;\
	for a in *seaqt-*; do cp -ar ../libseaqt/* $$a; done ;\
	for v in $(VERSIONS); do \
		for a in *seaqt-$$v; do cp -ar ../libseaqt-$$v/* $$a; done ;\
	done

gencommits: genbindings copy-libseaqt
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

genbranches:
	export GIT_BRANCH=$$(git rev-parse --abbrev-ref HEAD);\
	cd gen/ ;\
	for v in $(VERSIONS); do \
		for a in *seaqt-$$v; do cd $$a && git switch -C $$GIT_BRANCH-$$v && cd ..; done ;\
	done

.PHONY: all genbindings gencommits copy-libseaqt github-ssh

github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

$(VERSIONS:%=test-%): cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings /bin/bash -c 'cd gen/seaqt-$(@:test-%=%) && make -j$$(nproc) test'

.PHONY: test $(VERSIONS:%=test-%)
test: $(VERSIONS:%=test-%)

.PHONY: examples $(VERSIONS:%=examples-%)
$(VERSIONS:%=examples-%): cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings /bin/bash -c 'cd examples && make -j$$(nproc) seaqt-$(@:examples-%=%)'

examples: $(VERSIONS:%=examples-%)
