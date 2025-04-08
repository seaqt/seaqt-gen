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
	for a in *seaqt-*; do cp -ar ../libseaqt $$a/seaqt; done ;\
	for a in *seaqt-5.15; do cp -ar ../libseaqt-5.15/* $$a; done ;\
	for a in *seaqt-6.4; do echo $a; cp -ar ../libseaqt-6.4/* $$a; done ;

gencommits: copy-libseaqt
	cd gen/ ;\
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

.PHONY : all clean genbindings gencommits copy-libseaqt github-ssh

github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

test-gen-5.15: $(BUILDSTAMPS)
	$(DOCKEREXEC) 'cd gen/seaqt-5.15 && make -j$$(nproc) test'

test-gen-6.4: $(BUILDSTAMPS)
	$(DOCKEREXEC) 'cd gen/seaqt-6.4 && make -j$$(nproc) test'

test: test-gen-5.15 test-gen-6.4
