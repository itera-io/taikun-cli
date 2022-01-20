Context 'cmd/standaloneprofile/securitygroup'

  setup() {
    id=$(taikun standalone-profile add $(_rnd_name) --public-key "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy" -I)
    udp_id=$(taikun standalone-profile security-group add $id --name "doom" --protocol udp --min-port 666 --max-port 666 --remote-ip-prefix '192.0.2' -I)
    tcp_id=$(taikun standalone-profile security-group add $id --name "rsync" --protocol tcp --min-port 873 --max-port 873 --remote-ip-prefix '172.16.0' -I)
    icmp_id=$(taikun standalone-profile security-group add $id --name "icmp" --protocol icmp --remote-ip-prefix '192.168' -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun standalone-profile security-group delete $udp_id -q
    taikun standalone-profile security-group delete $tcp_id -q
    taikun standalone-profile security-group delete $icmp_id -q
    taikun standalone-profile delete $id -q
  }
  AfterAll 'cleanup'

  Example 'add, list and delete'
    When call taikun standalone-profile security-group list $id --no-decorate
    The status should equal 0
    The lines of output should equal 3
    The output should include 'doom'
    The output should include 'rsync'
    The output should include 'icmp'
  End

  Example 'calling add without standalone profile ID should cause error'
    When call taikun standalone-profile security-group add --name "doom" --protocol udp --min-port 666 --max-port 666 --remote-ip-prefix '192.0.2'
    The status should equal 1
    The stderr should equal 'Error: accepts 1 arg(s), received 0'
  End

  Example 'calling add without remote IP prefix should cause error'
    When call taikun standalone-profile security-group add --name "doom" --protocol udp --min-port 666 --max-port 666
    The status should equal 1
    The stderr should equal 'Error: required flag(s) "remote-ip-prefix" not set'
  End

  Example 'calling add without max port should cause error'
    When call taikun standalone-profile security-group add --name "doom" --protocol udp --min-port 666 --remote-ip-prefix '192.0.2'
    The status should equal 1
    The stderr should include 'max-port'
  End

  Example 'calling add without min port should cause error'
    When call taikun standalone-profile security-group add --name "doom" --protocol udp --max-port 666 --remote-ip-prefix '192.0.2'
    The status should equal 1
    The stderr should include 'min-port'
  End

  Example 'adding ICMP with port range should cause error'
    When call taikun standalone-profile security-group add --name "doom" --protocol udp --min-port 666 --max-port 666 --remote-ip-prefix '192.0.2'
    The status should equal 1
    The stderr should include 'port range'
  End

  Example 'duplicate name should cause error'
    When call taikun standalone-profile security-group add $id --name "doom" --protocol udp --min-port 666 --max-port 666 --remote-ip-prefix '192.0.2'
    The status should equal 1
    The stderr should include 'Please specify another name'
  End

End
