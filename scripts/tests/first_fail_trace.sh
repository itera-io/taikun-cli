# scripts/first_fail_trace.sh
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

export -f first_fail_trace_on first_fail_trace_off trace_run
export FIRST_FAIL_TRACE_LOADED=1
