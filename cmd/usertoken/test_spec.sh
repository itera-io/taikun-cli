Context 'usertoken'

    setup() {
        tokenname="test-cli-jvndkjbv"
    }

    BeforeAll setup

    Example 'create a basic user token'
        When call taikun usertoken add $tokenname --bind-all
        The status should equal 0
      The output should include 'ACCESS-KEY'
      The output should include 'SECRET-KEY'
    End

    Example 'create a user token with name already existing'
        When call taikun usertoken add $tokenname --bind-all
        The status should equal 1
        The stderr should include 'already exists'
    End

    Example 'delete one user token'
        When call taikun usertoken delete $tokenname
        The status should equal 0
        The output should include 'deleted successfully'
    End

    Example 'delete non existing user token'
        When call taikun usertoken delete $tokenname
        The status should equal 1
        The stderr should include "No user token found with name '$tokenname'"
    End


End