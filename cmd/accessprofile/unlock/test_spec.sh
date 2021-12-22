Context 'accessprofile/unlock'
  cleanup() {
    taikun access-profile delete $id
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context 'unlocked access profile'
    setup() {
      name=_rnd_name
      id=$(taikun access-profile create $name -I)
    }

    Example 'unlock an unlocked access profile'
      When call taikun access-profile unlock $id
      The stderr should include '400'
      The status should equal 1
    End
  End

  Context 'locked access profile'
    setup() {
      name=_rnd_name
      id=$(taikun access-profile create $name -I)
      taikun access-profile lock $id
    }

    Example 'unlock a locked access profile'
      When call taikun access-profile unlock $id
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End
End
