# shellcheck shell=bash

Context 'whoami'
  Example 'whoami does not fail'
    When call taikun whoami
    The status should equal 0
    The lines of output should equal 1
  End
End

Context 'rnd_name'
  setup() {
    username=$(_rnd_name)
  }
  BeforeAll 'setup'

  Example 'it works'
    When call echo "$username"
    The status should equal 0
    The output should include "tk-cli-"
    The output should not include "{a..z}"
  End
End