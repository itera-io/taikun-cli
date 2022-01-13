Context 'kubernetesprofile'

  setup() {
    oid=$(taikun organization create $(_rnd_name) --full-name $(_rnd_name) -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete $oid -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No kubernetes profiles'
    When call taikun kubernetes-profile list -o $oid --no-decorate
    The status should equal 0
    The lines of output should equal 1 # counting the default
  End

  Context 'add/delete'
    add_config() {
      scid=$(taikun kubernetes-profile add $(_rnd_name) -o $oid -I)
    }

    Before 'add_config'

    delete_config() {
      taikun kubernetes-profile delete $scid -q
    }

    After 'delete_config'

    Example 'add then delete kubernetes profile'
      When call taikun kubernetes-profile list -o $oid --no-decorate
      The status should equal 0
      The lines of output should equal 2 # counting the default
    End
  End

End
