Context 'billing/rule/organization/list'

  setup() {
    name=$(_rnd_name)
    pass=$PROMETHEUS_PASSWORD
    url=$PROMETHEUS_URL
    user=$PROMETHEUS_USERNAME

    cid=$(taikun billing cred create $name -p $pass -u $url -l $user -I)
    id=$(taikun billing rule create $name -b $cid -l foo=foo -m abc --price 1 --price-rate 1 --type count -I)
  }

  BeforeEach 'setup'

  cleanup() {
    taikun billing rule delete $id -q 2> /dev/null || true
    taikun billing credential delete $cid -q 2>/dev/null || true
  }

  AfterEach 'cleanup'

  Example 'no bindings'
    When call taikun billing rule org list $id --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Todo 'multiple bindings'

End
