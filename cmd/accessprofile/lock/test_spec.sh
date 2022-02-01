Context 'accessprofile/lock'
  cleanup() {
    taikun access-profile unlock $id -q
    taikun access-profile delete $id -q
  }

  add_profile() {
    name=$(_rnd_name)
    id=$(taikun access-profile add $name -I)
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context
    setup() {
      add_profile
    }

    Example 'unlocked access profile'
      When call taikun access-profile lock $id
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End

  Context
    setup() {
      add_profile
      taikun access-profile lock $id -q
    }

    Example 'locked access profile'
      When call taikun access-profile lock $id
      The stderr should include '400'
      The status should equal 1
    End
  End
End
