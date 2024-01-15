Context 'whoami'
  Example 'whoami does not fail'
    When call taikun whoami
    The status should equal 0
    The lines of output should equal 1
  End
End