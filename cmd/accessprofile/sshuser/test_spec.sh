Context 'accessprofile/sshuser'
  setup() {
    apname="$(_rnd_name)"
    sshname="$(_rnd_name)"
    oid=$(taikun organization add "$(_rnd_name)" --full-name "$(_rnd_name)" -I | xargs)
    apid=$(taikun access-profile add "$apname" -o "$oid" -I | xargs )
  }

  cleanup() {
    taikun access-profile delete "$apid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'Add ssh profile'
    When call taikun access-profile ssh-user add "$apid" --name "$sshname" --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy"
    The output should include "$sshid"
    The output should include "$sshname"
    The status should equal 0
  End

  Context 'SSH user added to acces-profile'
    add_ssh_user() {
      sshid=$(taikun access-profile ssh-user add "$apid" --name "$sshname" --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I | xargs)
    }
    BeforeEach 'add_ssh_user'

    delete_ssh_user() {
      taikun access-profile ssh-user delete "$sshid" -q 2>/dev/null || true
    }
    AfterEach 'delete_ssh_user'


    Example 'List the ssh user'
      When call taikun access-profile ssh-user list "$apid" --no-decorate
      The output should include "$sshid"
      The output should include "$sshname"
      The lines of output should equal 1
      The status should equal 0
    End

    Example 'Delete the ssh user'
      When call taikun access-profile ssh-user delete "$sshid"
      The output should include "SSH user with ID $sshid was deleted successfully."
      The lines of output should equal 1
      The status should equal 0
    End

  End

End