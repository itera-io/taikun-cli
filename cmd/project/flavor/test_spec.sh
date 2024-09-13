Context 'project/flavor'
    setup() {
      oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
      ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
      flavors=$(taikun cc flavors "$ccid" --no-decorate -C name | head -1 | xargs)
      pid=$(taikun project add "$(_rnd_name)" -o "$oid" --cloud-credential-id "$ccid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      if ! taikun project delete "$pid" -q 2>/dev/null; then
        taikun project delete --force "$pid" -q 2>/dev/null || true
      fi
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Example 'list empty flavor'
      When call taikun project flavor list "$pid"
      The status should equal 0
      The lines of output should equal 2
      The output should include "RAM"
    End

    Example 'bind flavor'
      When call taikun project flavor bind "$pid" --flavors "$flavors"
      The status should equal 0
      The lines of output should equal 1
      The output should include "Operation was successful."
    End

    Example 'bind already bound flavor'
      When call taikun project flavor bind "$pid" --flavors "$flavors"
      The status should equal 1
      The lines of stderr should equal 1
      The stderr should include "already bounded"
    End

    Example 'list bound flavors'
      When call taikun project flavor list "$pid"
      The status should equal 0
      The lines of output should equal 3
      The output should include "OPENSTACK"
      The output should include "$flavors"
    End

    Example 'unbind flavor'
      When call taikun project flavor unbind $(taikun project flavor list "$pid" --no-decorate | cut -d' ' -f1 | xargs)
      The status should equal 0
      The lines of output should equal 1
      The output should include "Operation was successful."
    End

End