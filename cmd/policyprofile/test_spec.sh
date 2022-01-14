Context 'policyprofile'

  setup() {
    oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name) -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete $oid -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No policy profiles'
    When call taikun policy-profile list -o $oid --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context 'add/delete'
    add_config() {
      scid=$(taikun policy-profile add $(_rnd_name) -o $oid -I)
    }

    Before 'add_config'

    delete_config() {
      taikun policy-profile delete $scid -q
    }

    After 'delete_config'

    Example 'add then delete policy profile'
      When call taikun policy-profile list -o $oid --no-decorate
      The status should equal 0
      The lines of output should equal 1
    End
  End

End
