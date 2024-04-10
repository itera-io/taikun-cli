#!/usr/bin/env bash

if [[ $# -ne 1 ]]; then
  echo "Rerun failed shellspec tests"
  echo "Usage: ./rerun_failed_tests.sh <logfile>"
  exit 2
fi

if [[ ! -d ./cmd ]]; then
  echo "Rerun failed shellspec tests"
  echo "Error: ./cmd/ not found, must run from project root"
  exit 2
fi

logfile="$1"
fctx=$(grep FAILED "$logfile" | cut -d '-' -f 2 | cut -d ' ' -f 2 | uniq)

if [[ -z $fctx ]]; then
  echo "No tests failed"
  exit 0
fi

failures=0

# Shellspec Context must be named EXACLY as the path leading to the folder containing the spec .sh file.
for context in $fctx; do
  if [[ ! -d "./cmd/$context" ]]; then
    echo "Error: invalid context $context, please fix context name"
    exit 1
  fi
  if ! shellspec --shell bash --format tap --jobs "$(nproc)" ./cmd/$context; then
    failures=$((failures+1))
  fi
done

if [[ $failures -gt 0 ]]; then
  exit 1
fi
