BUILDSTAMPS = docker/genbindings.docker-buildstamp
DOCKER = docker

all: genbindings copy-libseaqt

docker/genbindings.docker-buildstamp: docker/genbindings.Dockerfile
	$(DOCKER) build -t miqt/genbindings:latest -f docker/genbindings.Dockerfile .
	touch $@

clean:
	$(DOCKER) image rm -f miqt/genbindings:latest
	rm -f $(BUILDSTAMPS)

copy-libseaqt: genbindings
	cd gen/ ;\
	for a in *seaqt-*; do cp -ar ../libseaqt $$a/seaqt; done ;

gencommits: copy-libseaqt
	cd gen/ ;\
	git submodule foreach git add -A ;\
	git submodule foreach 'git commit -am "update bindings" || :'

genbindings: $(BUILDSTAMPS)
	mkdir -p ~/.cache/go-build
	$(DOCKER) run --user $$(id -u):$$(id -g) -v ~/.cache/go-build:/.cache/go-build -v $$PWD:/src -w /src miqt/genbindings:latest /bin/bash -c 'cd cmd/genbindings && go build && ./genbindings' 2>log.txt

.PHONY : all clean genbindings gencommits copy-libseaqt github-ssh

github-ssh:
	git config url."git@github.com:".insteadOf "https://github.com/"
	git submodule foreach --recursive 'git config url."git@github.com:".insteadOf "https://github.com/"'

