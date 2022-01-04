Context 'showback/rule/label/add'

  setup() {
    name=$(_rnd_name)
    id=$(taikun showback rule create $name -t sum -k general -m wat --global-alert-limit 4 --price 2 -I)
  }

  cleanup() {
    taikun showback rule delete $id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'add label to showback rule'

    When call taikun showback rule label add $id --label foo --value bar --no-decorate
    The status should equal 0
    The output should include foo
    The output should include bar
    The lines of output should equal 1

  End

End
