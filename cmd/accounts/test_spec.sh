Context 'accounts'
  setup(){
    if [ -z "$TAIKUN_EMAIL" ] && [ -z "$TAIKUN_ACCESS_KEY" ]; then
      Skip "No credentials found, skipping integration tests"
    fi

    account_name=$(_rnd_name)
  }
  BeforeAll setup

  clean_up(){
    if [ -n "$account_id" ]; then
      taikun accounts delete "$account_id" -q 2>/dev/null || true
    fi
  }
  AfterAll clean_up

  Example 'accounts create'
    When call taikun accounts create "$account_name" --email "${account_name}@itera.io"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts check'
    When call taikun accounts check "$account_name"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts list and find account_id'
    find_account(){
      account_id=$(taikun accounts list --limit 100 --columns id,name --no-decorate | grep "$account_name" | awk '{print $1}')
      export account_id
      echo "$account_id"
    }
    When call find_account
    The status should equal 0
    The output should not be empty
  End

  Example 'accounts details'
    get_details(){
      taikun accounts details "$account_id"
    }
    When call get_details
    The status should equal 0
    The output should include "$account_name"
    The output should include "$account_id"
  End

  Example 'accounts update'
    When call taikun accounts update "$account_id" --name "${account_name}-upd"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts enable-2fa'
    When call taikun accounts enable-2fa
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts disable-2fa'
    When call taikun accounts disable-2fa
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts add-admin'
    When call taikun accounts add-admin "$account_id" --user-id "some-user-id"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts transfer-ownership'
    When call taikun accounts transfer-ownership "some-user-id"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts organizations list'
    When call taikun accounts organizations list "$account_id"
    The status should equal 0
  End

  Example 'accounts users list'
    When call taikun accounts users list "$account_id"
    The status should equal 0
  End

  Example 'accounts projects list'
    When call taikun accounts projects list "$account_id"
    The status should equal 0
  End

  Example 'accounts groups list'
    When call taikun accounts groups list "$account_id"
    The status should equal 0
  End

  Example 'accounts delete'
    When call taikun accounts delete "$account_id"
    The status should equal 0
    The output should include "deleted successfully"
    unset account_id
  End
End
