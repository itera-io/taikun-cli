BINARY=taikun

default: install

.PHONY: build
build:
	go build -o ${BINARY} .

.PHONY: install
install: build
	mv -v ${BINARY} ${GOPATH}/bin
