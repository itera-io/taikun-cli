Context 'project/quota'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    projectname="$(_rnd_name)"
    pid=$(taikun project add "$projectname" -o "$oid" --cloud-credential-id "$ccid" -I)
    diskSize="$(_rnd_between 30 102400)"
    serverCpu="$(_rnd_between 2 1000000)"
    serverRam="$(_rnd_between 2 102400)"
    vmCpu="$(_rnd_between 1 1000000)"
    vmRam="$(_rnd_between 1 102400)"
    vmVolume="$(_rnd_between 1 102400)"
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
    The output should not include "$vmRam"
  End

  Example 'change quota for my project'
    When call taikun project quota edit "$pid" --disk-size "$diskSize" --server-cpu "$serverCpu" --server-ram "$serverRam" --vm-cpu "$vmCpu" --vm-ram "$vmRam" --vm-volume-size "$vmVolume"
    The status should equal 0
    The lines of output should equal 1
    The output should include "Operation was successful."
  End

  Example 'list changed quota'
    When call list_project
    The status should equal 0
    The lines of output should equal 1
    The output should include "$diskSize"
    The output should include "$serverCpu"
    The output should include "$serverRam"
    The output should include "$vmCpu"
    The output should include "$vmRam"
    The output should include "$vmVolume"
  End
End