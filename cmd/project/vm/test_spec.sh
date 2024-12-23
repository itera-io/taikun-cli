Context 'project/vm'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    cc=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I | xargs)
    sleep 0.1
    flavor=$(taikun cloud-credential flavors "$cc" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    id=$(taikun project add "$(_rnd_name)" -o "$oid" --cloud-credential-id "$cc" --flavors "$flavor" -I | xargs)
    img=$(taikun cc images "$cc" --no-decorate | grep -E -i '(ubuntu)|(focal)' | head -1 | cut -d ' ' -f 1 | xargs)
    profile=$(taikun standalone-profile add "$(_rnd_name)" -o "$oid" --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I | xargs)
    taikun project image bind "$id" --image-ids "$img" -q
    vm_onetag=$(taikun project vm add "$id" --name onetag --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" --tags foo=bar -I | xargs)
    vm_notags=$(taikun project vm add "$id" --name notags --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" -I | xargs)
    vm_manytags=$(taikun project vm add "$id" --name manytags --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" --tags foo=bar,editor=vim -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project vm delete "$id" --vm-ids $vm_onetag,$vm_notags,$vm_manytags -q
    taikun standalone-profile delete "$profile" -q 2>/dev/null || true
    if ! taikun project delete "$id" -q 2>/dev/null; then
      taikun project delete --force "$id" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$cc" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'add and then remove 3 project'
    When call taikun project vm list "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 3
  End

  notags() {
    taikun project vm list "$id" --columns name,tags --no-decorate | grep notags | tr -s ' ' | cut -d ' ' -f 2
  }

  Example 'add vm without tags'
    When call notags
    The status should equal 0
    The output should equal 'N/A'
  End

  onetag() {
    taikun project vm list "$id" --columns name,tags --no-decorate | grep onetag | tr -s ' ' | cut -d ' ' -f 2
  }

  Example 'add vm with one tag'
    When call onetag
    The status should equal 0
    The output should equal 'foo=bar'
  End

  manytags() {
    taikun project vm list "$id" --columns name,tags --no-decorate --show-large-values | grep manytags | tr -s ' ' | cut -d ' ' -f 2-
  }

  Example 'add vm with multiple tags'
    When call manytags
    The status should equal 0
    The output should include 'foo=bar,editor=vim'
  End
End
