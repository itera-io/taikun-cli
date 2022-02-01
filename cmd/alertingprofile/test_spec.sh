Context 'alertingprofile'

    setup() {
      oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name) -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun organization delete $oid -q
    }
    AfterAll 'cleanup'

    Context
      add_profile() {
        name=$(_rnd_name)
        pid=$(taikun alerting-profile add $name --reminder daily -o $oid -I)
      }
      BeforeEach 'add_profile'

      del_profile() {
        taikun alerting-profile delete $pid -q
      }
      AfterEach 'del_profile'

      Example 'add and then remove'
        When call taikun alerting-profile list -o $oid --columns name,reminder --no-decorate
        The status should equal 0
        The lines of output should equal 2 # counting the default
        The output should include "$name"
        The output should include "Daily"
      End

      Example 'duplicate name causes error'
        When call taikun alerting-profile add $name --reminder daily -o $oid
        The status should equal 1
        The stderr should include 'Please specify another name'
      End

      Example 'invalid reminder causes error'
        When call taikun alerting-profile add $name --reminder random -o $oid
        The status should equal 1
        The stderr should include 'reminder'
      End
    End

End
