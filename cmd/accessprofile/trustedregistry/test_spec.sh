# shellcheck shell=sh

Describe 'accessprofile/trustedregistry'
  setup() {
    apname="$(_rnd_name)"
    registryname="ghcr.io"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    apid=$(taikun access-profile add "$apname" -o "$oid" -I | xargs )
  }

  cleanup() {
    taikun access-profile delete "$apid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'Add trusted registry'
    When call taikun access-profile trusted-registry add "$apid" --registry "$registryname"
    The output should include "$registryname"
    The output should include "$apname"
    The status should equal 0
  End

  Context 'when trusted registry is added to access-profile'
    add_trusted_registry() {
      trid=$(taikun access-profile trusted-registry add "$apid" --registry "$registryname" -I | xargs)
    }
    BeforeEach 'add_trusted_registry'

    delete_trusted_registry() {
      taikun access-profile trusted-registry delete "$trid" -q 2>/dev/null || true
    }
    AfterEach 'delete_trusted_registry'


    Example 'List the trusted registry'
      When call taikun access-profile trusted-registry list "$apid" --no-decorate
      The output should include "$trid"
      The output should include "$registryname"
      The lines of output should equal 1
      The status should equal 0
    End

    Example 'Delete the trusted registry'
      When call taikun access-profile trusted-registry delete "$trid"
      The output should include "Trusted registry with ID $trid was deleted successfully."
      The lines of output should equal 1
      The status should equal 0
    End

  End

End
