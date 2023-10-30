Context 'project/autoscaler/edit'
    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid"  -o "$oid" --autoscaler  --autoscaler-name "$(_rnd_name)" --autoscaler-flavor "$AUTOSCALER_FLAVOR" -I)
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

    Example 'edit autoscaling succesfully'
        When call taikun project autoscaler edit "$pid" --max-size 10 --min-size 2
        The status should equal 0
        The output should include 'Operation was successful'
    End
End

Context 'project/autoscaler/edit'
    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid"  -o "$oid" -I)
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

    Example 'edit when not enabled'
        When call taikun project autoscaler edit "$pid" --max-size 9 --min-size 3
        The status should equal 1
        The stderr should include 'Project autoscaling is disabled'
    End
End