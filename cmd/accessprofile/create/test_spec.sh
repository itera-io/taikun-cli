Context 'accessprofile/create'
  setup() {
    name=$(_rnd_name)
  }

  cleanup() {
    if [[ -n $id ]]; then
      taikun access-profile delete $id -q || true
    fi
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'basic access profile'
    run() {
      id=$(taikun access-profile create $name -I)
      taikun access-profile list | grep $id
    }

    When call run
    The output should include "$name"
    The status should equal 0
  End

  Context
    create_access_profile() {
      id=$(taikun access-profile create $name -I)
    }
    Before 'create_access_profile'

    Example 'duplicate names'
      When call taikun access-profile create $name
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
