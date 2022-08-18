Context 'accessprofile/allowedhost/add'
  setup() {
    profile_id=$(taikun access-profile add $(_rnd_name) -I)
    ip="192.168.1.1"
    host_id=$(taikun access-profile host add $profile_id -i $ip -m 16 -I)
  }

  cleanup() {
    taikun access-profile host delete $host_id -q 2>/dev/null || true
    taikun access-profile delete $profile_id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'basic allowed host'
    When call taikun access-profile host list $profile_id
    The output should include "$ip"
    The status should equal 0
  End

  Example 'duplicate ip address'
    When call taikun access-profile host add $profile_id -i $ip -m 16 
    The stderr should include '400'
    The stderr should include 'already exists'
    The status should equal 1
  End
  
End
