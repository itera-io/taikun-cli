Context 'cloudcredential/google/add'

  Example 'Cannot set both import-project and billing-account-id'
    When call taikun cloud-credential google add g --region g --zone g --import-project --billing-account-id "foo" 
    The stderr should include "The flags --import-project and --billing-account-id are mutually exclusive"
    The status should equal 1
  End

  Example 'Cannot set both import-project and folder-id'
    When call taikun cloud-credential google add g --region g --zone g --import-project --folder-id "foo" 
    The stderr should include "The flags --import-project and --folder-id are mutually exclusive"
    The status should equal 1
  End

  Example 'Must set billing-account-id if not importing a project'
    When call taikun cloud-credential google add g --region g --zone g --folder-id "foo" 
    The stderr should include "Must set --billing-acount-id if not importing a project"
    The status should equal 1
  End

  Example 'Must set folder-id if not importing a project'
    When call taikun cloud-credential google add g --region g --zone g --billing-account-id "foo" 
    The stderr should include "Must set --folder-id if not importing a project"
    The status should equal 1
  End

End
