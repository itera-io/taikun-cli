Context 'alertingprofile/webhook'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    name="$(_rnd_name)"
    apid=$(taikun alerting-profile add "$name" --reminder daily -o "$oid" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun alerting-profile delete "$apid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_webhooks(){
    taikun alerting-profile webhook list "$apid" --no-decorate
  }

  Example 'list no webhooks'
    When call list_webhooks
    The status should equal 0
    The lines of output should equal 0
  End

  Example 'add webhook without url'
    When call taikun alerting-profile webhook add "$apid"
    The status should equal 1
    The stderr should include "url"
  End

  Example 'add webhook with url'
    When call taikun alerting-profile webhook add "$apid" -u zaphod.beeblebrox.local
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list one webhook'
    When call list_webhooks
    The status should equal 0
    The lines of output should equal 1
    The output should include "zaphod.beeblebrox.local"
  End

  Example 'clear webhooks'
    When call taikun alerting-profile webhook clear "$apid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list no webhooks'
    When call list_webhooks
    The status should equal 0
    The lines of output should equal 0
  End

End