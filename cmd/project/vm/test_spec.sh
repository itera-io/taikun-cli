Context 'project/vm'
  setup() {
    cc=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $cc --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    id=$(taikun project add $(_rnd_name) --cloud-credential-id $cc --flavors $flavor -I)
    img=$(taikun cc images $cc --columns id --limit 1 --no-decorate)
    profile=$(taikun standalone-profile add $(_rnd_name) --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I)
    taikun project image bind $id --image-ids $img -q
    vm=$(taikun project vm add $id --name $(_rnd_name) --flavor $flavor --image-id $img --volume-size 5 --profile $profile --tags foo=bar -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project vm delete $id --vm-ids $vm -q
    if ! taikun project delete $id -q 2>/dev/null; then
      taikun project delete --force $id -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $cc -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'list vms'
    When call taikun project vm list $id --no-decorate
    The status should equal 0
    The lines of output should equal 1
  End

  Example 'list tags of vm'
    When call taikun project vm tags $vm --no-decorate
    The status should equal 0
    The lines of output should equal 1
    The output should include 'foo'
    The output should include 'bar'
  End
End
