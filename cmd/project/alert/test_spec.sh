Context 'project/alert'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    apid=$(taikun alerting-profile add "$(_rnd_name)" -o "$oid" -I)
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    pid=$(taikun project add "$(_rnd_name)" -o "$oid" --cloud-credential-id "$ccid" -I)
    taikun project alert detach "$pid" -q
  }

  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete "$pid" -q 2>/dev/null; then
      taikun project delete --force "$pid" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun alerting-profile delete $apid -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Context
    attach_then_detach() {
      taikun project alert attach "$pid" --alerting-profile-id "$apid" -q
      taikun project alert detach "$pid" -q
    }

    Before 'attach_then_detach'

    project_has_alerting_profile() {
      taikun project info "$pid" | grep -i alerting
    }

    Example 'attach then detach'
      When call project_has_alerting_profile
      The status should equal 0
      The output should include 'ALERTING-PROFILE'
      The output should include 'N/A'
    End
  End

  Context
    attach() {
      taikun project alert attach "$pid" --alerting-profile-id "$apid" -q
    }

    Before 'attach'

    detach() {
      taikun project alert detach "$pid" -q
    }

    After 'detach'

    Example 'attach twice causes error'
      When call taikun project alert attach "$pid" --alerting-profile-id "$apid"
      The status should equal 1
      The stderr should include 'This alerting profile already assigned to this project'
    End
  End


End
