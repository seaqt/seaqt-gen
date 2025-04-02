SHELL := /bin/bash
GO := go

.PHONY: all
all: genbindings

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

.PHONY: test-cmd
test-cmd: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings go test ./cmd/...

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
