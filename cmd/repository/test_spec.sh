Context 'repository/'
  Example 'list the recommended repositories'
    When call taikun repository list-recommend
    The status should equal 0
    The output should include "REPOSITORY-NAME"
    The output should include "taikun-managed-apps"
  End

  Example 'list the public repositories'
    When call taikun repository list-public --limit 1 --no-decorate --sort-by REPOSITORY-NAME
    The status should equal 0
    The lines of output should equal 1
  End

# We cannot guarantee there will be some private repos
#  Example 'list the private repositories'
#    When call taikun repository list-private --limit 1 --no-decorate --sort-by REPOSITORY-NAME
#    The status should equal 0
#    The stdout should not be blank
#  End
End

Context 'repostory/'
  setup(){
    cname="$(_rnd_name)"
    oid=$(taikun organization add "$cname" --full-name "$(_rnd_name)" -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  askforrepo(){
    taikun repository list-private --no-decorate -o "$oid" | grep "tk-cli-susetest"
  }

  Example 'import private repo'
    When call taikun repository import tk-cli-susetest --url https://kubernetes-charts.suse.com -o "$oid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list the new repository'
    When call taikun repository list-private --no-decorate -o "$oid"
    The status should equal 0
    The output should include "$cname"
    The output should include "Enabled"
    The output should include "tk-cli-susetest"
  End

  Example 'disable the new repo'
    When call taikun repository disable tk-cli-susetest "$cname" -o "$oid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list the disabled repository'
    When call askforrepo
    The status should equal 0
    The output should include "$cname"
    The output should include "Disabled"
    The output should include "tk-cli-susetest"
  End

  Example 'enable the new repo'
    When call taikun repository enable tk-cli-susetest "$cname" -o "$oid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list the enabled repository'
    When call askforrepo
    The status should equal 0
    The output should include "$cname"
    The output should include "Enabled"
    The output should include "tk-cli-susetest"
  End

  Example 'disable the repo for deletion'
    When call taikun repository disable tk-cli-susetest "$cname" -o "$oid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'delete the repository'
    When call taikun repository delete tk-cli-susetest "$cname" -o "$oid"
    The status should equal 0
    The output should include "Operation was successful."
  End

  Example 'list without repo'
    When call taikun repository list-private --no-decorate -o "$oid"
    The status should equal 0
    The output should not include "$cname"
  End

End
