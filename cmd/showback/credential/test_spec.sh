Context 'showback/credential'
  setup(){
    cname="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun showback credential add "$cname" -p "$PROMETHEUS_PASSWORD" -u "$PROMETHEUS_URL" -l "$PROMETHEUS_USERNAME" -o "$oid" -I | xargs)
    taikun showback credential lock "$cid"
  }
  BeforeAll 'setup'

  cleanup() {
    taikun showback credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_cred(){
    taikun showback credential list | grep "$cname"
  }

  Example 'lock already locked'
    When call taikun showback credential lock "$cid"
    The status should equal 1
    The stderr should include "already lock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "Yes"
    The output should include "$PROMETHEUS_USERNAME"
    The output should include "$cname"
  End

  Example 'unlock locked'
    When call taikun showback credential unlock "$cid"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'unlock already unlocked'
    When call taikun showback credential unlock "$cid"
    The status should equal 1
    The stderr should include "already unlock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "No"
    The output should include "$PROMETHEUS_USERNAME"
    The output should include "$cname"
  End

End
