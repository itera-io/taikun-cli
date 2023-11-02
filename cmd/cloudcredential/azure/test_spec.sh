Context 'cloudcredential/azure'
  setup() {
    orgname="$(_rnd_name)"
    ccname="$(_rnd_name)"
    oid=$(taikun organization add "$orgname" -f "$orgname" -I)
    ccid=$(taikun cloud-credential azure add "$ccname" -o "$oid" --az-count="$AZ_COUNT" --client-id="$AZURE_CLIENT_ID" --client-secret="$AZURE_SECRET" --location="$AZURE_LOCATION" --subscription-id="$AZURE_SUBSCRIPTION" --tenant-id="$AZURE_TENANT" -I)
  }
  BeforeAll 'setup'

  cleanup() {
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  list_cc(){
    taikun cloud-credential list -o "$oid" --no-decorate
  }

  Example 'list azure cloud credential'
    When call list_cc
    The lines of output should equal 1
    The status should equal 0
    The output should include "$orgname"
    The output should include "$ccname"
    The output should include "$ccid"
    The output should include "Unlocked"
  End

  Example 'same name azure cloud credential'
    When call taikun cloud-credential azure add "$ccname" -o "$oid" --az-count="$AZ_COUNT" --client-id="$AZURE_CLIENT_ID" --client-secret="$AZURE_SECRET" --location="$AZURE_LOCATION" --subscription-id="$AZURE_SUBSCRIPTION" --tenant-id="$AZURE_TENANT"
    The lines of output should equal 0
    The lines of stderr should equal 1
    The status should equal 1
    The stderr should include "$ccname already exists"
  End

  Example 'check valid azure cloud credential'
    When call taikun cloud-credential azure check --client-id="$AZURE_CLIENT_ID" --client-secret="$AZURE_SECRET" --tenant-id="$AZURE_TENANT"
    The lines of output should equal 1
    The status should equal 0
    The output should include "Azure cloud credential is valid."
  End

  Example 'check invalid azure cloud credential'
    When call taikun cloud-credential azure check --client-id="$AZURE_CLIENT_ID" --client-secret="mockup.local" --tenant-id="$AZURE_TENANT"
    The lines of stderr should equal 1
    The status should equal 1
    The stderr should include "Azure cloud credential is not valid."
  End

  # Publishers
  Example 'list publishers'
    When call taikun cloud-credential azure publishers "$ccid" --limit 2
    The lines of output should equal 2
    The status should equal 0
  End

  # List offers from the first listed publisher but dont print them (cannot guarantee how many there will be)
  Example 'list offers from publisher'
    When call taikun cloud-credential azure offers "$ccid" --limit 1 -p "`taikun cloud-credential azure publishers $ccid --limit 1`" -q
    The status should equal 0
  End

End