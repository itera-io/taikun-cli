Context 'showback/rule/add'
  setup() {
    name=$(_rnd_name)
  }

  cleanup() {
    taikun showback rule delete $id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'missing flags'
    When call taikun showback rule add foo
    The stderr should include "Error: required flag(s)"
    The status should equal 1
  End


  Example 'basic showback rule'
    run() {
      id=$(taikun showback rule add $name -t sum -k general -m node --global-alert-limit 10 --price 1 -I)
      taikun showback rule list | grep $id
    }

    When call run
    The status should equal 0
    The word 1 of output should equal $id
    The output should include "$name"
  End

  Context
    add_showback_rule() {
      id=$(taikun showback rule add $name -t sum -k general -m node --global-alert-limit 10 --price 1 -I)
    }
    Before 'add_showback_rule'

    Example 'duplicate names'
      When call taikun showback rule add $name -t sum -k general -m node --global-alert-limit 10 --price 1
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
