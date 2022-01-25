Context 'project/vm'
  setup() {
    cc=$(taikun cloud-credential openstack add $(_rnd_name) -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    flavor=$(taikun cloud-credential flavors $cc --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1)
    id=$(taikun project add $(_rnd_name) --cloud-credential-id $cc --flavors $flavor -I)
    img=$(taikun cc images $cc --columns id --limit 1 --no-decorate)
    taikun project image bind $id --image-ids $img -q
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete $id -q 2>/dev/null; then
      taikun project delete --force $id -q 2>/dev/null || true
    fi
    taikun cloud-credential delete $cc -q 2>/dev/null || true
  }
  AfterAll 'cleanup'
End
