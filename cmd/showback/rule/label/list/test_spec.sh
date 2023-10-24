Context 'showback/rule/label/list'

  setup() {
    name="$(_rnd_name)"
    id=$(taikun showback rule add "$name" -t sum -k general -m wat --global-alert-limit 4 --price 2 -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun showback rule delete "$id" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Example 'empty list'
    When call taikun showback rule label list "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context
    add_labels() {
      taikun showback rule label add "$id" --label foo0 --value bar0 --quiet
      taikun showback rule label add "$id" --label foo1 --value bar1 --quiet
      taikun showback rule label add "$id" --label foo2 --value bar2 --quiet
    }
    BeforeAll 'add_labels'

    Example 'list only one label'
      When call taikun showback rule label list "$id" --no-decorate --limit 1
      The status should equal 0
      The lines of output should equal 1
    End

    Example 'list all label'
      When call taikun showback rule label list "$id" --no-decorate
      The status should equal 0
      The lines of output should equal 3
    End
  End
End
