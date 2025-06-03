Context 'cloudcredential/google/list'
    # ---
    # --- The story test simulates the whole process from nothing to Google project with VM and k8s server (no commit) ---
    # ---
    setup() {
      orgname="$(_rnd_name)"
      oid=$(taikun organization add "$orgname" -f "$orgname" -I)

      echo "$GCP_CONFIG_FILE" > gcp.json
      ccname="$(_rnd_name)"
      ccid=$(taikun cloud-credential google add "$ccname" -z "$GCP_AZ_COUNT" -c ./gcp.json -r "$GCP_REGION" --import-project  -o "$oid" -I)

      flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 2 --max-cpu 4 --min-ram 4 --max-ram 8 -C name --limit 1 | xargs)
      image=$(taikun cc images "$ccid" --no-decorate -C id --limit 1 --google-image-type all --google-latest --show-large-values | xargs)
      image_name=$(taikun cc images "$ccid" --no-decorate -C name --limit 1 --google-image-type all --google-latest --show-large-values | xargs)

      k8sprofilename="$(_rnd_name)"
      k8sprofileid=$(taikun kubernetes-profile add "$k8sprofilename" --enable-taikun-lb -o "$oid" -I)

      projectname="$(_rnd_name)"
      projectid=$(taikun project add "$projectname" --cloud-credential-id "$ccid" --kubernetes-profile-id "$k8sprofileid" --flavors "$flavor" -I)

      taikun project image bind "$projectid" --image-ids "$image_name" -q
      standaloneprofile="$(_rnd_name)"
      pubkey="ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIHshx25CJGDd0HfOQqNt65n/970dsPt0y12lfKKO9fAs dummy"
      standaloneprofileid=$(taikun standalone-profile add "$standaloneprofile" --public-key "$pubkey" -o "$oid" -I)
      vmname="tk-cli1"
      vmid=$(taikun project vm add "$projectid" --name "$vmname" --flavor "$flavor" --image-id "$image" --standalone-profile-id "$standaloneprofileid" --volume-size 42 -I)

      servername="tk-cli2"
      serverid=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername" --role Kubeworker -I )
    }
    BeforeAll 'setup'

    cleanup() {
      taikun project vm delete "$projectid" --vm-ids "$vmid" -q 2>/dev/null || true
      taikun project k8s delete "$projectid" --server-ids "$serverid" -q 2>/dev/null || true
      taikun project delete "$projectid" -q 2>/dev/null || true
      taikun kubernetes-profile delete "$k8sprofileid" -q 2>/dev/null || true
      taikun standalone-profile delete "$standaloneprofileid" -q 2>/dev/null || true
      taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
      taikun organization delete "$oid" -q 2>/dev/null || true
      rm gcp.json
    }
    AfterAll 'cleanup'

    Example 'list google cloud credential'
      When call taikun cloud-credential google list -o "$oid" --no-decorate
      The lines of output should equal 1
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
      The output should include "Unlocked"
    End

    Example 'list all cloud credentials'
      When call taikun cloud-credential list -o "$oid" --no-decorate
      The status should equal 0
      The output should include "$orgname"
      The output should include "$ccname"
      The output should include "$ccid"
    End

    Example 'list google kubernetes profile'
      When call taikun kubernetes-profile list
      The output should include "$k8sprofileid"
      The output should include "$k8sprofilename"
      The status should equal 0
    End

    Example 'list google project'
      When call taikun project list
      The output should include "$projectname"
      The output should include "$projectid"
      The status should equal 0
    End

    Example 'list google vm'
      When call taikun project vm list "$projectid"
      The output should include "$vmname"
      The output should include "$vmid"
      The status should equal 0
    End

    Example 'list google server'
      When call taikun project k8s list "$projectid"
      The output should include "$servername"
      The output should include "$serverid"
      The status should equal 0
    End

End