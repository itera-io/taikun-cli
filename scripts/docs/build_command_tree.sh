#!/bin/sh

set -e

cmd_tree_path="COMMAND_TREE.md"

# save previous command tree
if [ -f $cmd_tree_path ]; then
  mv $cmd_tree_path "${cmd_tree_path}.bak"
fi

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
./taikun > $cmd_tree_path

# reset main.go
mv main.go.bak main.go

# exit successfully if the command tree was generated for the first time
if [ ! -f "${cmd_tree_path}.bak" ]; then
  echo "Command tree was generated for the first time."
  exit 0
fi

# exit with non-zero status if the command tree has *not* changed
if diff $cmd_tree_path "${cmd_tree_path}.bak" 1>/dev/null; then
  # remove previous command tree
  rm "${cmd_tree_path}.bak"
  echo "Command tree hasn't changed."
  exit 1
fi

# remove previous command tree
rm "${cmd_tree_path}.bak"
echo "Command tree has changed."
