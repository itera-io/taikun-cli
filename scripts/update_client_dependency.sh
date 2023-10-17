#!/bin/sh

set -e

# check args
[[ $# -eq 1 ]] || (echo "Usage: $0 <dev|staging|main>" 1>&2 && exit 2)

# check requested client branch
branch="$1"
branch_is_valid=0
for valid_branch in dev staging main; do
    if [[ "$branch" = $valid_branch ]]; then
        branch_is_valid=1
        break
    fi
done
if [[ $branch_is_valid -eq 0 ]]; then
  echo -e "Usage: $0 <dev|staging|main>\n$branch is not a valid branch" 1>&2
  exit 2
fi

# check go.mod exists
[[ -f "./go.mod" ]] || (echo "go.mod not found, run from root of repo" 1>&2 && exit 2)

# client repo url
#repo="https://github.com/itera-io/taikungoclient.git"
repo="https://github.com/itera-io/taikungoclient.git"

# client dependency path
# path=github.com/itera-io/taikungoclient
path=github.com/itera-io/taikungoclient

# get latest commit hash
commit=$(git ls-remote $repo refs/heads/$branch | awk '{print $1}')

# update client dependency
go get $path@$commit

# tidy module dependencies
go mod tidy -compat=1.17
