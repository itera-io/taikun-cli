Context 'project/k8s/commit'
    setup() {
        oid=$(taikun organization add "$(_rnd_name)" -f "$(_rnd_name)" -I | xargs )
        ccid=$(taikun cloud-credential openstack add "$(_rnd_name)" -s "$OS_PASSWORD" --project "$OS_PROJECT_NAME" -r "$OS_REGION_NAME" -i "$OS_USERNAME" --public-network "$OS_INTERFACE" --url "$OS_AUTH_URL" -o "$oid" -I)
        flavor=$(taikun cloud-credential flavors "$ccid" --no-decorate --limit 1 -C name | xargs)
        pid=$(taikun project add "$(_rnd_name)" --cloud-credential-id "$ccid" --flavors "$flavor" -I | xargs)

    }

    BeforeAll 'setup'

    cleanup() {
        taikun project k8s delete "$pid" -a -q 2>/dev/null || true
        if ! taikun project delete "$pid" -q 2>/dev/null; then
                taikun project delete --force "$pid" -q 2>/dev/null || true
            fi
        taikun cloud-credential delete "$ccid" -q 2>/dev/null || true
        taikun organization delete "$oid" -q 2>/dev/null || true
    }

    AfterAll 'cleanup'

    example_setup() {
        taikun project k8s add "$pid" -n master --flavor "$flavor" -r kubemaster -q
        taikun project k8s add "$pid" -n master2 --flavor "$flavor" -r kubemaster -q
        taikun project k8s add "$pid" -n bastion --flavor "$flavor" -r bastion -q
        taikun project k8s add "$pid" -n worker --flavor "$flavor" -r kubeworker -q
        taikun project k8s commit "$pid"
    }

    Example 'commit with two master nodes'
        When call example_setup
        The status should equal 1
        The stderr should include 'odd number of master'
    End
End