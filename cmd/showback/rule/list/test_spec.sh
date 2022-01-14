Context 'showback/rule/list'

  setup() {
    ids=""
    for i in {1..5}; do
      name=$(_rnd_name)
      id=$(taikun showback rule add $name -t count -k general -m idk --global-alert-limit $i --price $i -I)
      ids="$ids $id"
    done
  }

  cleanup() {
    for id in $ids; do
      taikun showback rule delete $id -q 2>/dev/null || true
    done
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'negative limit causes error'
    When call taikun showback rule list --limit -42
    The status should equal 1
    The stderr should equal 'Error: The --limit flag must be positive.'
  End

  Example 'list two showback rules'
    When call taikun showback rule list --limit 2 --no-decorate
    The status should equal 0
    The lines of output should equal 2
  End

  Example 'list five showback rules'
    When call taikun showback rule list --limit 5 --no-decorate
    The status should equal 0
    The lines of output should equal 5
  End

  Example 'list cheapest showback rule'
    When call taikun showback rule list --limit 1 --sort-by price --no-decorate
    The status should equal 0
    The lines of output should equal 1
    The word 8 of output should equal 1
  End

End
