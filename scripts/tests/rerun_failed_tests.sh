#!/usr/bin/env bash
set -euo pipefail

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
if [[ ! -f "$logfile" ]]; then
  echo "Error: logfile '$logfile' not found"
  exit 2
fi

# Extract failing contexts from TAP log lines like:
# "not ok 83 - cloudcredential/google/list list google cloud credential # FAILED"
fctx=$(grep -F "FAILED" "$logfile" | cut -d '-' -f 2 | cut -d ' ' -f 2 | uniq)

if [[ -z ${fctx:-} ]]; then
  echo "No tests failed"
  exit 0
fi

failures=0

# Ensure spec_helper is loaded from ./cmd and enable tracing for reruns
export SHELLSPEC_LOAD_PATH="cmd"
export FIRST_FAIL_TRACE=1

for context in $fctx; do
  if [[ ! -d "./cmd/$context" ]]; then
    echo "Error: invalid context $context, please fix context name"
    exit 1
  fi
  if ! shellspec --shell bash --format tap --jobs "$(nproc)" --load-path cmd "./cmd/$context"; then
    failures=$((failures+1))
  fi
done

if [[ $failures -gt 0 ]]; then
  exit 1
fi
