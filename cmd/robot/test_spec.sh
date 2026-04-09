Context 'robot'
  setup(){
    # Skip if TAIKUN_EMAIL is not set (not in CI or no credentials provided)
    if [ -z "$TAIKUN_EMAIL" ] && [ -z "$TAIKUN_ACCESS_KEY" ]; then
      Skip "No credentials found, skipping integration tests"
    fi

    ## Create account
    account_name=$(_rnd_name)
    taikun accounts create "$account_name" --email "${account_name}@itera.io" -q
    account_id=$(taikun accounts list --limit 100 --columns id,name --no-decorate | grep "$account_name" | awk '{print $1}')

    ## Create organization
    org_id=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" --account-id "$account_id" -I | xargs )

    ## Create Robot
    robot_name=$(_rnd_name)
    echo $robot_name

    ## Getting date of expiration 5 months ahead of now
    expires_at=$(date -u -d "+5 months" +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || date -u -v+5m +"%Y-%m-%dT%H:%M:%SZ")
    
    if [ -z "$org_id" ] || [ -z "$account_id" ]; then
      Skip "No organization or account found"
    fi
  }
  BeforeAll setup

  clean_up(){
    if [ -n "$robot_id" ]; then
      taikun robot delete "$robot_id" -q 2>/dev/null || true
    fi
    if [ -n "$org_id" ]; then
      taikun organization delete "$org_id" -q 2>/dev/null || true
    fi
    if [ -n "$account_id" ]; then
      taikun accounts delete "$account_id" -q 2>/dev/null || true
    fi
  }
  AfterAll clean_up

  Example 'robot create'
    When call taikun robot create "$robot_name" --organization-id "$org_id" --account-id "$account_id" --description "test-robot" --expires-at "$expires_at" --scope "scope:kubernetes-profiles:read"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'robot list'
    When call taikun robot list "$account_id" --organization-id "$org_id" --no-decorate
    The status should equal 0
    The output should be present
  End

  robot_id=$(taikun robot list "$account_id" --organization-id "$org_id"  --no-decorate | xargs | awk '{print $1}')

  Example 'robot details'
    When call taikun robot details
    The status should equal 0
    The output should include "$robot_name"
    The output should include "$robot_id"
  End

  Example 'robot status deactivate'
    When call taikun robot status "$robot_id" --mode deactivate
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'robot status activate'
    When call taikun robot status "$robot_id" --mode activate
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'robot update'
    When call taikun robot update "$robot_id" --description "updated-description" --name "$robot_name"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'robot scope-list'
    When call taikun robot scope-list
    The status should equal 0
    The output should include "KEY"
  End

  Example 'robot update-scope'
    When call taikun robot update-scope "$robot_id" --scope "scope:kubernetes-profiles:read"
    The status should equal 0
    The output should include "Operation was successful"
  End

  new_expires_at=$(date -u -d "+5 months" +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || date -u -v+5m +"%Y-%m-%dT%H:%M:%SZ")

  Example 'robot regenerate'
    When call taikun robot regenerate "$robot_id" --expires_at "$new_expires_at"
    The status should equal 0
    The output should include "ACCESS_KEY"
    The output should include "SECRET_KEY"
  End

  Example 'robot delete'
    When call taikun robot delete "$robot_id"
    The status should equal 0
    The output should include "deleted successfully"
    unset robot_id
  End
End
