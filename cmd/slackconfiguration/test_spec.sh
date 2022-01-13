Context 'slackconfiguration'

  setup() {
    oid=$(taikun organization create $(_rnd_name) --full-name $(_rnd_name) -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete $oid -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No slack configurations'
    When call taikun slack-configuration list -o $oid --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

End
