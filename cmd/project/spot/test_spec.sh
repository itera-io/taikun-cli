Context 'project/spot'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I | xargs)
    flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" --flavors "$flavor" --spot-full -I | xargs)
  }
  BeforeAll 'setup'

  cleanup() {
    if ! taikun project delete "$pid" -q 2>/dev/null; then
      taikun project delete --force "$pid" -q 2>/dev/null || true
    fi
    taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
    taikun organization delete "$oid" -q 2>/dev/null || true
  }
  AfterAll 'cleanup'

  Context
    Example 'Enable full spot twice'
      When call taikun project spot full --enable "$pid"
      The status should equal 1
      The lines of stderr should equal 1
      The stderr should include 'Taikun Error:'
    End

    Example 'List if full spots are set correctly'
      When call taikun project info --columns SPOT-FULL "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

    Example 'Disable full spot'
      When call taikun project spot full --disable "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'List if full spots are set correctly'
      When call taikun project info --columns SPOT-FULL "$pid"
      The status should equal 0
      The output should include 'No'
    End

    Example 'Enable worker spot'
      When call taikun project spot worker --enable "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'Enable full spot while worker spot is enabled'
      When call taikun project spot full --enable  "$pid"
      The status should equal 1
      The stderr should include 'Spot workers option enabled for this project.'
    End

    Example 'List if worker spots are set correctly'
      When call taikun project info --columns SPOT-WORKERS "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

    Example 'Disable worker spot'
      When call taikun project spot worker --disable "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'List if worker spots are set correctly'
      When call taikun project info --columns SPOT-WORKERS "$pid"
      The status should equal 0
      The output should include 'No'
    End

    Example 'List if VMS spots are set correctly'
      When call taikun project info --columns  SPOT-VMS "$pid"
      The status should equal 0
      The output should include 'No'
    End

    Example 'Enable VMS spot while worker spot is enabled'
      When call taikun project spot vms --enable  "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'List if VMS spots are set correctly'
      When call taikun project info --columns  SPOT-VMS "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

  End
End
