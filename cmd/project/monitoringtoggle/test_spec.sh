Context 'project/monitoringtoggle'
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

  Context
    toggle() {
      taikun project monitoring-toggle $id -q
    }
    BeforeEach 'toggle'
    AfterEach 'toggle'

    Example 'toggle on and off'
      When call taikun project info $id --columns monitoring
      The status should equal 0
      The output should include 'Yes'
    End
  End
End
