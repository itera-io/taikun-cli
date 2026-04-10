Context 'catalog'
  setup() {
    orgname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
    ## Creating catalog managed-apps
    taikun catalog create "managed-apps" -d "managed-apps" -O "$oid"
    defcatid=`taikun catalog list -O "$oid" --no-decorate | grep "managed-apps" | cut -d ' ' -f1 | xargs`
    ## Making managed-apps catalog to be default one
    taikun catalog make-default "$defcatid"
  }
  BeforeAll 'setup'

  cleanup() {
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_cat(){
    taikun catalog list -O "$oid" | grep "$orgname-catalog"
  }

  makecatdefault(){
      catid=`taikun catalog list -O "$oid" --no-decorate | grep "$orgname-catalog" | cut -d ' ' -f1 | xargs`
      taikun catalog make-default "$catid"
  }

  unmakedefault(){
      managedid=`taikun catalog list -O "$oid" --no-decorate | grep "managed-apps" | cut -d ' ' -f1 | xargs`
      taikun catalog make-default "$managedid"
  }

  deletecat(){
      catid=`taikun catalog list -O "$oid" --no-decorate | grep "$orgname-catalog" | cut -d ' ' -f1 | xargs`
      taikun catalog delete "$catid"
  }

  Example 'create catalog'
    When call taikun catalog create "$orgname-catalog" -d "$orgname-catalog" -O "$oid"
    The lines of output should equal 1
    The status should equal 0
    The output should include 'Operation was successful.'
  End

  Example 'list catalog 1'
    When call list_cat
    The lines of output should equal 1
    The status should equal 0
    The output should include "$orgname-catalog"
    The output should include "No"
  End

  Example 'make catalog default'
    When call makecatdefault
    The status should equal 0
    The lines of output should equal 1
    The output should include 'Operation was successful.'
  End

  Example 'list catalog 2'
    When call list_cat
    The lines of output should equal 1
    The status should equal 0
    The output should include "$orgname-catalog"
    The output should include "Yes"
  End

  Example 'delete default catalog'
    When call deletecat
    The lines of stderr should equal 1
    The status should equal 1
    The stderr should include "400"
    The stderr should include "can not delete default catalog"
  End

  Example 'not make it default again'
    When call unmakedefault
    The status should equal 0
    The output should include "Operation was successful"
  End

  Example 'delete catalog'
    When call deletecat
    The lines of output should equal 1
    The status should equal 0
    The output should include "catalog"
    The output should include "$catid"
    The output should include "delete"
  End

  Example 'list catalog 3'
    When call taikun catalog list -O "$oid" --no-decorate
    The lines of output should equal 1
    The status should equal 0
    The output should not include "$orgname-catalog"
  End

End