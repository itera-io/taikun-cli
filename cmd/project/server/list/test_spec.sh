Context 'project/server/list'

  setup() {
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
  }

  AfterEach 'cleanup'

  Example 'empty project'
    When call taikun project server list $pid --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context
    add_servers() {
      bsid=$(taikun project server add -p $pid bastion -r bastion -f $flavor -I)
      msid=$(taikun project server add -p $pid master -r kubemaster -f $flavor -I)
      wsid=$(taikun project server add -p $pid worker -r kubeworker -f $flavor -I)
    }

    Before 'add_servers'

    remove_servers() {
      taikun project server delete -p $pid $bsid $msid $wsid -q 2>/dev/null || true
    }

    After 'remove_servers'

    Example 'project with 3 servers'
      When call taikun project server list $pid --no-decorate
      The status should equal 0
      The lines of output should equal 3
      The output should include 'bastion'
      The output should include 'master'
      The output should include 'worker'
    End
  End
End
