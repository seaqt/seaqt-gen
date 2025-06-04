SHELL := /bin/bash
GO := go

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

copy-libseaqt: genbindings
	cd gen/ ;\
	for a in *seaqt-*; do cp -ar ../libseaqt/* $$a; done ;\
	for a in *seaqt-5.15; do cp -ar ../libseaqt-5.15/* $$a; done ;\
	for a in *seaqt-6.4; do echo $a; cp -ar ../libseaqt-6.4/* $$a; done ;

gencommits: copy-libseaqt
	cd gen/ ;\
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

.PHONY: all genbindings gencommits copy-libseaqt github-ssh

github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

.PHONY: test-gen-5.15
test-gen-5.15: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings /bin/bash -c 'cd gen/seaqt-5.15 && make -j$$(nproc) test'

.PHONY: test-gen-6.4
test-gen-6.4: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings /bin/bash -c 'cd gen/seaqt-6.4 && make -j$$(nproc) test'

.PHONY: test
test: test-gen-5.15 test-gen-6.4

.PHONY: examples
examples:
	$(MAKE) -C $@

