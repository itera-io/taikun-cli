Context 'application/app'
  Example 'list one application'
    When call taikun application list --limit 1
    The status should equal 0
    The output should include "REPOSITORY"
    The lines of output should equal 3
  End

  Example 'list 10 applications'
    When call taikun application list --limit 10 --no-decorate
    The status should equal 0
    The lines of output should equal 10
  End
End
