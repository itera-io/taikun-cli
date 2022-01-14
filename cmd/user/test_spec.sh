Context 'user'

    setup() {
      oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name))
      ccid=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -o $oid -I)
      pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid -o $oid -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun project delete $pid
      taikun cloud-credential delete $ccid
      taikun organization delete $oid -q
    }
    AfterAll 'cleanup'

    add_user() {
      username=$(_rnd_name)
      uid=$(taikun user add $username --role user --email "${username}@mailinator.com" -o $oid -I)
    }
    BeforeEach 'add_user'

    del_user() {
      taikun user delete $uid -q
    }
    AfterEach 'del_user'

    Example 'add and then remove a user'
      When call taikun user list -o $oid --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include "$username"
    End

    Example 'duplicate name causes error'
      When call taikun user add $username --role manager --email "${username}2@mailinator.com" -o $oid
      The status should equal 1
      The stderr should include 'duplicate'
    End

    Example 'duplicate email causes error'
      When call taikun user add "${username}2" --role manager --email "${username}@mailinator.com" -o $oid
      The status should equal 1
      The stderr should include 'duplicate'
    End

    Context

      bind_user() {
        taikun user project bind $uid --project-id $pid -q
      }

      unbind_user() {
        taikun user project unbind $uid --project-id $pid -q
      }

      Example 
        When call taikun user project list $uid
        The status should equal 0
        The lines of output should equal 1
      End

    End

End
