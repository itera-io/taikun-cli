Context 'cmd/standaloneprofile'

  setup() {
    oid=$(taikun organization add $(_rnd_name) --full-name $(_rnd_name) -I)

    name=$(_rnd_name)
    pubkey="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy"
    id=$(taikun standalone-profile add $name --public-key "$pubkey" -o $oid -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun standalone-profile delete $id -q
    taikun organization delete $oid -q
  }
  AfterAll 'cleanup'

  Example 'add, list and delete'
    When call taikun standalone-profile list -o $oid --no-decorate --show-large-values
    The status should equal 0
    The lines of output should equal 1
    The output should include "$name"
    The output should include "$pubkey"
  End

  Example 'duplicate name should cause error'
    When call taikun standalone-profile add $name --public-key "$pubkey" -o $oid
    The status should equal 1
    The stderr should include 'Please specify another name'
  End

  Example 'calling add without name should cause error'
    When call taikun standalone-profile add --public-key "$pubkey" -o $oid
    The status should equal 1
    The stderr should equal 'Error: accepts 1 arg(s), received 0'
  End

  Example 'calling add without public key should cause error'
    When call taikun standalone-profile add $name -o $oid
    The status should equal 1
    The stderr should equal 'Error: required flag(s) "public-key" not set'
  End

  Context
    lock() {
      taikun standalone-profile lock $id -q
    }
    BeforeEach 'lock'

    unlock() {
      taikun standalone-profile unlock $id -q
    }
    AfterEach 'unlock'

    Example 'lock then unlock'
      When call taikun standalone-profile list -o $oid --columns lock --no-decorate
      The status should equal 0
    The lines of output should equal 1
      The output should include 'Locked'
    End
  End

  Context 'rename'
    give_new_name() {
      new_name=$(_rnd_name)
      taikun standalone-profile rename $id --name $new_name -q
    }
    BeforeEach 'give_new_name'

    restore_old_name() {
      taikun standalone-profile rename $id --name $name -q
    }
    AfterEach 'restore_old_name'

    Example 'rename then restore old name'
      When call taikun standalone-profile list -o $oid --columns name --no-decorate
      The status should equal 0
      The lines of output should equal 1
      The output should include "$name"
    End
  End
End
