Context 'project/add'

    setup() {
        oid=$(taikun organization add $(_rnd_name) -f $(_rnd_name) -I)
        ccid=$(taikun cloud-credential openstack add $(_rnd_name) -o $oid -d $OS_USER_DOMAIN_NAME -p $OS_PASSWORD --project $OS_PROJECT_NAME -r $OS_REGION_NAME -u $OS_USERNAME --public-network $OS_INTERFACE --url $OS_AUTH_URL -I)
    }

    BeforeAll 'setup'

    cleanup() {
        taikun cloud-credential delete $ccid -q 2>/dev/null || true
        taikun organization delete $oid
    }

    AfterAll 'cleanup'

    Context
        autoscaler_default_project() {
            pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid  -o $oid --autoscaler  --autoscaler-name $(_rnd_name) --autoscaler-flavor "m1.extra_tiny" -I)
        }

        cleanup() {
            if ! taikun project delete $pid -q 2>/dev/null; then
                taikun project delete --force $pid -q 2>/dev/null || true
            fi
        }

        After 'cleanup'

        Example 'basic autoscaler project'
            When call autoscaler_default_project
            The status should equal 0
        End
    End

    Context
        autoscaler_project() {
            pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid -o $oid --autoscaler  --autoscaler-name $(_rnd_name) --autoscaler-flavor "m1.extra_tiny" --autoscaler-disk-size 32 --autoscaler-min-size 2 --autoscaler-max-size 10 -I)
        }

        cleanup() {
            if ! taikun project delete $pid -q 2>/dev/null; then
                taikun project delete --force $pid -q 2>/dev/null || true
            fi
        }

        After 'cleanup'

        Example 'autoscaler project'
            When call autoscaler_project
            The status should equal 0
        End
    End

    Context
        not_autoscaler_project() {
            pid=$(taikun project add $(_rnd_name) --cloud-credential-id $ccid -o $oid  --autoscaler-name $(_rnd_name) --autoscaler-flavor "m1.extra_tiny" -I)
            taikun project list -o $oid --limit 1 --format json
            if ! taikun project delete $pid -q 2>/dev/null; then
                taikun project delete --force $pid -q 2>/dev/null || true
            fi
        }

        Example 'autoscaler project'
            When call not_autoscaler_project
            The status should equal 0
            The output should include '"isAutoscalingEnabled": false'
        End
    End
End