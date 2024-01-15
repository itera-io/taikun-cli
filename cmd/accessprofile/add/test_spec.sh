Context 'accessprofile/add'
  setup() {
    name="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
  }

  cleanup() {
    taikun access-profile delete "$id" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'basic access profile'
    run() {
      taikun access-profile add "$name" -o "$oid" -q
      taikun access-profile list
    }

    When call run
    The output should include "$name"
    The status should equal 0
  End

  Context
    add_access_profile() {
      taikun access-profile add "$name" -o "$oid" -q
    }
    Before 'add_access_profile'

    Example 'duplicate names'
      When call taikun access-profile add "$name" -o "$oid"
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
