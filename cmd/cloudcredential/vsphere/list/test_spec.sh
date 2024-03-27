Context 'cloudcredential/vsphere/list'
    # ---
    # --- The story test simulates the whole process from nothing to Proxmox project with VM and k8s server (no commit) ---
    # ---
    setup() {
      orgname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)

      ccname="$(_rnd_name)"
      ccid=$(taikun cloud-credential vsphere add "$ccname" --username "$VSPHERE_USERNAME" --password "$VSPHERE_PASSWORD" --url "$VSPHERE_API_URL" --datacenter "$VSPHERE_DATACENTER" --resource-pool "$VSPHERE_RESOURCE_POOL" --data-store "$VSPHERE_DATA_STORE" --drs-enabled  --vm-template "$VSPHERE_VM_TEMPLATE" --continent "$VSPHERE_CONTINENT" --public-network-name "$VSPHERE_PUBLIC_NETWORK_NAME" --public-network "$VSPHERE_PUBLIC_NETWORK_ADDRESS" --public-netmask "$VSPHERE_PUBLIC_NETMASK" --public-gateway "$VSPHERE_PUBLIC_GATEWAY" --public-begin-range "$VSPHERE_PUBLIC_BEGIN_RANGE" --public-end-range "$VSPHERE_PUBLIC_END_RANGE" --private-network-name "$VSPHERE_PRIVATE_NETWORK_NAME" --private-network "$VSPHERE_PRIVATE_NETWORK_ADDRESS" --private-netmask "$VSPHERE_PRIVATE_NETMASK" --private-gateway "$VSPHERE_PRIVATE_GATEWAY" --private-begin-range "$VSPHERE_PRIVATE_BEGIN_RANGE" --private-end-range "$VSPHERE_PRIVATE_END_RANGE" --organization "$oid" -I)

      flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 2 --max-cpu 4 --min-ram 4 --max-ram 8 -C name --limit 1 | xargs)
      image=$(taikun cc images "$ccid" --no-decorate -C id --limit 1 | xargs)

      projectname="$(_rnd_name)"
      projectid=$(taikun project add "$projectname" --cloud-credential-id "$ccid" --flavors "$flavor" -o "$oid" -I)

      taikun project image bind "$projectid" --image-ids "$image" -q
      standaloneprofile="$(_rnd_name)"
      pubkey="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy"
      standaloneprofileid=$(taikun standalone-profile add "$standaloneprofile" --public-key "$pubkey" -o "$oid" -I)
      vmname="tk-cli1"
      vmid=$(taikun project vm add "$projectid" --name "$vmname" --flavor "$flavor" --image-id "$image" --standalone-profile-id "$standaloneprofileid" --volume-size 42 -I)

      servername="tk-cli2"
      serverid=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername" --role Kubeworker -I )
    }
    BeforeAll 'setup'

    cleanup() {
      taikun project vm delete "$projectid" --vm-ids "$vmid" -q 2>/dev/null || true
      taikun project k8s delete "$projectid" --server-ids "$serverid" -q 2>/dev/null || true
      taikun project delete "$projectid" -q 2>/dev/null || true
      taikun standalone-profile delete "$standaloneprofileid" -q 2>/dev/null || true
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    Example 'list vsphere cloud credential'
      When call taikun cloud-credential vsphere list -o "$oid" --no-decorate
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

    Example 'list vsphere project'
      When call taikun project list
      The output should include "$projectname"
      The output should include "$projectid"
      The status should equal 0
    End

    Example 'list vsphere vm'
      When call taikun project vm list "$projectid"
      The output should include "$vmname"
      The output should include "$vmid"
      The output should inculde "$flavor"
      The output should inculde "$image"
      The status should equal 0
    End

    Example 'list vsphere server'
      When call taikun project k8s list "$projectid"
      The output should include "$servername"
      The output should include "$serverid"
      The status should equal 0
    End

End