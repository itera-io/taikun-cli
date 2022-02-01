Context 'accessprofile/unlock'
  cleanup() {
    taikun access-profile delete $id -q
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context
    setup() {
      name=$(_rnd_name)
      id=$(taikun access-profile add $name -I)
    }

    Example 'unlocked access profile'
      When call taikun access-profile unlock $id
      The stderr should include '400'
      The status should equal 1
    End
  End

  Context
    setup() {
      name=$(_rnd_name)
      id=$(taikun access-profile add $name -I)
      taikun access-profile lock $id -q
    }

    Example 'locked access profile'
      When call taikun access-profile unlock $id
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End
End
