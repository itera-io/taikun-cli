Context 'billing/rule/organization/bind'

  setup() {
    name="$(_rnd_name)"
    pass="$PROMETHEUS_PASSWORD"
    url="$PROMETHEUS_URL"
    user="$PROMETHEUS_USERNAME"

    oid=$(taikun organization add "$name" --full-name "$name" -I | xargs)
    cid=$(taikun billing credential add "$name" -p "$pass" -u "$url" -l "$user" -o "$oid" -I | xargs)
    id=$(taikun billing rule add "$name" -b "$cid" -l foo=foo -m abc --price 1 --price-rate 1 --type count -I | xargs)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun billing rule delete "$id" -q 2> /dev/null || true
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Context
    bind_org() {
      taikun billing rule organization bind "$id" -o "$oid" -d 42 -q
    }

    Before 'bind_org'
    Example 'bind an organization'
      When call taikun billing rule organization list "$id" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include "$name"
      The output should include "$oid"
    End
  End

  Example 'bind a nonexistent organization'
    When call taikun billing rule organization bind "$id" -o 0 -d 42 -q
    The status should equal 1
    The stderr should include 'Failed to add organization(s) for prometheus rule'
    The stderr should include '400'
  End
End
