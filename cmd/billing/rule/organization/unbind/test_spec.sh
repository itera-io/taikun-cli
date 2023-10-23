Context 'billing/rule/organization/unbind'

  setup() {
    name="$(_rnd_name)"
    pass="$PROMETHEUS_PASSWORD"
    url="$PROMETHEUS_URL"
    user="$PROMETHEUS_USERNAME"

    oid=$(taikun organization add "$name" --full-name "$name" -I)
    cid=$(taikun billing credential add "$name" -p "$pass" -u "$url" -l "$user" -o "$oid" -I)
    id=$(taikun billing rule add "$name" -b "$cid" -l foo=foo -m abc --price 1 --price-rate 1 --type count -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun billing rule delete "$id" -q 2> /dev/null || true
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Context
    bind_unbind_org() {
      taikun billing rule organization bind "$id" -o "$oid" -d 42 -q
      taikun billing rule organization unbind "$id" -o "$oid" -q
    }

    Before 'bind_unbind_org'

    Example 'unbind an organization'
      When call taikun billing rule organization list "$id" --no-decorate
      The status should equal 0
      The lines of output should equal 0
    End
  End

  Example 'unbind a nonexistent organization'
    When call taikun billing rule organization unbind "$id" -o 0 -q
    The status should equal 1
    The stderr should include 'Can not find organization'
    The stderr should include '400'
  End
End
