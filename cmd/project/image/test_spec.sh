Context 'project/image'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
    cc=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    sleep 0.1
    id=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$cc" -o "$oid" -I)
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete "$id" -q 2>/dev/null; then
      taikun project delete --force "$id" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$cc" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'no image bounds initially'
    When call taikun project image list "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context
    bind() {
      img=$(taikun cloud-credential images "$cc" --limit 1 --columns id --no-decorate | xargs)
      taikun project image bind "$id" --image-ids "$img" -q
    }
    BeforeEach 'bind'

    unbind() {
      binding=$(taikun project image list "$id" --no-decorate --columns binding-id)
      taikun project image unbind $binding -q
    }
    AfterEach 'unbind'

    Example 'bind then unbind image'
      When call taikun project image list "$id" --no-decorate --columns image-id
      The status should equal 0
      The lines of output should equal 1
      The output should include "$img"
    End
  End
End
