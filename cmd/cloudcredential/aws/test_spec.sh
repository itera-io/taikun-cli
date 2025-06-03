Context 'cloudcredential/aws'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid=$(taikun cloud-credential aws add "$ccname" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I)
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

    Example 'list aws cloud credential'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'check valid aws cloud credential'
      When call taikun cloud-credential aws check -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION"
      The lines of output should equal 1
      The status should equal 0
      The output should include "AWS cloud credential is valid."
    End

    Example 'check invalid aws cloud credential'
      When call taikun cloud-credential aws check -a "$AWS_ACCESS_KEY_ID" -s "mockup" -r "$AWS_DEFAULT_REGION"
      The lines of stderr should equal 1
      The status should equal 1
      The stderr should include "Error: AWS cloud credential is not valid"
    End

End