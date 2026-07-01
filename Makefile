SHELL := /bin/bash
GO := go

VERSIONS := 5.15 6.4 6.5 6.8 6.11

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
genbindings: $(VERSIONS:%=genbindings-%)

# Build the shared base Docker image (Go + dev tools + aqtinstall, no Qt)
.PHONY: genbindings-base
genbindings-base: cmd/miqt-docker/miqt-docker docker/genbindings-base.Dockerfile
	./cmd/miqt-docker/miqt-docker genbindings-base true

# Generate bindings for each Qt version
$(VERSIONS:%=genbindings-%): genbindings-base cmd/miqt-docker/miqt-docker cmd/genbindings/genbindings
	export VERSION=$(@:genbindings-%=%) && \
	cd cmd/genbindings && ../miqt-docker/miqt-docker genbindings-qt$${VERSION} ./genbindings -version=$${VERSION} ./genbindings && cd ../.. && \
	for a in libseaqt/*; do cp -aR "$$a" "gen/seaqt-$${VERSION}/"; done && \
	for a in libseaqt-$${VERSION%.*}/*; do cp -aR "$$a" gen/seaqt-$${VERSION}/; done && \
	for a in libseaqt/*; do cp -aR "$$a" "gen/nim-seaqt-$${VERSION}/"; done && \
	for a in libseaqt-$${VERSION%.*}/*; do cp -aR "$$a" gen/nim-seaqt-$${VERSION}/; done

.PHONY: genbindings test-cmd
test-cmd: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings go test ./cmd/...

.PHONY: build-all
build-all: cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings go build ./...

# reset-gen resets each seaqt submodule's version branch to its first commit,
# then overwrites the first commit's tree with master's contents.
reset-gen:
	git submodule update --init --recursive
	for name in gen/seaqt-*; do \
		branch=$$(git -C $$name config -f .gitmodules submodule.$$name.branch); \
		first=$$(git -C $$name rev-list --max-parents=0 HEAD); \
		cd $$name && \
		git switch "$$branch" && \
		git reset --hard "$$first" && \
		git checkout master -- . && \
		git add -A && \
		git commit --amend --no-edit && \
		cd ..; \
	done

# reset-gen-nim resets each nim-seaqt submodule's version branch to its first commit,
# then overwrites the first commit's tree with master's contents.
reset-gen-nim:
	git submodule update --init --recursive
	for name in gen/nim-seaqt-*; do \
		branch=$$(git -C $$name config -f .gitmodules submodule.$$name.branch); \
		first=$$(git -C $$name rev-list --max-parents=0 HEAD); \
		cd $$name && \
		git switch "$$branch" && \
		git reset --hard "$$first" && \
		git checkout master -- . && \
		git add -A && \
		git commit --amend --no-edit && \
		cd ..; \
	done

gencommits: genbindings
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

genbranches:
	export GIT_BRANCH=$$(git rev-parse --abbrev-ref HEAD);\
	cd gen/ ;\
	for v in $(VERSIONS); do \
		for a in *seaqt-$$v; do cd $$a && git switch -C $$GIT_BRANCH-$$v && cd ..; done ;\
	done

.PHONY: all genbindings gencommits github-ssh
github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

.PHONY: test $(VERSIONS:%=test-%)
$(VERSIONS:%=test-%): genbindings-base cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings-qt$(@:test-%=%) /bin/bash -c 'cd gen/seaqt-$(@:test-%=%) && make -j$$(nproc) test'
test: $(VERSIONS:%=test-%)

.PHONY: examples $(VERSIONS:%=examples-%)
$(VERSIONS:%=examples-%): genbindings-base cmd/miqt-docker/miqt-docker
	./cmd/miqt-docker/miqt-docker genbindings-qt$(@:examples-%=%) /bin/bash -c 'cd examples && make -j$$(nproc) seaqt-$(@:examples-%=%) nim-seaqt-$(@:examples-%=%)'
examples: $(VERSIONS:%=examples-%)
