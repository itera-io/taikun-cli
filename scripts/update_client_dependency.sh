#!/bin/sh

set -e

# check go.mod exists
[[ -f "./go.mod" ]] || (echo "go.mod not found, run from root of repo" && exit 2)

# client repo url
repo="https://github.com/itera-io/taikungoclient.git"

# client branch
branch=dev

# client dependency path
path=github.com/itera-io/taikungoclient

# get latest commit hash
commit=$(git ls-remote $repo refs/heads/$branch | awk '{print $1}')

# update client dependency
go get $path@$commit

# tidy module dependencies
go mod tidy
