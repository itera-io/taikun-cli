#!/bin/bash

# User sets the environment variables
# -----------------------------------
export TAIKUN_EMAIL=""
export TAIKUN_PASSWORD=""
export TAIKUN_API_HOST="api.taikun.dev"
ORGANIZATION_ID=""
CLOUD_CREDENTAIL_ID=""
PROJECT_NAME="tk-cli-quickstart"

# Program creates a BMW cluster
# -----------------------------
project_real_name="$PROJECT_NAME-$(shuf --echo --repeat --head-count=8 {a..z} | tr -d '\n')"
flavor=$(taikun cloud-credential flavors "$CLOUD_CREDENTAIL_ID" --no-decorate --min-cpu 2 --max-cpu 4 --min-ram 4 --max-ram 8 -C name --limit 1 | xargs)
projectid=$(taikun project add "$project_real_name" --cloud-credential-id "$CLOUD_CREDENTAIL_ID" --flavors "$flavor" -o "$ORGANIZATION_ID" -I)

servername1="tk-cli-w"
serverid1=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername1" --role Kubeworker -I )
servername2="tk-cli-c"
serverid2=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername2" --role Kubemaster -I )
servername3="tk-cli-b"
serverid3=$(taikun project k8s add "$projectid" --flavor "$flavor" --name "$servername3" --role Bastion -I )

taikun project k8s commit "$projectid"