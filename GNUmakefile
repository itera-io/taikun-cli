BINARY=taikun

default: install

build:
	go build -o ${BINARY} .

dockerbuild:
	DOCKER_BUILDKIT=1 docker build --rm --target bin --output . .

install: build
	mv -v ${BINARY} ${GOPATH}/bin

.PHONY: install dockerbuild build