Context 'project/quota'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    projectname="$(_rnd_name)"
    profilename="$(_rnd_name)"
    kid=$(taikun kubernetes-profile add "$profilename" -o "$oid" --enable-wasm --enable-octavia -I)
    pid=$(taikun project add "$projectname" -o "$oid" --cloud-credential-id "$ccid" --kubernetes-profile-id "$kid" -I)
  }
  Before 'setup'

  cleanup() {
    if ! taikun project delete "$pid" -q 2>/dev/null; then
      taikun project delete --force "$pid" -q 2>/dev/null || true
    fi
    taikun kubernetes-profile delete "$kid" -q 2>/dev/null || true
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  After 'cleanup'

  getwasm(){
    taikun project info "$pid" | grep "WASM"
  }

  Example 'Get detailed info about project with wasm enabled'
    When call getwasm
    The status should equal 0
    The lines of output should equal 1
    The output should include "Yes"
  End

  Example 'Get detailed info about project with wasm enabled'
    When call taikun project info "$pid"
    The status should equal 0
    The lines of output should equal 33
    The output should include "$oid"
    The output should include "$ccid"
    The output should include "$projectname"
    The output should include "$profilename"
    The output should include "$kid"
    The output should include "$pid"
    The output should include "OpenStack"
    The output should include "Unlocked"
  End

End