Context 'project/set/expiration'
    setup() {
      oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs )
      ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_APPLICATION_CREDENTIAL_ID" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
      projectname="$(_rnd_name)"
      pid=$(taikun project add "$projectname" --cloud-credential-id "$ccid" -I | xargs)
    }
    BeforeAll 'setup'

    cleanup() {
      if ! taikun project delete "$pid" -q 2>/dev/null; then
        taikun project delete --force "$pid" -q 2>/dev/null || true
      fi
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    list_project(){
      taikun project list -o "$oid" -C id,name,org,expires,delete-on-expiration --no-decorate
    }

    Example 'list no expiration'
      When call list_project
      The status should equal 0
      The lines of output should equal 1
      The output should include "$pid"
      The output should include "No"
    End

    Example 'set expiration'
      When call taikun project set expiration "$pid" --expiration-date 01.01.3000 -d
      The status should equal 0
      The lines of output should equal 1
      The output should include "Operation was successful"
    End

    Example 'list set expiration'
      When call list_project
      The status should equal 0
      The lines of output should equal 1
      The output should include "$pid"
      The output should include "Yes"
      The output should include "3000-01-01 00:00:00"
    End

    Example 'remove expiration'
      When call taikun project set expiration "$pid" --remove-expiration
      The status should equal 0
      The lines of output should equal 1
      The output should include "Operation was successful"
    End

    Example 'list no expiration'
      When call list_project
      The status should equal 0
      The lines of output should equal 1
      The output should include "$pid"
      The output should include "No"
    End

End