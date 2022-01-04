Context 'showback/rule/list'

  # TODO add --no-decorate flag where needed

  setup() {
    ids=""
    for i in {1..5}; do
      name=$(_rnd_name)
      id=$(taikun showback rule create $name -t count -k general -m idk --global-alert-limit $i --price $i -I)
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

  Example 'list one showback rule'
    When call taikun showback rule list --limit 1
    The status should equal 0
    The lines of output should equal 1
  End

  Example 'list five showback rules'
    When call taikun showback rule list --limit 5
    The status should equal 0
    The lines of output should equal 5
  End

  Example 'list cheapest showback rule'
    When call taikun showback rule list --limit 1 --sort-by price
    The status should equal 0
    The lines of output should equal 1
    The word 9 of output should equal 1
  End

End
