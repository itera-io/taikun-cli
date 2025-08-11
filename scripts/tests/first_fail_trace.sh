# scripts/first_fail_trace.sh
# First-fail tracer + taikun error wrapper. Safe to source many times.

if [[ -n "${FIRST_FAIL_TRACE_LOADED:-}" ]]; then
  return 0
fi

first_fail_trace_on() {
  set -Eeuo pipefail
  shopt -s inherit_errexit 2>/dev/null || true
  export PS4='+ ${BASH_SOURCE##*/}:${LINENO}: $(date +%H:%M:%S)  '
  exec 3>&2
  export BASH_XTRACEFD=3 2>/dev/null || true
  set -x
  trap 'rc=$?;
        cmd=$BASH_COMMAND;
        file=${BASH_SOURCE[0]};
        line=${LINENO};
        {
          echo "ERROR first-fail -> ${file}:${line}";
          echo "  command: ${cmd}";
          if ( : "${#FUNCNAME[@]}" ) 2>/dev/null; then
            echo "  stack:";
            for i in "${!FUNCNAME[@]}"; do
              [[ "$i" == 0 ]] && continue
              echo "    #$i ${FUNCNAME[$i]} at ${BASH_SOURCE[$i]}:${BASH_LINENO[$((i-1))]}"
            done
          fi
        } >&3
        exit $rc' ERR
}

first_fail_trace_off() { set +x; trap - ERR || true; exec 3>&- 2>/dev/null || true; }

trace_run() { echo ">> ${BASH_SOURCE[1]##*/}:${BASH_LINENO[0]}: $*" >&3; "$@"; }

# ---- NEW: wrap taikun so printed errors become non-zero exits ----
install_taikun_wrapper() {
  [[ -n "${TAIKUN_WRAP_INSTALLED:-}" ]] && return 0

  taikun() {
    # Keep stdout clean for command substitution, mirror stderr as usual.
    local tmp_out tmp_err rc
    tmp_out="$(mktemp)"; tmp_err="$(mktemp)"

    # Optional: inject --debug if you want endpoint details and CLI supports it.
    local args=("$@")
    if [[ "${TAIKUN_DEBUG_ON:-0}" == "1" ]]; then
      args=(--debug "${args[@]}")
    fi

    command taikun "${args[@]}" >"$tmp_out" 2>"$tmp_err"
    rc=$?

    # Show stderr from real CLI (so you still see its messages)
    cat "$tmp_err" >&2

    # Treat printed Taikun errors (or HTTP 4xx/5xx) as failures.
    if (( rc != 0 )) || grep -Eiq 'Taikun Error|HTTP [45][0-9]{2}\b' "$tmp_err"; then
      cat "$tmp_out"   # preserve stdout for callers that capture it
      rm -f "$tmp_out" "$tmp_err"
      return 1
    fi

    cat "$tmp_out"
    rm -f "$tmp_out" "$tmp_err"
    return 0
  }

  export -f taikun
  export TAIKUN_WRAP_INSTALLED=1
}

export -f first_fail_trace_on first_fail_trace_off trace_run install_taikun_wrapper
export FIRST_FAIL_TRACE_LOADED=1
