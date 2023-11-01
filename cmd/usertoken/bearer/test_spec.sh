Context 'usertoken/bearer'
  Example 'Get a usertoken'
    When call taikun usertoken get-bearer
    The status should equal 0
    The lines of output should equal 1
    The output should include "Bearer "
  End
End