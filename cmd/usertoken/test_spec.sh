Context 'usertoken/create'

    setup() {
        tokenname=$(_rnd_name)
    }

    BeforeAll setup

    login() {
        tokenout=$(taikun usertoken add "$tokenname" --format json)
        accesskey=$(echo "$tokenout" | jq .accessKey)
        secretkey=$(echo "$tokenout" | jq .secretKey)
        curl -X POST "https://$TAIKUN_API_HOST/api/v1/Auth/login" -H  "accept: application/json" -H  "Content-Type: application/*+json" -d "{  \"mode\": \"token\",  \"accessKey\": $accesskey,  \"secretKey\": $secretkey}"
    }

    delete() {
        taikun usertoken delete "$tokenname" -q
    }

    AfterAll delete

    Example 'basic usertoken'
        When call taikun usertoken add "$tokenname" --bind-all
        The status should equal 0
        The output should include 'ACCESS-KEY'
        The output should include 'SECRET-KEY'
    End

    Example 'name already exists'
        When call taikun usertoken add "$tokenname" --bind-all
        The status should equal 1
        The stderr should include 'already exists'
    End

    Example 'usertoken list'
        When call taikun usertoken list
        The status should equal 0
        The output should include 'ACCESS-KEY'
        The output should include "$tokenname"
    End

    Example 'delete usertoken'
        When call taikun usertoken delete "$tokenname"
        The status should equal 0
        The output should include 'deleted successfully'
    End

    Example 'delete nonexisting usertoken'
        When call taikun usertoken delete "$tokenname"
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

Context 'usertoken/binding'
      setup() {
          tokenname="$(_rnd_name)"
          taikun usertoken add "$tokenname" -q
          taikun usertoken bind "$tokenname" --endpoints Kubernetes/POST/cli -q
      }
      BeforeAll setup

      delete() {
          taikun usertoken delete "$tokenname" -q
      }
      AfterAll delete

      Example 'show bound endpoints of token'
        When call taikun usertoken show-endpoints "$tokenname"
        The status should equal 0
        The output should include 'DESCRIPTION'
        The output should include 'CONTROLLER'
        The output should include 'METHOD'
        The output should include 'PATH'
        The output should include 'ID'
        The output should include 'Kubernetes'
        The output should include 'POST'
        The output should include 'cli'
        The output should include 'Execute k8s web socket namespaced pod'
      End

      Example 'show unbound endpoints of token'
        When call taikun usertoken show-endpoints "$tokenname" --unbound
        The status should equal 0
        The output should include 'DESCRIPTION'
        The output should include 'CONTROLLER'
        The output should include 'METHOD'
        The output should include 'PATH'
        The output should include 'ID'
        The output should include 'GET'
        The output should not include 'Execute k8s web socket namespaced pod'
      End
End

Context 'usertoken/unbinding'
      setup() {
          tokenname="$(_rnd_name)"
          taikun usertoken add "$tokenname" --bind-all -q
          taikun usertoken unbind "$tokenname" --endpoints Kubernetes/POST/cli -q
      }
      BeforeAll setup

      delete() {
          taikun usertoken delete "$tokenname" -q
      }
      AfterAll delete

      Example 'check if token unbinded'
        When call taikun usertoken show-endpoints "$tokenname"
        The status should equal 0
        The output should include 'DESCRIPTION'
        The output should include 'Kubernetes'
        The output should include 'POST'
        The output should not include 'Execute k8s web socket namespaced pod'
      End

      Example 'bind unbound endpoint'
        When call taikun usertoken bind "$tokenname" --endpoints Kubernetes/POST/cli
        The status should equal 0
        The output should include 'Operation was successful.'
      End
End