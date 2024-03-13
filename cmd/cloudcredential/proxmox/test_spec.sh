Context 'cloudcredential/proxmox'
    setup() {
      orgname="$(_rnd_name)"
      ccname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)
      ccid=$(taikun cloud-credential proxmox add "$ccname" --api-host "$PROXMOX_API_HOST" --client-id "$PROXMOX_CLIENT_ID" --client-secret "$PROXMOX_CLIENT_SECRET" --storage "$PROXMOX_STORAGE" --vm-template "$PROXMOX_VM_TEMPLATE_NAME" --hypervisors "$PROXMOX_HYPERVISOR,$PROXMOX_HYPERVISOR2" --continent "$PROXMOX_CONTINENT" --private-network "$PROXMOX_PRIVATE_NETWORK" --private-netmask "$PROXMOX_PRIVATE_NETMASK" --private-gateway "$PROXMOX_PRIVATE_GATEWAY" --private-begin-range "$PROXMOX_PRIVATE_BEGIN_RANGE" --private-end-range "$PROXMOX_PRIVATE_END_RANGE" --private-bridge "$PROXMOX_PRIVATE_BRIDGE" --public-network "$PROXMOX_PUBLIC_NETWORK" --public-netmask "$PROXMOX_PUBLIC_NETMASK" --public-gateway "$PROXMOX_PUBLIC_GATEWAY" --public-begin-range "$PROXMOX_PUBLIC_BEGIN_RANGE" --public-end-range "$PROXMOX_PUBLIC_END_RANGE" --public-bridge "$PROXMOX_PUBLIC_BRIDGE" -o "$oid" -I)
    }
    BeforeAll 'setup'

    cleanup() {
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
    }
    AfterAll 'cleanup'

    list_cc(){
      taikun cloud-credential proxmox list -o "$oid" --no-decorate
    }

    list_cc_all(){
      taikun cloud-credential list -o "$oid" --no-decorate
    }

    Example 'list proxmox cloud credential'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'list all cloud credentials'
      When call list_cc
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'check valid proxmox cloud credential'
      When call taikun cloud-credential proxmox check "$ccname" -u "$PROXMOX_API_HOST" -i "$PROXMOX_CLIENT_ID" -s "$PROXMOX_CLIENT_SECRET"
      The lines of output should equal 1
      The status should equal 0
      The output should include "Proxmox cloud credential is valid."
    End

    Example 'check invalid proxmox cloud credential'
      When call taikun cloud-credential proxmox check "$ccname" -u "$PROXMOX_API_HOST" -i "$PROXMOX_CLIENT_ID" -s "mockup"
      The lines of stderr should equal 1
      The status should equal 1
      The stderr should include "Error: Proxmox cloud credential is not valid."
    End

End