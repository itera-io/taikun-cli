Context 'showback/rule/label/list'

  setup() {
    name=$(_rnd_name)
    id=$(taikun showback rule create $name -t sum -k general -m wat --global-alert-limit 4 --price 2 -I)
  }

  cleanup() {
    taikun showback rule delete $id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'empty list'
    When call taikun showback rule label list $id --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context

    add_labels() {
      : # TODO add 3 labels
    }

    BeforeEach 'add_labels'


    Example 'list only one label'
      When call taikun showback rule label list $id --no-decorate --limit 1
      The status should equal 0
      The lines of output should equal 1
    End

    Example 'list all label'
      When call taikun showback rule label list $id --no-decorate
      The status should equal 0
      The lines of output should equal 3
    End

  End

End
