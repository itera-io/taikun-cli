Context 'alertingprofile/integration'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    name="$(_rnd_name)"
    apid=$(taikun alerting-profile add "$name" --reminder daily -o "$oid" -I | xargs)
    integration_id=$(taikun alerting-profile integration add "$apid" --token "mock" --url "https://zaphod.beeblebrox.local" --type opsgenie -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun alerting-profile delete "$apid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_integrations(){
    taikun alerting-profile integration list "$apid" --no-decorate
  }

  Example 'list one integration'
    When call list_integrations
    The status should equal 0
    The lines of output should equal 1
    The output should include "zaphod.beeblebrox.local"
    The output should include "$integration_id"
  End

  Example 'add integration without token'
    When call taikun alerting-profile integration add "$apid" --url "https://arthur.dent.local" --type opsgenie
    The status should equal 1
    The stderr should include "token"
  End

  Example 'add integration with url'
    When call taikun alerting-profile integration add "$apid" --token "mock" --url "https://arthur.dent.local" --type opsgenie
    The lines of output should equal 2
    The output should include "arthur.dent.local"
    The output should include "URL"
    The output should include "ID"
  End

  Example 'list one integration'
    When call list_integrations
    The status should equal 0
    The lines of output should equal 2
    The output should include "zaphod.beeblebrox.local"
    The output should include "$integration_id"
    The output should include "arthur.dent.local"
  End

  Example 'delete integration'
    When call taikun alerting-profile integration delete "$integration_id"
    The status should equal 0
    The output should include "Alerting integration with ID $integration_id was deleted successfully."
  End

  Example 'list one integration'
    When call list_integrations
    The status should equal 0
    The lines of output should equal 1
    The output should include "arthur.dent.local"
    The output should not include "zaphod.beeblebrox.local"
    The output should not include "$integration_id"
  End

End