Context 'project/k8s/remove'

  setup() {
    ccid=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $ccid --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid --flavors $flavor -I)
  }

  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete $pid -q 2>/dev/null; then
      taikun project delete --force $pid -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $ccid -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Context

    add_and_remove_master() {
      msid=$(taikun project k8s add $pid -n master -r kubemaster -f $flavor -I)
      taikun project k8s delete $pid --server-ids $msid -q
    }

    BeforeEach 'add_and_remove_master'

    Example 'delete one server'
      When call taikun project k8s list $pid --no-decorate
      The status should equal 0
      The lines of output should equal 0
    End

    Example 'delete the same server twice'
      When call taikun project k8s delete $pid --server-ids $msid
      The status should equal 1
      The stderr should include '404'
    End
  End

  Context

    add_servers_then_remove_all() {
      msid=$(taikun project k8s add $pid -n m -r kubemaster -f $flavor -I)
      wsid=$(taikun project k8s add $pid -n w -r kubeworker -f $flavor -I)
      bsid=$(taikun project k8s add $pid -n b -r bastion -f $flavor -I)
      taikun project k8s delete $pid --all-servers -q
    }

    Before 'add_servers_then_remove_all'

    Example 'delete all servers'
      When call taikun project k8s list $pid --no-decorate
      The status should equal 0
      The lines of output should equal 0
    End
  End
End
