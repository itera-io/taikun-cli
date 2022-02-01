Context 'accessprofile/add'
  setup() {
    name=$(_rnd_name)
  }

  cleanup() {
    taikun access-profile delete $id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'basic access profile'
    run() {
      id=$(taikun access-profile add $name -I)
      taikun access-profile list | grep $id
    }

    When call run
    The output should include "$name"
    The status should equal 0
  End

  Context
    add_access_profile() {
      id=$(taikun access-profile add $name -I)
    }
    Before 'add_access_profile'

    Example 'duplicate names'
      When call taikun access-profile add $name
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
