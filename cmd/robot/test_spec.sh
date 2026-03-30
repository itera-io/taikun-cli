

Context 'robot'
  setup(){
    # Skip if TAIKUN_EMAIL is not set (not in CI or no credentials provided)
    if [ -z "$TAIKUN_EMAIL" ] && [ -z "$TAIKUN_ACCESS_KEY" ]; then
      Skip "No credentials found, skipping integration tests"
    fi

    org_id=$(taikun organization list --limit 1 --columns id --no-decorate | head -n 1)
    account_id=$(taikun accounts list --limit 1 --columns id --no-decorate | head -n 1)
    robot_name=$(_rnd_name)
    expires_at="2030-01-01T00:00:00Z"
    
    if [ -z "$org_id" ] || [ -z "$account_id" ]; then
      Skip "No organization or account found"
    fi
  }
  BeforeAll setup

  clean_up(){
    if [ -n "$robot_id" ]; then
      taikun robot delete "$robot_id" -q 2>/dev/null || true
    fi
  }
  AfterAll clean_up

  Example 'robot create'
    When call taikun robot create "$robot_name" --organization-id "$org_id" --account-id "$account_id" --description "test robot" --expires-at "$expires_at"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'robot list and find robot_id'
    find_robot(){
      robot_id=$(taikun robot list "$account_id" --organization-id "$org_id" --search "$robot_name" --columns userid --no-decorate | head -n 1)
      export robot_id
      echo "$robot_id"
    }
    When call find_robot
    The status should equal 0
    The output should not be empty
  End

  Example 'robot details'
    get_details(){
      taikun robot details "$robot_id"
    }
    When call get_details
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
    When call taikun robot update "$robot_id" --description "updated description"
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

  Example 'robot regenerate'
    When call taikun robot regenerate "$robot_id"
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
