Context 'user/info'
      setup() {
        oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -o "$oid" -I)
        username="$(_rnd_name)"
        uid=$(taikun user add "$username" --role user --email "${username}@mailinator.com" -o "$oid" -I)
      }
      BeforeAll 'setup'

      cleanup() {
        taikun user delete "$uid" -q
        taikun project delete "$pid" -q
        taikun cloud-credential delete "$ccid" -q
        taikun organization delete "$oid" -q 2>/dev/null || true
      }
      AfterAll 'cleanup'

      Example 'get info about user'
        When call taikun user info "$uid"
        The status should equal 0
        The lines of output should equal 17
        The output should include "$uid"
        The output should include "$username"
        The output should include "mailinator.com"
        The output should include "MUST-RESET-PASSWORD"
      End
End