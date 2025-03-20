Context 'cloudcredential/openstack-appcred'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid=$(taikun cloud-credential openstack add "$ccname" -o "$oid" -i "$OS_APPLICATION_CREDENTIAL_ID" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME"  --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
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

    Example 'list openstack cloud credential appcred'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'check valid openstack cloud credential appcred'
      When call taikun cloud-credential openstack check -i "$OS_APPLICATION_CREDENTIAL_ID" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --url "$OS_AUTH_URL"
      The lines of output should equal 1
      The status should equal 0
      The output should include "OpenStack cloud credential is valid."
    End

    Example 'check invalid openstack cloud credential appcred'
      When call taikun cloud-credential openstack check -i "$OS_APPLICATION_CREDENTIAL_ID" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --url "mockup.local"
      The lines of stderr should equal 1
      The status should equal 1
      The stderr should include "OpenStack cloud credential is not valid."
    End

End
