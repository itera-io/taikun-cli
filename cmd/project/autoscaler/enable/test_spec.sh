Context 'project/autoscaler/enable'
    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid"  -o "$oid" --flavors "$AUTOSCALER_FLAVOR" -I)
        AUTOSCALER_NAME="auto"
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

    Example 'enable spot autoscaling on openstack'
        When call taikun project autoscaler enable "$pid" -n "$AUTOSCALER_NAME" -f "$AUTOSCALER_FLAVOR" --max-size 4 --min-size 1 --spot-enable
        The status should equal 1
        The stderr should include 'There is no spot'
    End

    Example 'enable autoscaling succesfully'
        When call taikun project autoscaler enable "$pid" -n "$AUTOSCALER_NAME" -f "$AUTOSCALER_FLAVOR" --max-size 4 --min-size 1
        The status should equal 0
        The output should include 'Operation was successful'
    End

    Example 'enable autoscaling succesfully'
        When call taikun project info "$pid" --all-columns
        The output should include "$AUTOSCALER_FLAVOR"
        The output should include "$AUTOSCALER_NAME"
        The output should include " 1 "
        The output should include " 4 "
        The status should equal 0
    End

    Example 'enable two times'
        When call taikun project autoscaler enable "$pid" -n "$(_rnd_name)" -f "$AUTOSCALER_FLAVOR"
        The status should equal 1
        The stderr should include 'Project autoscaling already enabled'
    End
End