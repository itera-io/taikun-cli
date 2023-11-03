Context 'project/lockmanager'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    pid=$(taikun project add "$(_rnd_name)" -o "$oid" --cloud-credential-id "$ccid" -I)
    taikun project lock "$pid" -q
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

  Example 'lock with already locked'
    When call taikun project lock "$pid"
    The status should equal 1
    The stderr should include "Project already loc"
  End

  Example 'unlock'
    When call taikun project unlock "$pid"
    The status should equal 0
    The output should include "Operation was successful"
  End

End