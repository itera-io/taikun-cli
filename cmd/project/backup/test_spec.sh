Context 'project/backup'

  setup() {
    bid=$(taikun backup-credential add $(_rnd_name) -a $AWS_ACCESS_KEY_ID -e $S3_ENDPOINT -r $S3_REGION -s $AWS_SECRET_ACCESS_KEY -I)
    ccid=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid -I)
  }

  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete $pid -q 2>/dev/null; then
      taikun project delete --force $pid -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $ccid -q 2>/dev/null || true
    taikun backup-credential delete $bid -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  disable_backup() {
      taikun project backup disable $pid -q 2>/dev/null || true
  }

  BeforeEach 'disable_backup'

  get_backup_status() {
    taikun project info $1 --no-decorate | grep -i backup
  }

  Context
    enable_backup() {
      $(taikun project backup enable $pid -b $bid -q)
    }
    Before 'enable_backup'

    Example 'enable backup'
      When call get_backup_status $pid
      The status should equal 0
      The output should include 'Yes'
    End
  End

  Context
    enable_and_disable_backup() {
      $(taikun project backup enable $pid -b $bid -q)
      taikun project backup disable $pid -q
    }
    Before 'enable_and_disable_backup'

    Example 'disable backup'
      When call get_backup_status $pid
      The status should equal 0
      The output should include 'No'
    End
  End

  Example 'disable backup for project with backup already disabled'
    When call taikun project backup disable $pid
    The status should equal 1
    The stderr should include 'Backup already disabled for this project'
  End
End
