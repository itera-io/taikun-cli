BINARY=taikun
GOLANGCI_LINTERS_VERSION := v2.11.4

default: install

deps: shellspec-install goreleaser-install go-linters-install ## Installing development prerequisites locally

shellspec-install: ## Installs shellspec locally (to run shell-based unit tests)
	- curl -fsSL https://git.io/shellspec | sh

goreleaser-install: ## Installs goreleaser binary with go install
	go install github.com/goreleaser/goreleaser/v2@latest

go-linters-install: ## Installs Golang's linters locally for verification
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin ${GOLANGCI_LINTERS_VERSION}

.PHONY: build
build: go-vendor ## Builds taikun-cli binary
	go build -o ${BINARY} .

.PHONY: dockerbuild
dockerbuild: ## Builds Docker image for taikun-cli
	DOCKER_BUILDKIT=1 docker build --rm --target bin --output . .

.PHONY: install
install: build ## Installs built Go binary as a Go module
	mv -v ${BINARY} ${GOPATH}/bin

.PHONY: test
test: install ## Runs unit tests against Taikun's API
	shellspec --shell bash --format tap ${TESTARGS} | tee shellspec.log

.PHONY: vimtest
vimtest: ## Runs unit tests against Taikun's API and shows only failed ones
	@shellspec --shell bash --format failures ${TESTARGS}

.PHONY: lint
lint: go-linters-install ## Performs linting against codebase
	golangci-lint run --timeout 5m

.PHONY: release
release: install ## Releases Go binary
	goreleaser --snapshot --clean

.PHONY: ci_local
ci_local: install lint test release ## Mimics CI/CD behavior locally and runs test and release phases
	echo "======================================"
	echo "Local mock of the CI pipeline complete"

go-tidy: ## Runs go mod tidy
	go mod tidy

go-vendor: go-tidy ## Runs go mod tidy && go mod vendor
	go mod vendor

clean-vendor: ## Removes vendor folder
	rm -rf vendor

.PHONY: help
help: # Credits to https://gist.github.com/prwhite/8168133 for this handy oneliner
	@awk 'BEGIN {FS = ":.*##"; printf "Usage: make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
