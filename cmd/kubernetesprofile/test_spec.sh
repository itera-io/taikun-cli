Context 'kubernetesprofile'

  setup() {
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I)
  }

  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  AfterAll 'cleanup'

  Example 'No kubernetes profiles'
    When call taikun kubernetes-profile list -o "$oid" --no-decorate
    The status should equal 0
    The lines of output should equal 1 # counting the default
  End

  Context 'add/delete'
    add_config() {
      profile_name="$(_rnd_name)"
      profile_name2="$(_rnd_name)"
      scid=$(taikun kubernetes-profile add "$profile_name" -o "$oid" --enable-octavia --enable-gpu --enable-wasm -I)
      scid2=$(taikun kubernetes-profile add "$profile_name2" -o "$oid" --enable-octavia --enable-gpu --enable-wasm --proxmox-storage "OpenEBS" -I)
    }

    Before 'add_config'

    delete_config() {
      taikun kubernetes-profile delete "$scid" -q
      taikun kubernetes-profile delete "$scid2" -q
    }

    After 'delete_config'

    Example 'add then delete kubernetes profile'
      When call taikun kubernetes-profile list -o "$oid" --no-decorate
      The status should equal 0
      The lines of output should equal 3 # counting the default
      The output should include "$scid"
      The output should include "$scid2"
      The output should include "$profile_name"
      The output should include "$profile_name2"
      The output should include "OpenEBS"
    End

    Example 'Check if GPU got enabled'
      When call taikun kubernetes-profile list -o "$oid" --columns=nvidia-gpu --no-decorate
      The status should equal 0
      The lines of output should equal 3 # counting the default
      The output should include "Yes"
    End

    Example 'Check if Wasm got enabled'
      When call taikun kubernetes-profile list -o "$oid" --columns=wasm --no-decorate
      The status should equal 0
      The lines of output should equal 3 # counting the default
      The output should include "Yes"
    End
  End


    Context 'lock/unlock'
      add_config() {
        ppid=$(taikun kubernetes-profile add "$(_rnd_name)" -o "$oid" --enable-octavia -I)
        taikun kubernetes-profile lock "$ppid" -q
      }

      Before 'add_config'

      delete_config() {
        taikun kubernetes-profile delete "$ppid" -q
      }

      After 'delete_config'

      Example 'lock with already locked'
        When call taikun kubernetes-profile lock "$ppid"
        The status should equal 1
        The stderr should include "already lock"
      End

      Example 'unlock'
        When call taikun kubernetes-profile unlock "$ppid"
        The status should equal 0
        The output should include "Operation was successful"
      End

    End


End
