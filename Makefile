
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
	CGO_ENABLED=1; go build -ldflags " \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaCommitId=${commitId}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaBranch=${branch}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaVersion=${version}' \
		-X 'github.com/skrbox/ioseek/pkg/conf.metaBuildAt=${buildAt}' \
	" \
	-o _output/ioseek ioseek.go

.phony: image
image:
	docker build -t ${imageName} --build-arg pkgDir=${pkgDir} .
	docker push ${imageName}

.phony: pkg
pkg: binary
	mkdir -p "_output/${pkgDir}"
	cp _output/ioseek \
		ioseek.yml \
		README.md \
	  	LICENSE \
	_output/${pkgDir} ; \
 	tar -cvf _output/${pkgDir}.tar.gz _output/${pkgDir}

.phony: init
init:
	rm -rf _output/**
