Context 'user/info'
    Example 'get info about current user'
      When call taikun user info
      The status should equal 0
      The output should include "ID"
      The output should include "USERNAME"
      The output should include "EMAIL"
      The output should include "ROLE"
    End
End
