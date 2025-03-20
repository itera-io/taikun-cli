BINARY=taikun

default: install

.PHONY: build
build:
	go mod tidy -compat=1.23
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

.PHONY: lint
lint: install
	golangci-lint run

.PHONY: release
release: install
	goreleaser --snapshot --clean

.PHONY: ci_local
ci_local: install lint test release
	echo "======================================"
	echo "Local mock of the CI pipeline complete"