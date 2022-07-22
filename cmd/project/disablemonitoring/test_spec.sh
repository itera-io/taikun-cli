Context 'project/disablemonitoring'
  setup() {
    cc=$(taikun cloud-credential aws add $(_rnd_name) --access-key-id $AWS_ACCESS_KEY_ID --secret-access-key $AWS_SECRET_ACCESS_KEY --region $AWS_DEFAULT_REGION --availability-zone $AWS_AVAILABILITY_ZONE -I)
    id=$(taikun project add $(_rnd_name) --cloud-credential-id $cc -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun project delete $id -q
    taikun cloud-credential delete $cc -q
  }
  AfterAll 'cleanup'

disable_monitoring() {
      taikun project disable-monitoring $pid -q 2>/dev/null || true
  }

  BeforeEach 'disable_monitoring'

  get_monitoring_status() {
    taikun project info $1 --no-decorate | grep -i monitoring
  }

  Context
    enablemonitoring() {
      taikun project enable-monitoring $id -q
    }
    Before 'enablemonitoring'

    Example 'enable monitoring'
      When call taikun project info $id --columns monitoring
      The status should equal 0
      The output should include 'Yes'
    End
  End

  Context
    enable_and_disable_monitoring() {
      taikun project enable-monitoring $pid -q
      taikun project disable-monitoring $pid -q
    }
    Before 'enable_and_disable_monitoring'

    Example 'disable monitoring'
      When call get_monitoring_status $pid
      The status should equal 0
      The output should include 'No'
    End
  End

  Example 'disable monitoring for project with monitoring already disabled'
    When call taikun project disable-monitoring $pid
    The status should equal 1
    The stderr should equal 'Error: Project monitoring already disabled'
  End

End
