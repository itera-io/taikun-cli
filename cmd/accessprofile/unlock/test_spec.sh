Context 'accessprofile/unlock'
  cleanup() {
    taikun access-profile delete "$id" -q
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Context
    setup() {
      oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
      name=$(_rnd_name)
      id=$(taikun access-profile add "$name" -o "$oid" -I | xargs)
    }

    Example 'unlocked access profile'
      When call taikun access-profile unlock "$id"
      The stderr should include '400'
      The status should equal 1
    End
  End

  Context
    setup() {
      oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
      name=$(_rnd_name)
      id=$(taikun access-profile add "$name" -o "$oid" -I | xargs)
      taikun access-profile lock "$id" -q
    }

    Example 'locked access profile'
      When call taikun access-profile unlock "$id"
      The output should equal 'Operation was successful.'
      The status should equal 0
    End
  End
End
