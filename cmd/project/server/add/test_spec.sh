Context 'project/server/add'

  setup() {
    ccid=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $ccid --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid --flavors $flavor -I)
  }

  BeforeEach 'setup'

  cleanup() {
    if ! taikun project delete $pid -q 2>/dev/null; then
      taikun project delete --force $pid -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $ccid -q 2>/dev/null || true
  }

  AfterEach 'cleanup'

  Context
    add_master() {
      msid=$(taikun project server add $pid --name master -r kubemaster -k foo=bar,bar=foo -f $flavor -I)
    }

    BeforeEach 'add_master'

    remove_master() {
      taikun project server delete $pid --server-ids $msid -q 2>/dev/null || true
    }

    AfterEach 'remove_master'

    Example 'add one server'
      When call taikun project server list $pid --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'master'
    End

    Example 'add two servers with the same name'
      When call taikun project server add $pid --name master -r kubemaster -f $flavor
      The status should equal 1
      The stderr should include 'Duplicate name occured'
    End
  End

End
