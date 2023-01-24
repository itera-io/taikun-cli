Context 'usertoken'

    setup() {
        tokenname="test-cli-jvndkjbv"
    }

    BeforeAll setup

    login() {
        tokenout=$(taikun usertoken add $tokenname --format json)
        accesskey=$(echo $tokenout | jq .accessKey)
        secretkey=$(echo $tokenout | jq .secretKey)
        curl -X POST "https://api.taikun.dev/api/v1/Auth/login" -H  "accept: application/json" -H  "Content-Type: application/*+json" -d "{  \"mode\": \"token\",  \"accessKey\": $accesskey,  \"secretKey\": $secretkey}"
    }

    delete() {
        taikun usertoken delete $tokenname
    }

    AfterAll delete

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

    Example 'usertoken list'
        When call taikun usertoken list
        The status should equal 0
        The output should include 'ACCESS-KEY'
        The output should include $tokenname
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

    Example 'login'
        When call login
        The status should equal 0
        The output should include 'token'
        The output should include 'refreshToken'
        The output should include 'refreshTokenExpireTime'
        The stderr should include ""
    End

End