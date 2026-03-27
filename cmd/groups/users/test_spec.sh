Context 'groups/users'
  setup() {
    accountID=`taikun accounts list --no-decorate | cut -d ' ' -f1 | xargs`
    ## Seeding group
    taikun groups create test-grp --account-id "$accountID" -q
    groupID=`taikun groups list "$accountID" --no-decorate | cut -d ' ' -f1 | xargs`
    ## Seeding organization
    orgname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
    ## Seeding user
    username="$(_rnd_name)"
    uid=$(taikun user add "$username" --email "${username}@cloudera.com" -O "$oid" -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun user delete "$uid" -q
    taikun organization delete "$oid" -q
    taikun groups delete "$groupID" -q
  }
  AfterAll 'cleanup'

  Example 'add'
    When call taikun groups users add "$groupID" --user-id "$uid"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End

  Example 'delete'
    When call taikun groups users delete "$groupID" --user-id "$uid"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End
End
