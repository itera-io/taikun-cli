Context 'accessprofile/create'
  setup() {
    name=$(_rnd_name)
  }

  cleanup() {
    if [[ -n $id ]]; then
      taikun access-profile delete $id || true
    fi
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'create basic access profile'

    run() {
      id=$(taikun access-profile create $name -I)
      taikun access-profile list | grep $id
    }

    When call run
    The output should include "$name"
    The status should equal 0
  End

End
