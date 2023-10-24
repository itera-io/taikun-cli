Context 'billing/rule/list'

  setup() {
    cname="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun billing credential add -p "$PROMETHEUS_PASSWORD" -u "$PROMETHEUS_URL" -l "$PROMETHEUS_USERNAME" -o "$oid" "$cname" -I)

    flags="-b $cid -l foo=bar -m foo --price 1 --price-rate 5 -t count"
    name1="$(_rnd_name)"
    name2="$(_rnd_name)"
    id1=$(taikun billing rule add "$name1" $flags -I)
    id2=$(taikun billing rule add "$name2" $flags -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing rule delete "$id1" -q 2>/dev/null || true
    taikun billing rule delete "$id2" -q 2>/dev/null || true
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'list all billing rules'
    When call taikun billing rule list --no-decorate
    The status should equal 0
    The output should include "$name1"
    The output should include "$name2"
  End

  Example 'list only one billing rule'
    When call taikun billing rule list --no-decorate --limit 1
    The status should equal 0
    The lines of output should equal 1
  End

End
