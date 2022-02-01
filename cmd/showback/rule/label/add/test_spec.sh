Context 'showback/rule/label/add'

  setup() {
    name=$(_rnd_name)
    id=$(taikun showback rule add $name -t sum -k general -m wat --global-alert-limit 4 --price 2 -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun showback rule delete $id -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'add label to showback rule'
    When call taikun showback rule label add $id --label foo --value bar --no-decorate
    The status should equal 0
    The output should equal 'Operation was successful.'
  End
End
