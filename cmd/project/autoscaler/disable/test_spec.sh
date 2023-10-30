Context 'project/autoscaler/disable'
    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid"  -o "$oid" --autoscaler  --autoscaler-name "$(_rnd_name)" --autoscaler-flavor "$AUTOSCALER_FLAVOR" -I)
        taikun project autoscaler disable "$pid"
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

    Example 'disable autoscaling succesfully'
        When call taikun project list -o "$oid" --limit 1 --format json
        The status should equal 0
        The output should include '"isAutoscalingEnabled": false'
    End

    Example 'disable two times'
        When call taikun project autoscaler disable "$pid"
        The status should equal 1
        The stderr should include 'Project autoscaling already disabled'
    End
End