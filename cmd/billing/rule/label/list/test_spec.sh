Context 'billing/rule/label/list'

  setup() {
    name=$(_rnd_name)
    pass=$PROMETHEUS_PASSWORD
    url=$PROMETHEUS_URL
    user=$PROMETHEUS_USERNAME

    cid=$(taikun billing credential add $name -p $pass -u $url -l $user -I)
    id=$(taikun billing rule add $name -b $cid -l ed=vim,lang=rust -m abc --price 1 --price-rate 1 --type count -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing rule delete $id -q
    taikun billing credential delete $cid -q
  }
  AfterAll 'cleanup'

  Example 'list all labels'
    When call taikun billing rule label list $id --no-decorate
    The status should equal 0
    The lines of output should equal 2
    The output should include vim
    The output should include rust
    The output should include ed
    The output should include lang
  End

  Example 'list only one label'
    When call taikun billing rule label list $id --no-decorate --limit 1
    The status should equal 0
    The lines of output should equal 1
  End

End
