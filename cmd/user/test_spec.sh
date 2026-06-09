Context 'user'

    setup() {
      accountID=$(taikun accounts list --no-decorate | cut -d ' ' -f1 | xargs)
    }
    BeforeAll 'setup'

    Context
      add_user() {
        username="$(_rnd_name)"
        uid=$(taikun user add "$username" --email "${username}@mailinator.com" --account-id "$accountID" -I)
      }
      BeforeAll 'add_user'

      del_user() {
        taikun user delete "$uid" -q
      }
      AfterAll 'del_user'

      Example 'add and then remove'
        When call taikun user list --no-decorate
        The status should equal 0
        The output should include "$username"
      End

      Example 'duplicate name causes error'
        When call taikun user add "$username" --email "${username}2@mailinator.com" --account-id "$accountID"
        The status should equal 1
        The stderr should include 'already exists'
        The stderr should include "$username"
      End

      Example 'duplicate email causes error'
        When call taikun user add "${username}2" --email "${username}@mailinator.com" --account-id "$accountID"
        The status should equal 1
        The stderr should include 'already exists'
        The stderr should include "${username}@mailinator.com"
      End

    End
End
