Context 'billing/rule/organization/list'

  setup() {
    name=$(_rnd_name)
    pass=$PROMETHEUS_PASSWORD
    url=$PROMETHEUS_URL
    user=$PROMETHEUS_USERNAME

    org_name1=$(_rnd_name)
    oid1=$(taikun organization create $org_name1 --full-name $org_name1 -I)
    org_name2=$(_rnd_name)
    oid2=$(taikun organization create $org_name2 --full-name $org_name2 -I)

    cid=$(taikun billing credential create $name -p $pass -u $url -l $user -I)
    id=$(taikun billing rule create $name -b $cid -l foo=foo -m abc --price 1 --price-rate 1 --type count -I)
  }

  BeforeEach 'setup'

  cleanup() {
    taikun billing rule delete $id -q 2>/dev/null || true
    taikun billing credential delete $cid -q 2>/dev/null || true
    taikun organization delete $oid1 -q 2>/dev/null || true
    taikun organization delete $oid2 -q 2>/dev/null || true
  }

  AfterEach 'cleanup'

  Example 'no bindings'
    When call taikun billing rule organization list $id --no-decorate
    The status should equal 0
    The lines of output should equal 0
  End

  Context

    bind_org() {
      taikun billing rule organization bind $id -o $oid1 -d 42 -q
      taikun billing rule organization bind $id -o $oid2 -d 42 -q
    }

    BeforeEach 'bind_org'

    Example 'list all bindings'
      When call taikun billing rule organization list $id --no-decorate
      The status should equal 0
      The lines of output should equal 2
      The output should include $org_name1
      The output should include $org_name2
    End

    Example 'list only one binding'
      When call taikun billing rule organization list $id --no-decorate --limit 1
      The status should equal 0
      The lines of output should equal 1
    End

  End

End
