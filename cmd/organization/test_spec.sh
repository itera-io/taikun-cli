Context 'organization'
    add_org(){
      orgname=$(_rnd_name)
      orgnamefull=$(_rnd_name)
      oid=$(taikun organization add $orgname --full-name $orgnamefull -I)
      fakeoid=$(_rnd_number)
    }
    BeforeAll add_org

    del_org(){
      taikun organization delete $oid
    }
    AfterAll del_org

    Example 'add, list, and remove'
      When call taikun organization list
      The status should equal 0
      The output should include "$orgname"
      The output should include "$orgnamefull"
      The output should include "$oid"
      The output should not include "$fakeoid"
    End

    Example 'add, info, and remove'
      When call taikun organization info "$oid"
      The status should equal 0
      The output should include "$orgname"
      The output should include "$orgnamefull"
      The output should include "$oid"
      The output should include "CREATED-AT"
      The output should not include "$fakeoid"
    End

    Example 'Info on bad organization id'
      When call taikun organization info "$fakeoid"
      The status should equal 1
      The stderr should include "not found"
      The stderr should include "$fakeoid"
      The stderr should not include "$orgname"
    End
End
