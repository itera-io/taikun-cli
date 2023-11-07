Context 'root/version'
  Example 'Get CLI version'
    When call taikun --version
    The status should equal 0
    The lines of output should equal 1
    The output should include "Taikun CLI version"
  End

  Example 'Get CLI version in subcommand'
    When call taikun whoami --version
    The status should equal 1
    The lines of stderr should equal 1
    The stderr should include "Error: unknown flag: --version"
  End
End