Context 'alertingprofile'

    setup() {
      oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Context
      add_profile() {
        name="$(_rnd_name)"
        pid=$(taikun alerting-profile add "$name" --reminder daily -o "$oid" -I | xargs)
      }
      BeforeEach 'add_profile'

      del_profile() {
        taikun alerting-profile delete "$pid" -q
      }
      AfterEach 'del_profile'

      Example 'add and then remove'
        When call taikun alerting-profile list -o "$oid" --columns name,reminder --no-decorate
        The status should equal 0
        The lines of output should equal 2 # counting the default
        The output should include "$name"
        The output should include "Daily"
      End

      Example 'duplicate name causes error'
        When call taikun alerting-profile add "$name" --reminder daily -o "$oid"
        The status should equal 1
        The stderr should include 'Please specify another name'
      End

      Example 'invalid reminder causes error'
        When call taikun alerting-profile add "$name" --reminder random -o "$oid"
        The status should equal 1
        The stderr should include 'reminder'
      End
    End

    Context 'lock and unlock'
      add_profile() {
        name="$(_rnd_name)"
        apid=$(taikun alerting-profile add "$name" --reminder daily -o "$oid" -I | xargs)
        taikun alerting-profile lock "$apid"
      }
      BeforeAll 'add_profile'

      del_profile() {
        taikun alerting-profile delete "$apid" -q
      }
      AfterAll 'del_profile'

      list_profile(){
        taikun alerting-profile list -o "$oid" --columns id,name,reminder,lock --no-decorate | grep "$pid"
      }

      Example 'list locked'
        When call list_profile
        The status should equal 0
        The lines of output should equal 2 # counting the default
        The output should include "$name"
        The output should include "Daily"
        The output should include "Locked"
        The output should include "Unlocked" # Default in Unlocked
      End

      Example 'lock with already locked'
        When call taikun alerting-profile lock "$apid"
        The status should equal 1
        The stderr should include "Alerting profile already lock"
      End

      Example 'unlock'
        When call taikun alerting-profile unlock "$apid"
        The status should equal 0
        The output should include "Operation was successful"
      End

      Example 'list unlocked'
        When call list_profile
        The status should equal 0
        The lines of output should equal 2 # counting the default
        The output should include "$name"
        The output should include "Daily"
        The output should include "Unlocked"
        The output should not include "Locked"
      End

    End
End
