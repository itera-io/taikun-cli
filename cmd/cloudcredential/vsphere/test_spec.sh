Context 'cloudcredential/vsphere'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid=$(taikun cloud-credential vsphere add "$ccname" --username "$VSPHERE_USERNAME" --password "$VSPHERE_PASSWORD" --url "$VSPHERE_API_URL" --datacenter "$VSPHERE_DATACENTER" --resource-pool "$VSPHERE_RESOURCE_POOL" --data-store "$VSPHERE_DATA_STORE" --drs-enabled  --vm-template "$VSPHERE_VM_TEMPLATE" --continent "$VSPHERE_CONTINENT" --public-network-name "$VSPHERE_PUBLIC_NETWORK_NAME" --public-network "$VSPHERE_PUBLIC_NETWORK_ADDRESS" --public-netmask "$VSPHERE_PUBLIC_NETMASK" --public-gateway "$VSPHERE_PUBLIC_GATEWAY" --public-begin-range "$VSPHERE_PUBLIC_BEGIN_RANGE" --public-end-range "$VSPHERE_PUBLIC_END_RANGE" --private-network-name "$VSPHERE_PRIVATE_NETWORK_NAME" --private-network "$VSPHERE_PRIVATE_NETWORK_ADDRESS" --private-netmask "$VSPHERE_PRIVATE_NETMASK" --private-gateway "$VSPHERE_PRIVATE_GATEWAY" --private-begin-range "$VSPHERE_PRIVATE_BEGIN_RANGE" --private-end-range "$VSPHERE_PRIVATE_END_RANGE" --organization "$oid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    list_cc(){
      taikun cloud-credential vsphere list -o "$oid" --no-decorate
    }

    list_cc_all(){
      taikun cloud-credential list -o "$oid" --no-decorate
    }

    Example 'list vsphere cloud credential'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'list all cloud credentials'
      When call list_cc_all
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'check valid vsphere cloud credential'
      When call taikun cloud-credential vsphere check -u "$VSPHERE_API_URL" -i "$VSPHERE_USERNAME" -s "$VSPHERE_PASSWORD"
      The lines of output should equal 1
      The status should equal 0
      The output should include "vSphere cloud credential is valid."
    End

    Example 'check invalid vsphere cloud credential'
      When call taikun cloud-credential proxmox check -u "$VSPHERE_API_URL" -i "mockup" -s "$VSPHERE_PASSWORD"
      The lines of stderr should equal 1
      The status should equal 1
      The stderr should include "HTTP 400"
    End

End