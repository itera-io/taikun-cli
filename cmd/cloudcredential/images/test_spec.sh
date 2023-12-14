Context 'cloudcredential/images'
  setup() {
    orgname="$(_rnd_name)"
    ccname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
    ccid_openstack=$(taikun cloud-credential openstack add "$ccname" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    ccid_aws=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I | xargs)
    ccid_azure=$(taikun cloud-credential azure add "$(_rnd_name)" --az-count "$AZ_COUNT" --client-id "$AZURE_CLIENT_ID" --client-secret "$AZURE_SECRET" --location "$AZURE_LOCATION" --subscription-id "$AZURE_SUBSCRIPTION" --tenant-id "$AZURE_TENANT" -o "$oid" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun cloud-credential delete "$ccid_azure" -q 2>/dev/null || true
    taikun cloud-credential delete "$ccid_aws" -q 2>/dev/null || true
    taikun cloud-credential delete "$ccid_openstack" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'List images for openstack'
    When call taikun cloud-credential images "$ccid_openstack" --limit 3
    The status should equal 0
    The lines of output should not equal 0
    The lines of output should not equal 1
    The lines of output should not equal 2
    The lines of output should equal 5
    The output should include "NAME"
    The output should include "ID"
  End

  Example 'List images for aws'
    When call taikun cloud-credential images "$ccid_aws" --limit 3
    The status should equal 0
    The lines of output should not equal 0
    The lines of output should not equal 1
    The lines of output should not equal 2
    The lines of output should equal 5
    The output should include "NAME"
    The output should include "ID"
  End

    Example 'List images for azure'
      When call taikun cloud-credential images "$ccid_azure" --limit 3
      The status should equal 0
      The lines of output should not equal 0
      The lines of output should not equal 1
      The lines of output should not equal 2
      The lines of output should equal 5
      The output should include "NAME"
      The output should include "ID"
    End
End