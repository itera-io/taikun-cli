Context 'billing/credential'
  setup(){
    cname="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun billing credential add "$cname" -p "$PROMETHEUS_PASSWORD" -u "$PROMETHEUS_URL" -l "$PROMETHEUS_USERNAME" -o "$oid" -I | xargs)
    taikun billing credential lock "$cid" -q
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_cred(){
    taikun billing credential list | grep "$cid"
  }

  Example 'lock already locked'
    When call taikun billing credential lock "$cid"
    The status should equal 1
    The stderr should include "already lock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "Locked"
    The output should include "$cname"
  End

  Example 'unlock locked'
    When call taikun billing credential unlock "$cid"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'unlock already unlocked'
    When call taikun billing credential unlock "$cid"
    The status should equal 1
    The stderr should include "already unlock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "Unlocked"
    The output should include "$cname"
  End

End
