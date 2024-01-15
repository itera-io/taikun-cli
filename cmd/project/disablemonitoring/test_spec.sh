Context 'project/disablemonitoring'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    cc=$(taikun cloud-credential aws add "$(_rnd_name)" -o "$oid" --access-key-id "$AWS_ACCESS_KEY_ID" --secret-access-key "$AWS_SECRET_ACCESS_KEY" --region "$AWS_DEFAULT_REGION" --az-count "$AWS_AZ_COUNT" -I | xargs)
    id=$(taikun project add "$(_rnd_name)" -o "$oid" --cloud-credential-id "$cc" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project delete "$id" -q
    taikun cloud-credential delete "$cc" -q
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

disable_monitoring() {
      taikun project disable-monitoring "$id" -q 2>/dev/null || true
  }

  BeforeEach 'disable_monitoring'

  get_monitoring_status() {
    taikun project info "$1" --no-decorate | grep -i monitoring
  }

  Context
    enablemonitoring() {
      taikun project enable-monitoring "$id" -q
    }
    Before 'enablemonitoring'

    Example 'enable monitoring'
      When call taikun project info "$id" --columns monitoring
      The status should equal 0
      The output should include 'Yes'
    End
  End

  Context
    enable_and_disable_monitoring() {
      taikun project enable-monitoring "$id" -q
      taikun project disable-monitoring "$id" -q
    }
    Before 'enable_and_disable_monitoring'

    Example 'disable monitoring'
      When call get_monitoring_status $id
      The status should equal 0
      The output should include 'No'
    End
  End

  Example 'disable monitoring for project with monitoring already disabled'
    When call taikun project disable-monitoring $id
    The status should equal 1
    The stderr should include 'monitoring already disabled'
  End

End
