Context 'cloudcredential/proxmox/list'
    # ---
    # --- The story test simulates the whole process from nothing to Proxmox project with VM and k8s server (no commit) ---
    # ---
    setup() {
      orgname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)

      ccname="$(_rnd_name)"
      ccid=$(taikun cloud-credential proxmox add "$ccname" --api-host "$PROXMOX_API_HOST" --client-id "$PROXMOX_CLIENT_ID" --client-secret "$PROXMOX_CLIENT_SECRET" --storage "$PROXMOX_STORAGE" --vm-template "$PROXMOX_VM_TEMPLATE_NAME" --hypervisors "$PROXMOX_HYPERVISOR,$PROXMOX_HYPERVISOR2" --continent "$PROXMOX_CONTINENT" --private-network "$PROXMOX_PRIVATE_NETWORK" --private-netmask "$PROXMOX_PRIVATE_NETMASK" --private-gateway "$PROXMOX_PRIVATE_GATEWAY" --private-begin-range "$PROXMOX_PRIVATE_BEGIN_RANGE" --private-end-range "$PROXMOX_PRIVATE_END_RANGE" --private-bridge "$PROXMOX_PRIVATE_BRIDGE" --public-network "$PROXMOX_PUBLIC_NETWORK" --public-netmask "$PROXMOX_PUBLIC_NETMASK" --public-gateway "$PROXMOX_PUBLIC_GATEWAY" --public-begin-range "$PROXMOX_PUBLIC_BEGIN_RANGE" --public-end-range "$PROXMOX_PUBLIC_END_RANGE" --public-bridge "$PROXMOX_PUBLIC_BRIDGE" -o "$oid" -I)

      flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 2 --max-cpu 4 --min-ram 4 --max-ram 8 -C name --limit 1 | xargs)
      image=$(taikun cc images "$ccid" --no-decorate -C id --limit 1)

      k8sprofilename="$(_rnd_name)"
      k8sprofileid=$(taikun kubernetes-profile add "$k8sprofilename" --enable-taikun-lb -o "$oid" -I)

      projectname="$(_rnd_name)"
      projectid=$(taikun project add "$projectname" --cloud-credential-id "$ccid" --kubernetes-profile-id "$k8sprofileid" --flavors "$flavor" -o "$oid" -I)

      taikun project image bind "$projectid" --image-ids "$image" -q
      standaloneprofile="$(_rnd_name)"
      pubkey="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy"
      standaloneprofileid=$(taikun standalone-profile add "$standaloneprofile" --public-key "$pubkey" -o "$oid" -I)
      vmname="tk-cli1"
      vmid=$(taikun project vm add "$projectid" --name "$vmname" --hypervisor "$PROXMOX_HYPERVISOR2" --flavor "$flavor" --image-id "$image" --standalone-profile-id "$standaloneprofileid" --volume-size 42 -I)

      servername="tk-cli2"
      serverid=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername" --proxmox-disk 42 --role Kubeworker -I )
    }
    BeforeAll 'setup'

    cleanup() {
      taikun project vm delete "$projectid" --vm-ids "$vmid" -q 2>/dev/null || true
      taikun project k8s delete "$projectid" --server-ids "$serverid" -q 2>/dev/null || true
      taikun project delete "$projectid" -q 2>/dev/null || true
      taikun kubernetes-profile delete "$k8sprofileid" -q 2>/dev/null || true
      taikun standalone-profile delete "$standaloneprofileid" -q 2>/dev/null || true
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Example 'list proxmox cloud credential'
      When call taikun cloud-credential proxmox list -o "$oid" --no-decorate
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'list all cloud credentials'
      When call taikun cloud-credential list -o "$oid" --no-decorate
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
    End

    Example 'list proxmox kubernetes profile'
      When call taikun kubernetes-profile list
      The output should include "$k8sprofileid"
      The output should include "$k8sprofilename"
      The status should equal 0
    End

    Example 'list proxmox project'
      When call taikun project list
      The output should include "$projectname"
      The output should include "$projectid"
      The status should equal 0
    End

    Example 'list proxmox vm'
      When call taikun project vm list "$projectid"
      The output should include "$vmname"
      The output should include "$vmid"
      The status should equal 0
    End

    Example 'list proxmox server'
      When call taikun project k8s list "$projectid"
      The output should include "$servername"
      The output should include "$serverid"
      The status should equal 0
    End

End