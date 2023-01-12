Context 'usertoken'

    Example 'create a basic user token'
        When call taikun usertoken add $(_rnd_name)
        The status should equal 0
    End

    Example 'create a user token with name already existing'
        When call taikun usertoken add $(_rnd_name)
        The status should equal 1
        The stderr should include 'already exists'
    End

    Example 'delete one user token'
        When call taikun usertoken delete $(_rnd_name)
        The status should equal 0
    End

    Example 'delete non existing user token'
        When call taikun usertoken delete $(_rnd_name)
        The status should equal 1
        The stderr should include "No user token found with name '$(_rnd_name)'"
    End


End