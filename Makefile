
buildAt:=$(shell date "+%Y-%m-%d %H:%M:%S")
commitId:=$(shell git rev-parse --short HEAD)
branch:=$(shell git symbolic-ref --short -q HEAD)
version:=latest
imageName:=jeyrce/ioseek:${version}
pkgDir:=ioseek-${version}

.phony: all
all: image pkg
	@echo "make all"

.phony: binary
binary: init
	CGO_ENABLED=0 go build -ldflags " \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaCommitId=${commitId}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaBranch=${branch}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaVersion=${version}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaBuildAt=${buildAt}' \
	" \
	-o _output/ioseek ioseek.go

.phony: image
image:
	docker buildx rm ioseek
	docker buildx create --name ioseek --bootstrap --use
	docker buildx build -t ${imageName} \
		--build-arg pkgDir=${pkgDir} \
		--build-arg commitId=${commitId} \
		--build-arg goProxy=${goProxy} \
		--platform linux/386,linux/amd64,linux/arm64 \
		--push \
		.

.phony: pkg
pkg: binary
	mkdir -p "_output/${pkgDir}"
	cp _output/ioseek \
		README.md \
	  	LICENSE \
	_output/${pkgDir} ; \
 	tar -cvf _output/${pkgDir}.tar.gz _output/${pkgDir}

.phony: init
init:
	rm -rf _output/**
