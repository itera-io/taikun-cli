Context 'billing/rule/label/remove'

  setup() {
    name="$(_rnd_name)"
    pass="$PROMETHEUS_PASSWORD"
    url="$PROMETHEUS_URL"
    user="$PROMETHEUS_USERNAME"

    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    cid=$(taikun billing credential add "$name" -p "$pass" -u "$url" -l "$user" -o "$oid" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun billing credential delete "$cid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  add_rule() {
    id=$(taikun billing rule add "$name" -b "$cid" -l edit=vim,lang=rust -m abc --price 1 --price-rate 1 --type count -I)
    edit_lid=$(taikun billing rule label list "$id" --no-decorate -C id,label | grep edit | cut -d ' ' -f 1)
    lang_lid=$(taikun billing rule label list "$id" --no-decorate -C id,label | grep lang | cut -d ' ' -f 1)
  }
  BeforeEach 'add_rule'

  rm_rule() {
    taikun billing rule delete "$id" -q 2> /dev/null || true
  }
  AfterEach 'rm_rule'

  Context
    delete_edit_label() {
      taikun billing rule label delete "$edit_lid" --billing-rule-id "$id" -q
    }
    Before 'delete_edit_label'

    Example 'delete one of two existing labels'
      When call taikun billing rule label list "$id" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'lang'
      The output should include 'rust'
      The output should include "$lang_lid"
    End
  End

  Context
    delete_lang_label() {
      taikun billing rule label delete "$lang_lid" --billing-rule-id "$id" -q
    }
    Before 'delete_lang_label'

    Example 'delete other of two existing labels'
      When call taikun billing rule label list "$id" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'edit'
      The output should include 'vim'
      The output should include "$edit_lid"
    End
  End

  Example 'delete two existing labels twice'
    When call taikun billing rule label delete "$edit_lid" "$lang_lid" "$edit_lid" "$lang_lid" --billing-rule-id "$id"
    The status should equal 1
    The output should include 'was deleted successfully'
    The output should include "edit_lid"
    The output should include "lang_lid"
    The stderr should include 'Error: Failed to delete one or more resources'
  End
End
