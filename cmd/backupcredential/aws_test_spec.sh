Context 'backupcredential /credential'
  setup(){
    cname="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun backup-credential add "$cname" --s3-access-key "$AWS_ACCESS_KEY_ID" --s3-endpoint "$AWS_S3_ENDPOINT" --s3-region "$AWS_DEFAULT_REGION" --s3-secret-key "$AWS_SECRET_ACCESS_KEY" -o "$oid" -I | xargs)
    taikun backup-credential lock "$cid" -q
  }
  BeforeAll 'setup'

  cleanup() {
    taikun backup-credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_cred(){
    taikun backup-credential list | grep "$cname"
  }

  Example 'lock already locked'
    When call taikun backup-credential lock "$cid"
    The status should equal 1
    The stderr should include "already lock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "Locked"
    The output should include "$cname"
  End

  Example 'unlock locked'
    When call taikun backup-credential unlock "$cid"
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'unlock already unlocked'
    When call taikun backup-credential unlock "$cid"
    The status should equal 1
    The stderr should include "already unlock"
  End

  Example 'list locked credential'
    When call list_cred
    The status should equal 0
    The output should include "$cid"
    The output should include "Unlocked"
    The output should include "$cname"
  End

End
