BINARY=taikun

default: install

.PHONY: build
build:
	go build -o ${BINARY} .

.PHONY: dockerbuild
dockerbuild:
	DOCKER_BUILDKIT=1 docker build --rm --target bin --output . .

.PHONY: install
install: build
	mv -v ${BINARY} ${GOPATH}/bin

.PHONY: test
test: install
	shellspec --shell bash --format tap ${TESTARGS} | tee shellspec.log

.PHONY: vimtest
vimtest:
	@shellspec --shell bash --format failures ${TESTARGS}
