Context 'showback/rule/label/clear'

  setup() {
    name=$(_rnd_name)
    id=$(taikun showback rule add "$name" -t count -k general -m any --global-alert-limit 3 --price 3 -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun showback rule delete "$id" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  add_and_clear_labels() {
    taikun showback rule label add "$id" --label foo0 --value bar0 --quiet
    taikun showback rule label add "$id" --label foo1 --value bar1 --quiet
    taikun showback rule label add "$id" --label foo2 --value bar2 --quiet
    taikun showback rule label clear "$id" --quiet
  }

  BeforeEach 'add_and_clear_labels'

  Example 'no labels after clear'
    When call taikun showback rule label list "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

End
