Context 'showback/rule/list'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs )
    ids=""
    for i in {1..5}; do
      id=$(taikun showback rule add "$(_rnd_name)" -t count -k general -m idk --global-alert-limit "$i" --price "$i" -o "$oid" -I | xargs)
      ids="$ids $id"
    done
  }
  BeforeAll 'setup'

  cleanup() {
    for id in $ids; do
      taikun showback rule delete "$id" -q
    done
    taikun organization delete "$oid" -q
  }
  AfterAll 'cleanup'

  Example 'negative limit causes error'
    When call taikun showback rule list --limit -42 -o "$oid"
    The status should equal 1
    The stderr should equal 'Error: the --limit flag must be positive'
  End

  Example 'list two showback rules'
    When call taikun showback rule list --limit 2 -o "$oid" --no-decorate
    The status should equal 0
    The lines of output should equal 2
  End

  Example 'list five showback rules'
    When call taikun showback rule list --limit 5 -o "$oid" --no-decorate
    The status should equal 0
    The lines of output should equal 5
  End

  Example 'list cheapest showback rule'
    When call taikun showback rule list --limit 1 -o "$oid" --sort-by price --no-decorate
    The status should equal 0
    The lines of output should equal 1
    The word 8 of output should equal 1
  End

End
