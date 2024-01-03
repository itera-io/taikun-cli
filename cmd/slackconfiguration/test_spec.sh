Context 'slackconfiguration'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No slack configurations'
    When call taikun slack-configuration list -o "$oid" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context 'add bad slack url'

    Example 'add slack config which does not pass verifycation'
      When call taikun slack-configuration add "$(_rnd_name)" -c foo -u 'http://foo.bar.test' -t alert -o "$oid" -I
      The status should equal 1
      The stderr should include 'Wrong channel or slack webhook url detected'
    End
  End

  Context 'add/delete, channel name with dash'
    add_config() {
      name="$(_rnd_name)"
      scid=$(taikun slack-configuration add "$name" -c cli-test -u "$SLACK_WEBHOOK" -t alert -o "$oid" -I)
    }

    Before 'add_config'

    delete_config() {
      taikun slack-configuration delete "$scid" -q
    }

    After 'delete_config'

    Example 'add then delete slack config'
      When call taikun slack-configuration list -o "$oid" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include "$name"
    End
  End

  Context 'add/delete, channel name with number'
    add_config() {
      name="$(_rnd_name)"
      scid=$(taikun slack-configuration add "$name" -c 4aaa -u "$SLACK_WEBHOOK" -t alert -o "$oid" -I)
    }
    Before 'add_config'

    delete_config() {
      taikun slack-configuration delete "$scid" -q
    }
    After 'delete_config'

    Example 'add then delete slack config'
      When call taikun slack-configuration list -o "$oid" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include "$name"
    End
  End

End
