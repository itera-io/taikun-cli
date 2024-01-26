Context 'project/add'

    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I)
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -o "$oid" -d "$OS_USER_DOMAIN_NAME" -p "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -u "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -I)
    }

    BeforeAll 'setup'

    cleanup() {
        taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
        taikun organization delete "$oid" -q 2>/dev/null || true
    }

    AfterAll 'cleanup'

    Context 'with autoscaler default'
        autoscaler_default_project() {
            pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid"  -o "$oid" --autoscaler  --autoscaler-name "auto" --autoscaler-flavor "$AUTOSCALER_FLAVOR" -I | xargs )
        }
        BeforeAll 'autoscaler_default_project'

        cleanup_project() {
            if ! taikun project delete "$pid" -q 2>/dev/null; then
                taikun project delete --force "$pid" -q 2>/dev/null || true
            fi
        }
        AfterAll 'cleanup_project'

        Example 'test project info command'
            When call taikun project info "$pid"
            The status should equal 0
            The lines of output should equal 33
            The output should include "K8S-PROFILE"
            The output should include "NAME"
            The output should include "$oid"
            The output should include "Unlocked"
        End

    End

    Context 'with autoscaler'
        autoscaler_project() {
            pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -o "$oid" --autoscaler  --autoscaler-name "auto" --autoscaler-flavor "$AUTOSCALER_FLAVOR" --autoscaler-disk-size 32 --autoscaler-min-size 2 --autoscaler-max-size 10 -I)
        }

        cleanup() {
            if ! taikun project delete "$pid" -q 2>/dev/null; then
                taikun project delete --force "$pid" -q 2>/dev/null || true
            fi
        }

        After 'cleanup'

        Example 'autoscaler project'
            When call autoscaler_project
            The status should equal 0
        End
    End

    Context 'without autoscaler'
        not_autoscaler_project() {
            pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" -o "$oid"  --autoscaler-name "auto" --autoscaler-flavor "$AUTOSCALER_FLAVOR" -I)
            taikun project list -o "$oid" --limit 1 --format json
            taikun project list -o "$oid" --limit 1 --format json
            if ! taikun project delete "$pid" -q 2>/dev/null; then
                taikun project delete --force "$pid" -q 2>/dev/null || true
            fi
        }

        Example 'autoscaler project'
            When call not_autoscaler_project
            The status should equal 0
            The output should include '"isAutoscalingEnabled": false'
        End
    End
End