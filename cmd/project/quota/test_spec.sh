Context 'project/quota'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    projectname="$(_rnd_name)"
    pid=$(taikun project add "$projectname" -o "$oid" --cloud-credential-id "$ccid" -I)
    num1="$(_rnd_number)"
    num2="$(_rnd_number)"
    num3="$(_rnd_number)"
    num4="$(_rnd_number)"
    num5="$(_rnd_number)"
    num6="$(_rnd_number)"
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete "$pid" -q 2>/dev/null; then
      taikun project delete --force "$pid" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_project(){
    taikun project quota list -o "$oid" --no-decorate
  }

  Example 'list quota for my project'
    When call list_project
    The status should equal 0
    The lines of output should equal 1
    The output should include "$projectname"
    The output should include "$pid"
    The output should not include "$num5"
  End

  Example 'change quota for my project'
    When call taikun project quota edit "$pid" --disk-size "$num1" --server-cpu "$num2" --server-ram "$num3" --vm-cpu "$num4" --vm-ram "$num5" --vm-volume-size "$num6"
    The status should equal 0
    The lines of output should equal 1
    The output should include "Operation was successful."
  End

  Example 'list changed quota'
    When call list_project
    The status should equal 0
    The lines of output should equal 1
    The output should include "$num1"
    The output should include "$num2"
    The output should include "$num3"
    The output should include "$num4"
    The output should include "$num5"
    The output should include "$num6"
  End
End