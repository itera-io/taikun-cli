Context 'user'

    setup() {
      oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I)
      ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_APPLICATION_CREDENTIAL_ID" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -O "$oid" -I)
      pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun project delete "$pid" -q
      taikun cloud-credential delete "$ccid" -q
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Context
      add_user() {
        username="$(_rnd_name)"
        uid=$(taikun user add "$username" --email "${username}@mailinator.com" -O "$oid" -I)
      }
      BeforeAll 'add_user'

      del_user() {
        taikun user delete "$uid" -q
      }
      AfterAll 'del_user'

      Example 'add and then remove'
        When call taikun user list -O "$oid" --no-decorate
        The status should equal 0
        The lines of output should equal 1
        The output should include "$username"
      End

      Example 'duplicate name causes error'
        When call taikun user add "$username" --email "${username}2@mailinator.com" -O "$oid"
        The status should equal 1
        The stderr should include 'already exists'
        The stderr should include "$username"
      End

      Example 'duplicate email causes error'
        When call taikun user add "${username}2" --email "${username}@mailinator.com" -O "$oid"
        The status should equal 1
        The stderr should include 'already exists'
        The stderr should include "${username}@mailinator.com"
      End

      Context
        bind() {
          taikun user project bind "$uid" --project-id "$pid" -q
        }
        BeforeEach 'bind'

        unbind() {
          taikun user project unbind "$uid" --project-id "$pid" -q 2>/dev/null || true # TODO remove failsafe

        }
        AfterEach 'unbind'

        Example 'bind and unbind'
          When call taikun user project list "$uid" --no-decorate
          The status should equal 0
          The lines of output should equal 1
        End
      End
    End
End
