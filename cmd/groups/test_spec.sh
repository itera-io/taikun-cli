Context 'groups'
  setup() {
    accountID=`taikun accounts list --no-decorate | cut -d ' ' -f1 | xargs`
  }
  BeforeAll 'setup'

  Example 'group add'
    When call taikun groups create test-grp --account-id "$accountID"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful.'
  End

  Example 'group list'
    When call taikun groups list "$accountID" --no-decorate
    The lines of output should equal 1
    The status should equal 0
    The output should include 'test-grp'
  End

  Example 'group check duplicate exists'
    When call taikun groups check-duplicate-entity "$accountID" --name "test-grp"
    The lines of output should equal 0
    The status should equal 1
    The stderr should include 'already exists'
  End

  Example 'group check duplicate does NOT exist'
    When call taikun groups check-duplicate-entity "$accountID" --name "test-grpp"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End

  ## Extracting group ID
  groupID=`taikun groups list "$accountID" --no-decorate | cut -d ' ' -f1 | xargs`

  Example 'group update'
    When call taikun groups update "$groupID" --name "test-group" --claim-value "true"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful'
  End

  Example 'group list'
    When call taikun groups list "$accountID" --no-decorate
    The lines of output should equal 1
    The status should equal 0
    The output should include 'test-group'
    The output should include 'true'
  End

  Example 'group delete'
    When call taikun groups delete "$groupID"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'deleted successfully'
  End

End
