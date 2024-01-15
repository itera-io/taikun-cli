#!/bin/sh

set -e

# create documentation directory
mkdir -p docs/markdown

# replace main.go
cat > main.go << EOF
package main

import (
	"github.com/itera-io/taikun-cli/cmd/root"
	"github.com/spf13/cobra/doc"
)

func main() {
	rootCmd := root.NewCmdRoot()
	if err := doc.GenMarkdownTree(rootCmd, "docs/markdown"); err != nil {
		panic(err)
	}
}
EOF

# update dependencies
go mod tidy -compat=1.20

# build
go build -o taikun

# generate docs
./taikun
