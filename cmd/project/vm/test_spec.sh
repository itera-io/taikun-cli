Context 'project/vm'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    cc=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_APPLICATION_CREDENTIAL_ID" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
    sleep 0.1
    flavor=$(taikun cloud-credential flavors "$cc" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    id=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$cc" --flavors "$flavor" -I | xargs)
    img=$(taikun cc images "$cc" --no-decorate | grep -E -i '(ubuntu)|(focal)' | head -1 | cut -d ' ' -f 1 | xargs)
    image=$(taikun cc images "$cc" --no-decorate | grep -E -i 'Debian' | head -1 | cut -d ' ' -f 1 | xargs)
    image_name=$(taikun cc images "$cc" -C name --no-decorate | grep -E -i 'Debian' | head -1 | xargs)
    profile=$(taikun standalone-profile add "$(_rnd_name)" -o "$oid" --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I | xargs)
    taikun project image bind "$id" --image-ids "$img" -q
    vm_onetag=$(taikun project vm add "$id" --name onetag --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" --tags foo=bar -I | xargs)
    vm_notags=$(taikun project vm add "$id" --name notags --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" -I | xargs)
    vm_manytags=$(taikun project vm add "$id" --name manytags --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" --tags foo=bar,editor=vim -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project vm delete "$id" --vm-ids $vm_onetag,$vm_notags,$vm_manytags -q
    taikun project vm delete "$id" --all-project -q
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

  Context 'add two VMs with count'
      Example 'list image empty'
          When call taikun project image list "$id" --no-decorate
          The status should equal 0
          The lines of output should equal 1
      End
      Example 'bind image'
          When call taikun project image bind "$id" -i "$image"
          The status should equal 0
          The lines of output should equal 1
          The stdout should include 'successful'
      End
      Example 'list image exists'
          When call taikun project image list "$id" --no-decorate
          The status should equal 0
          The lines of output should equal 2
          The output should include "$image"
          The output should include "$image_name"
      End
      Example 'create 3 vms with image'
          When call taikun project vm add "$id" --count 3 --name countvms --flavor "$flavor" --image-id "$image" --volume-size 25 --standalone-profile-id "$profile"
          The status should equal 0
          The output should include "countvms"
      End
      Example 'list 3 vms with image'
          When call taikun project vm list "$id" --no-decorate
          The status should equal 0
          The output should include "countvms-1"
          The output should include "countvms-2"
          The output should include "countvms-3"
          The output should include "onetag"
          The output should include "notags"
          The output should include "manytags"
          The output should include "$image_name"
          The output should include "$flavor"
      End
  End

End
