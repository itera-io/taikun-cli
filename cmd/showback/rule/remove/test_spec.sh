Context 'showback/rule/remove'

  Example 'delete nonexistent showback rule'

    When call taikun showback rule delete 0
    The status should equal 1
    The stderr should include 404
    The stderr should include '(0) was not found'

  End

  Context

    setup() {
      name=$(_rnd_name)
      id=$(taikun showback rule add "$name" -t sum -k general -m node --global-alert-limit 10 --price 1 -I)
    }
    BeforeEach 'setup'

    cleanup() {
      taikun showback rule delete "$id" -q 2>/dev/null || true
    }
    AfterEach 'cleanup'

    Example 'delete existing showback rule'
      When call taikun showback rule delete "$id"
      The status should equal 0
      The output should include 'was deleted successfully'
      The output should include "$id"
    End

    Example 'delete existing and nonexistent showback rules'
      When call taikun showback rule delete "$id" 0
      The status should equal 1
      The output should include 'was deleted successfully'
      The output should include "$id"
      The stderr should include 404
      The stderr should include '(0) was not found'
    End

  End

End
