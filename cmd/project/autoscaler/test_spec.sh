Context 'project/autoscaler'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I | xargs)
    flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -o "$oid" --flavors "$flavor" --spot-vms -I | xargs)
    AUTOSCALER_NAME="autosc"
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
    Example 'Enable autoscaler with spot'
      When call taikun project autoscaler enable "$pid" -n "$AUTOSCALER_NAME" -f "$flavor" --max-size 4 --min-size 1 --spot-enable
      The status should equal 1
      The lines of stderr should equal 1
      The stderr should include 'full spot or worker spot should be enabled'
    End

    Example 'Disable project spot'
      When call taikun project spot vms --disable "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'Enable worker spot'
      When call taikun project spot worker --enable "$pid"
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'Enable autoscaler with spot'
      When call taikun project autoscaler enable "$pid" -n "$AUTOSCALER_NAME" -f "$flavor" --max-size 4 --min-size 1 --spot-enable
      The status should equal 0
      The lines of output should equal 1
      The output should include 'Operation was successful'
    End

    Example 'List if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-SPOT "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

    Example 'disable project spots while spot-autoscaler is running'
      When call taikun project spot worker "$pid" --disable
      The status should equal 1
      The stderr should include 'Taikun Error:'
    End

  End
End



Context 'project/autoscaler/aws-project'
  setup() {
    oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs)
    ccid=$(taikun cloud-credential aws add "$(_rnd_name)" -a "$AWS_ACCESS_KEY_ID" -s "$AWS_SECRET_ACCESS_KEY" -r "$AWS_DEFAULT_REGION" -z 1 -o "$oid" -I | xargs)
    flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --min-cpu 4 --max-cpu 4 --min-ram 8 --max-ram 8 -C name --limit 1 | xargs)
    AUTOSCALER_NAME="autos"
    pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -o "$oid" --flavors "$flavor" --spot-full --autoscaler-name "$AUTOSCALER_NAME" --autoscaler-min-size 1 --autoscaler-max-size 3 --autoscaler-flavor "$flavor" --autoscaler-disk-size 31 --autoscaler-spot -I | xargs)
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
    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-GROUP "$pid"
      The status should equal 0
      The output should include "$AUTOSCALER_NAME"
    End

    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-MIN "$pid"
      The status should equal 0
      The output should include '1'
    End

    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-MAX "$pid"
      The status should equal 0
      The output should include '3'
    End

    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-FLAVOR "$pid"
      The status should equal 0
      The output should include "$flavor"
    End

    Example 'list if autoscaler spots are set correctly'
      When call taikun project info --columns AUTOSCALING-SPOT "$pid"
      The status should equal 0
      The output should include 'Yes'
    End

  End
End
