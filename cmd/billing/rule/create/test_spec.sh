Context 'billing/rule/create'

  flags="" # FIXME

  setup() {
    name=$(_rnd_name)
    id=$(taikun billing rule create $name $flags -I)
  }

  cleanup() {
    taikun billing rule delete $id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'create single billing rule'
    When call taikun billing rule list --no-decorate
    The status should equal 0
    The output should include $name
  End

  Example 'duplicate name causes error'
    When call taikun billing rule create $name $flags
    The status should equal 1
    The stderr should include '400'
    The stderr should include 'already exists'
  End

End
