Context 'groups/organizations'
  setup() {
    accountID=`taikun accounts list --no-decorate | cut -d ' ' -f1 | xargs`
    ## Seeding group
    taikun groups create test-grp --account-id "$accountID" -q
    groupID=`taikun groups list "$accountID" --no-decorate | cut -d ' ' -f1 | xargs`
    ## Seeding organization
    orgname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q
    taikun groups delete "$groupID" -q
  }
  AfterAll 'cleanup'

  Example 'add'
    When call taikun groups organizations add "$groupID" --organization-id "$oid" --role "Member"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End

  Example 'update'
    When call taikun groups organizations update "$groupID" --organization-id "$oid" --role "Manager"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End

  Example 'delete'
    When call taikun groups organizations delete "$groupID" --organization-id "$oid"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End
End
