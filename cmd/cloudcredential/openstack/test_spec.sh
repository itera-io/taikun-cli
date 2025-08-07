Context 'cloudcredential/openstack'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid=$(taikun cloud-credential openstack add "$ccname" -s "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    list_cc(){
      taikun cloud-credential list -o "$oid" --no-decorate
    }

    Example 'list openstack cloud credential'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'check invalid openstack cloud credential'
      When call taikun cloud-credential openstack check -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" -u "$OS_USERNAME" --url "mockup.local"
      The lines of stderr should equal 1
      The status should equal 1
      The stderr should include "Error: OpenStack cloud credential is not valid"
    End

End