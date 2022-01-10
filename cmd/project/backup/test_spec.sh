Context 'project/backup'

  setup() {
    bid=$(taikun backup-credential create foo -a $AWS_ACCESS_KEY_ID -e $S3_ENDPOINT -r $S3_REGION -s $AWS_SECRET_ACCESS_KEY -I)
    ccid=$(taikun cloud-credential openstack create $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $ccid --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    pid=$(taikun project create $(_rnd_name) --cloud-credential-id $ccid --flavors $flavor -I)
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

End
