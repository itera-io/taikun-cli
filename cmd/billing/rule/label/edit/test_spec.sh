Context 'billing/rule/label/edit'

  setup() {
    name="$(_rnd_name)"
    pass="$PROMETHEUS_PASSWORD"
    url="$PROMETHEUS_URL"
    user="$PROMETHEUS_USERNAME"

    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun billing credential add "$name" -p "$pass" -u "$url" -l "$user" -o "$oid" -I | xargs)
    id=$(taikun billing rule add "$name" -b "$cid" -l ed=vim -m abc --price 1 --price-rate 1 --type count -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing rule delete "$id" -q 2> /dev/null || true
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  edit_label() {
    taikun billing rule label edit "$id" -l lang=rust -q
  }
  Before 'edit_label'

  Example 'edit a label'
    When call taikun billing rule label list "$id" --no-decorate
    The status should equal 0
    The lines of output should equal 1
    The output should not include vim
    The output should include rust
    The output should not include ed
    The output should include lang
  End

End
