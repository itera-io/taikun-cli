Context 'accessprofile/lock'
  cleanup() {
    taikun access-profile unlock $id
    taikun access-profile delete $id
  }

  create_profile() {
    name=$(_rnd_name)
    id=$(taikun access-profile create $name -I)
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context 'unlocked access profile'
    setup() {
      create_profile
    }

    Example 'lock an unlocked access profile'
      When call taikun access-profile lock $id
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End

  Context 'locked access profile'
    setup() {
      create_profile
      taikun access-profile lock $id
    }

    Example 'lock a locked access profile'
      When call taikun access-profile lock $id
      The stderr should include '400'
      The status should equal 1
    End
  End
End
