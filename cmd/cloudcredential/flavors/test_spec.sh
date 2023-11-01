Context 'cloudcredential/flavors'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid_openstack=$(taikun cloud-credential openstack add "$ccname" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
      ccid_aws=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun cloud-credential delete "$ccid_aws" -q 2>/dev/null || true
      taikun cloud-credential delete "$ccid_openstack" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Example 'List flavors for openstack'
      When call taikun cloud-credential flavors "$ccid_openstack"
      The status should equal 0
      The lines of output should not equal 0
      The lines of output should not equal 1
      The lines of output should not equal 2
      The output should include "NAME"
      The output should include "CPU"
      The output should include "RAM"
    End

    Example 'List flavors for aws'
      When call taikun cloud-credential flavors "$ccid_aws"
      The status should equal 0
      The lines of output should not equal 0
      The lines of output should not equal 1
      The lines of output should not equal 2
      The output should include "NAME"
      The output should include "CPU"
      The output should include "RAM"
    End

End