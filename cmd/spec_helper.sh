# cmd/spec_helper.sh

# shellcheck source=/dev/null
source "$(dirname "${BASH_SOURCE[0]}")/../scripts/first_fail_trace.sh"

# Turn on tracing only when asked (keeps normal runs quiet)
[[ "${FIRST_FAIL_TRACE:-0}" = "1" ]] && first_fail_trace_on

# Always install the taikun wrapper for specs (or gate it with env if you prefer)
install_taikun_wrapper
