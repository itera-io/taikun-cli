Context 'lock/unlock'
  setup() {
    orgname="$(_rnd_name)"
    ccname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
    ccid=$(taikun cloud-credential openstack add "$ccname" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    taikun cloud-credential lock "$ccid" -q
  }
  BeforeAll 'setup'

  cleanup() {
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'list if lock is successful'
    When call taikun cloud-credential list -o "$oid"
    The status should equal 0
    The output should include "Locked"
  End

  Example 'lock with already locked'
    When call taikun cloud-credential lock "$ccid"
    The status should equal 1
    The stderr should include "Cloud credential already lock"
  End

  Example 'unlock'
    When call taikun cloud-credential unlock "$ccid"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'list if unlock is successful'
    When call taikun cloud-credential list -o "$oid"
    The status should equal 0
    The output should include "Unlocked"
  End

End