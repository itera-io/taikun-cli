Context 'slackconfiguration'

  setup() {
    oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name) -I)
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

  Context 'add/delete'
    add_config() {
      scid=$(taikun slack-configuration add $(_rnd_name) -c foo -u 'http://foo.bar.test' -t alert -o $oid -I)
    }

    Before 'add_config'

    delete_config() {
      taikun slack-configuration delete $scid -q
    }

    After 'delete_config'

    Example 'add then delete slack config'
      When call taikun slack-configuration list -o $oid --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'http://foo.bar.test'
    End
  End
End
