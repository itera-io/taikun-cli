Context 'policyprofile'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No policy profiles'
    When call taikun policy-profile list -o "$oid" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context 'add/delete'
    add_config() {
      ppid=$(taikun policy-profile add "$(_rnd_name)" -o "$oid" -I)
    }

    Before 'add_config'

    delete_config() {
      taikun policy-profile delete "$ppid" -q
    }

    After 'delete_config'

    Example 'add then delete policy profile'
      When call taikun policy-profile list -o "$oid" --no-decorate
      The status should equal 0
      The lines of output should equal 1
    End
  End

  Context 'lock/unlock'
    add_config() {
      ppid=$(taikun policy-profile add "$(_rnd_name)" -o "$oid" -I)
      taikun policy-profile lock "$ppid" -q
    }

    BeforeAll 'add_config'

    delete_config() {
      taikun policy-profile delete "$ppid" -q
    }

    AfterAll 'delete_config'

    Example 'lock with already locked'
      When call taikun policy-profile lock "$ppid"
      The status should equal 1
      The stderr should include "Opa profile already lock"
    End

    Example 'unlock'
      When call taikun policy-profile unlock "$ppid"
      The status should equal 0
      The output should include "Operation was successful"
    End

  End


End
