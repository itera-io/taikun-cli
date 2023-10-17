Context 'project/k8s/list'

  setup() {
    oid=$(taikun organization add $(_rnd_name) -f $(_rnd_name) -I)
    ccid=$(taikun cloud-credential openstack add $(_rnd_name) -o $oid -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $ccid --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    pid=$(taikun project add $(_rnd_name) -o $oid --cloud-credential-id $ccid --flavors $flavor -I)
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete $pid -q 2>/dev/null; then
      taikun project delete --force $pid -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $ccid -q 2>/dev/null || true
    taikun organization delete $oid
  }
  AfterAll 'cleanup'

  Example 'empty project'
    When call taikun project k8s list $pid --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context
    add_servers() {
      bsid=$(taikun project k8s add $pid -n bastion -r bastion -f $flavor -I)
      msid=$(taikun project k8s add $pid -n master -r kubemaster -f $flavor -I)
      wsid=$(taikun project k8s add $pid -n worker -r kubeworker -f $flavor -I)
    }
    Before 'add_servers'

    remove_servers() {
      taikun project k8s delete $pid --all-servers -q
    }
    After 'remove_servers'

    Example 'project with 3 servers'
      When call taikun project k8s list $pid --no-decorate
      The status should equal 0
      The lines of output should equal 3
      The output should include 'bastion'
      The output should include 'master'
      The output should include 'worker'
    End
  End
End

#Context 'project/k8s/list'
#
#  setup() {
#    oid=$(taikun organization add $(_rnd_name) -f $(_rnd_name) -I)
#    ccid=$(taikun cloud-credential aws add $(_rnd_name) -a $AWS_ACCESS_KEY_ID -s $AWS_SECRET_ACCESS_KEY -r $AWS_DEFAULT_REGION -z 3 -o $oid -I)
#    flavor=$(taikun cloud-credential flavors $ccid --no-decorate --limit 1 -C name)
#    pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid -o $oid --flavors $flavor -I)
#  }
#  BeforeAll 'setup'
#
#  cleanup() {
#    if ! taikun project delete $pid -q 2>/dev/null; then
#      taikun project delete --force $pid -q 2>/dev/null || true
#    fi
#    taikun cloud-credential delete $ccid -q 2>/dev/null || true
#    taikun organization delete $oid
#  }
#  AfterAll 'cleanup'
#
#  Context
#    add_servers() {
#      bsid=$(taikun project k8s add $pid -n bastion -r bastion -f $flavor -a a -I)
#      msid=$(taikun project k8s add $pid -n master -r kubemaster -f $flavor -a b -I)
#      wsid=$(taikun project k8s add $pid -n worker -r kubeworker -f $flavor -a c -I)
#    }
#    Before 'add_servers'
#
#    remove_servers() {
#      taikun project k8s delete $pid -a -q 2>/dev/null || true
#    }
#    AfterAll 'remove_servers'
#
#    Example 'project with 3 servers with availability zone'
#      When call taikun project k8s list $pid --no-decorate
#      The status should equal 0
#      The lines of output should equal 3
#      The output should include $AWS_DEFAULT_REGION'a'
#      The output should include $AWS_DEFAULT_REGION'b'
#      The output should include $AWS_DEFAULT_REGION'c'
#    End
#  End
#End
