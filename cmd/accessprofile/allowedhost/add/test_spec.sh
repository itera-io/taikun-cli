Context 'accessprofile/allowedhost/add'
  setup() {
    profile_id=$(taikun access-profile add $(_rnd_name) -I)
  }

  cleanup() {
    taikun access-profile host delete $host_id -q 2>/dev/null || true
    taikun access-profile delete $profile_id -q 2>/dev/null || true
  }

  BeforeEach 'setup'
  AfterEach 'cleanup'

  Example 'basic allowed host'
    run() {
      ip="192.168.1.1"
      host_id=$(taikun access-profile host add $profile_id -ip $ip -m 16 -I)
      taikun access-profile host list | grep $host_id
    }

    When call run
    The output should include "$ip"
    The status should equal 0
  End

  Context
    add_allowed_host() {
      host_id=$(taikun access-profile host add $profile_id -ip $ip -m 16 -I)
    }
    Before 'add_access_profile'

    Example 'duplicate ip address'
      When call taikun access-profile host add $profile_id -ip $ip -m 16 
      The stderr should include '400'
      The stderr should include 'already exists'
      The status should equal 1
    End
  End
End
