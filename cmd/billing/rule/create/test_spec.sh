Context 'billing/rule/create'

  setup() {
    name=$(_rnd_name)
    cname=$(_rnd_name)
    cid=$(taikun billing credential create -p $PROMETHEUS_PASSWORD -u $PROMETHEUS_URL -l $PROMETHEUS_USERNAME $cname -I)
    flags="-b $cid -l foo=bar -m foo --price 1 --price-rate 5 -t count"
    id=$(taikun billing rule create $name $flags -I)
  }

  cleanup() {
    taikun billing rule delete $id -q 2>/dev/null || true
    taikun billing credential delete $cid -q 2>/dev/null || true
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
    The stderr should include 'Duplicate rule occured'
  End

End
