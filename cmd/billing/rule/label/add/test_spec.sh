Context 'billing/rule/label/add'

  setup() {
    name=$(_rnd_name)
    pass=$PROMETHEUS_PASSWORD
    url=$PROMETHEUS_URL
    user=$PROMETHEUS_USERNAME

    cid=$(taikun billing cred create $name -p $pass -u $url -l $user -I)
    id=$(taikun billing rule create $name -b $cid -l ed=vim -m abc --price 1 --price-rate 1 --type count -I)
  }

  BeforeEach 'setup'

  cleanup() {
    taikun billing rule delete $id -q 2> /dev/null || true
    taikun billing credential delete $cid -q 2>/dev/null || true
  }

  AfterEach 'cleanup'

  add_label() {
    taikun billing rule label add $id -l lang -v rust -q
  }

  Before 'add_label'

  Example 'Add a label'
    When call taikun billing rule label list $id --no-decorate
    The status should equal 0
    The lines of output should equal 2
    The output should include vim
    The output should include rust
    The output should include ed
    The output should include lang
  End

End
