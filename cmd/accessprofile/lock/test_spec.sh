Context 'accessprofile/lock'
  cleanup() {
    taikun access-profile unlock $id
    taikun access-profile delete $id
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context 'unlocked access profile'
    setup() {
      id=$(taikun access-profile create $RSC_PREFIX-unlocked | jq ".id" | tr -d '"')
    }

    Example 'lock an unlocked access profile'
      When call taikun access-profile lock $id
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End

  Context 'locked access profile'
    setup() {
      id=$(taikun access-profile create $RSC_PREFIX-unlocked | jq ".id" | tr -d '"')
      taikun access-profile lock $id
    }

    Example 'lock a locked access profile'
      When call taikun access-profile lock $id
      The stderr should include '400'
      The status should equal 1
    End
  End
End
