Context 'project/k8s/add'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs )
    ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
    flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    profilename="$(_rnd_name)"
    kid=$(taikun kubernetes-profile add "$profilename" -o "$oid" --enable-wasm --enable-octavia -I)
    pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" --flavors "$flavor" --kubernetes-profile-id "$kid" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete "$pid" -q 2>/dev/null; then
      taikun project delete --force "$pid" -q 2>/dev/null || true
    fi
    taikun kubernetes-profile delete "$kid" -q 2>/dev/null || true
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Context
    add_master() {
      msid=$(taikun project k8s add "$pid" --name tk-cli-master -r kubemaster -k foo=bar,bar=foo -f "$flavor" --enable-wasm -I | xargs)
    }
    BeforeAll 'add_master'

    remove_master() {
      taikun project k8s delete "$pid" --server-ids "$msid" -q 2>/dev/null || true
    }
    AfterAll 'remove_master'

    getwasm(){
      taikun project k8s list "$pid" --columns wasm
    }

    add_wasm_bastion() {
      taikun project k8s add "$pid" --name tk-cli-bastion -r bastion -f "$flavor" --enable-wasm
    }

    Example 'add one server'
      When call taikun project k8s list "$pid" --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include 'master'
    End

    Example 'See if wasm was enabled'
      When call getwasm
      The status should equal 0
      The lines of output should equal 3
      The output should include 'WASM'
      The output should include 'Yes'
    End

    Example 'add two servers with the same name'
      When call taikun project k8s add "$pid" --name tk-cli-master -r kubemaster -f "$flavor"
      The status should equal 1
      The stderr should include 'Duplicate name occurred'
    End

    Example 'Try to add wasm enabled bastion'
      When call add_wasm_bastion
      The status should equal 1
      The lines of output should equal 0
      The lines of stderr should equal 1
      The stderr should include "Wasm not available for bastion"
      The stderr should include "400"
    End

  End
End

Context 'project/k8s/add'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I | xargs)
    #flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --limit 1 -C name) # Selects m4.4xlarge (16 CPU, 64 RAM) which is total overkill for testmachine
    flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" --flavors "$flavor" -I | xargs)
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

  Context
    remove_master() {
      taikun project k8s delete "$pid" -a -q 2>/dev/null || true
    }
    AfterAll 'remove_master'

    Example 'add one server with availability zone'
      When call taikun project k8s add "$pid" -n master --flavor "$flavor" -r kubemaster -a a
      The status should equal 0
      The lines of output should equal 8
      The output should include 'master'
    End

    Example 'add one server with bad availability zone'
      When call taikun project k8s add "$pid" -n master2 --flavor "$flavor" -r kubemaster -a f
      The status should equal 1
      The stderr should include 'There is no zone f for this cloud credential'
    End
  End
End
