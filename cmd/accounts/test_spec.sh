Context 'accounts'
  setup(){
    if [ -z "$TAIKUN_EMAIL" ] && [ -z "$TAIKUN_ACCESS_KEY" ]; then
      Skip "No credentials found, skipping integration tests"
    fi

    account_name=$(_rnd_name)
  }
  BeforeAll setup

  clean_up(){
    if [ -n "$uid" ]; then
      taikun user delete "$uid" -q 2>/dev/null || true
    fi
    if [ -n "$groupID" ]; then
      taikun groups delete "$groupID" -q 2>/dev/null || true
    fi
    if [ -n "$org_id" ]; then
      taikun organization delete "$org_id" -q 2>/dev/null || true
    fi
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
    The status should equal 1
    The error should include "already exists"
  End

  Example 'accounts check non-existing account'
    When call taikun accounts check "$account_name-sdfgh"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts list'
    When call taikun accounts list --limit 100 --columns id,name --no-decorate
    The status should equal 0
    The output should be present
  End

  Example 'find account_id'
    find_account(){
      account_id=$(taikun accounts list --limit 100 --columns id,name --no-decorate | grep "$account_name" | awk '{print $1}')
      export account_id
      echo "$account_id"
    }
    When call find_account
    The output should be present
  End

  account_id=$(taikun accounts list --limit 100 --columns id,name --no-decorate | grep "$account_name" | awk '{print $1}')
  ## Create organization
  org_id=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" --account-id "$account_id" -I | xargs )
  group_name=$(_rnd_name)
  taikun groups create "$group_name" --account-id "$account_id" -q
  groupID=`taikun groups list "$account_id" | grep "$group_name" | xargs | awk '{print $1}'`

  Example 'accounts details'
    When call taikun accounts details "$account_id"
    The status should equal 0
    The output should include "$account_name"
    The output should include "$account_id"
  End

  Example 'accounts update'
    When call taikun accounts update "$account_id" --name "${account_name}-upd"
    The status should equal 0
    The output should include "Operation was successful"
  End

## FIXME: 2FA tests are hard to implement due to dependency on the external app
#  Example 'accounts enable-2fa'
#    When call taikun accounts enable-2fa
#    The status should equal 0
#    The output should include "Operation was successful"
#  End

#  Example 'accounts disable-2fa'
#    When call taikun accounts disable-2fa
#    The status should equal 0
#    The output should include "Operation was successful"
#  End

  username1="$(_rnd_name)"
  uid1=$(taikun user add "$username1" --email "${username1}@mailinator.com" --account-id "$account_id" -I)
  username2="$(_rnd_name)"
  uid2=$(taikun user add "$username2" --email "${username2}@mailinator.com" --account-id "$account_id" -I)

  Example 'accounts add-admin'
    When call taikun accounts add-admin "$account_id" --user-id "$uid1"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts transfer-ownership'
    When call taikun accounts transfer-ownership "$uid2"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'accounts organizations list'
    When call taikun accounts organizations list "$account_id" --limit 100 --no-decorate
    The status should equal 0
    The output lines should equal 1 # Only one organization was created under same account
  End

  Example 'accounts users list'
    When call taikun accounts users list "$account_id" --limit 100 --no-decorate
    The status should equal 0
    The output lines should equal 2 # Two users were created under same account
  End

  Example 'accounts projects list'
    When call taikun accounts projects list "$account_id" --limit 100 --no-decorate
    The status should equal 0
    The output lines should equal 0 # No projects were created under this account
  End

  Example 'accounts groups list'
    When call taikun accounts groups list "$account_id" --limit 100 --no-decorate
    The status should equal 0
    The output lines should equal 1 # Only one group was created under the same account
  End
End
