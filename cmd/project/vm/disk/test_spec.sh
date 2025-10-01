Context 'project/vm/disk'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    cc=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_APPLICATION_CREDENTIAL_SECRET" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_APPLICATION_CREDENTIAL_ID" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
    sleep 0.1
    flavor=$(taikun cloud-credential flavors "$cc" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    id=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$cc" --flavors "$flavor" -I | xargs)
    img=$(taikun cc images "$cc" --no-decorate | grep -E -i '(ubuntu)|(focal)' | head -1 | cut -d ' ' -f 1 | xargs)
    profile=$(taikun standalone-profile add "$(_rnd_name)" -o "$oid" --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I | xargs)
    taikun project image bind "$id" --image-ids "$img" -q
    vm=$(taikun project vm add "$id" --name foo --flavor "$flavor" --image-id "$img" --volume-size 8 --standalone-profile-id "$profile" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project vm delete "$id" --vm-ids "$vm" -q
    taikun standalone-profile delete "$profile" -q 2>/dev/null || true
    if ! taikun project delete "$id" -q 2>/dev/null; then
      taikun project delete --force "$id" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$cc" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'no disks initially'
    When call taikun project vm disk list "$vm" --project-id "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context
    add() {
      disk=$(taikun project vm disk add "$vm" --name 'ext' --size 8 --openstack-volume-type "__DEFAULT__" -I | xargs)
    }
    BeforeAll 'add'

    remove() {
      taikun project vm disk delete "$vm" --disk-ids "$disk" -q
    }
    AfterAll 'remove'

    Example 'add and then remove disk'
      When call taikun project vm disk list "$vm" --project-id "$id" --columns name,size,volume-type --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'ext'
      The output should include '8'
      The output should include '__DEFAULT__'
    End

    Context
      resize() {
        taikun project vm disk resize "$disk" --size 10 -q
      }
      BeforeEach 'resize'

      Example 'resize disk'
        When call taikun project vm disk list "$vm" --project-id "$id" --columns target-size --no-decorate
        The status should equal 0
        The lines of output should equal 1
        The output should include '10'
      End
    End
  End
End
