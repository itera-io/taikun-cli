#!/bin/sh

set -e

# replace main.go
cp main.go main.go.bak
cat > main.go << EOF
package main

import (
	"github.com/itera-io/taikun-cli/cmd/root"
	"github.com/itera-io/taikun-cli/utils/docs"
)

func main() {
	docs.PrintCommandTree(root.NewCmdRoot())
}
EOF

# build
go build -o taikun

# write command tree
./taikun > COMMAND_TREE.md

# reset main.go
mv main.go.bak main.go
