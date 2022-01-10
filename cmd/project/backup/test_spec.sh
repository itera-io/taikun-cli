Context 'project/backup'

  setup() {
    bid=$(taikun backup-credential create $(_rnd_name) -a $AWS_ACCESS_KEY_ID -e $S3_ENDPOINT -r $S3_REGION -s $AWS_SECRET_ACCESS_KEY -I)
    ccid=$(taikun cloud-credential openstack create $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    pid=$(taikun project create $(_rnd_name) --cloud-credential-id $ccid -I)
  }

  BeforeEach 'setup'

  cleanup() {
    if ! taikun project delete $pid -q 2>/dev/null; then
      taikun project delete --force $pid -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $ccid -q 2>/dev/null || true
    taikun backup-credential delete $bid -q 2>/dev/null || true
  }

  AfterEach 'cleanup'

  get_backup_status() {
    taikun project info $1 --no-decorate | grep -i backup | tr -d ' ' -f 2
  }

  Context
    enable_backup() {
      taikun project backup enable $pid -b $bid -q
    }
    Before 'enable_backup'

    Example 'enable backup'
      When call get_backup_status $pid
      The status should equal 0
      The output should equal true
    End
  End

  Context
    disable_backup() {
      taikun project backup disable $pid -q
    }
    Before 'disable_backup'

    Example 'disable backup'
      When call get_backup_status $pid
      The status should equal 0
      The output should equal false
    End
  End
End
