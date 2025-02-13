BUILDSTAMPS = docker/genbindings.docker-buildstamp
DOCKER = docker

all: genbindings

docker/genbindings.docker-buildstamp: docker/genbindings.Dockerfile
	$(DOCKER) build -t miqt/genbindings:latest -f docker/genbindings.Dockerfile .
	touch $@

clean:
	$(DOCKER) image rm -f miqt/genbindings:latest
	rm -f $(BUILDSTAMPS)

gencommits:
	cd gen/ ;\
	for a in qt-*; do cp -ar ../libmiqt $$a; done ; \
	git submodule foreach git add -A ;\
	git submodule foreach git commit -am "update bindings"

genbindings: $(BUILDSTAMPS)
	mkdir -p ~/.cache/go-build
	$(DOCKER) run --user $$(id -u):$$(id -g) -v ~/.cache/go-build:/.cache/go-build -v $$PWD:/src -w /src miqt/genbindings:latest /bin/bash -c 'cd cmd/genbindings && go build && ./genbindings'

.PHONY : all clean genbindings gencommits
