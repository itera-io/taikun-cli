# cmd/spec_helper.sh
# Load first-fail tracer for all specs in ./cmd

# shellcheck source=/dev/null
source "$(dirname "${BASH_SOURCE[0]}")/../scripts/first_fail_trace.sh"

# Only enable if requested (quiet by default)
if [[ "${FIRST_FAIL_TRACE:-0}" = "1" ]]; then
  first_fail_trace_on
fi
