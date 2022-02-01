Context 'showback/rule/add'
  setup() {
    oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name) -I)
    name=$(_rnd_name)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun showback rule delete $id -q 2>/dev/null || true
    taikun organization delete $oid -q
  }
  AfterAll 'cleanup'

  Example 'missing flags'
    When call taikun showback rule add foo
    The stderr should include "Error: required flag(s)"
    The status should equal 1
  End

  Context
    add_rule() {
      id=$(taikun showback rule add $name -t sum -k general -m node --global-alert-limit 10 --price 1 -o $oid -I)
    }
    BeforeAll 'add_rule'

    Example 'basic showback rule'
      list() {
        taikun showback rule list | grep $id
      }

      When call list
      The status should equal 0
      The word 1 of output should equal $id
      The output should include "$name"
    End

    Example 'duplicate names'
      When call taikun showback rule add $name -t sum -k general -m node --global-alert-limit 10 --price 1 -o $oid
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
