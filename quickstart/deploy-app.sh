#!/bin/bash

# User sets the environment variables
# -----------------------------------
export TAIKUN_EMAIL=""
export TAIKUN_PASSWORD=""
export TAIKUN_API_HOST="api.taikun.dev"
ORGANIZATION_ID=""
PROJECT_ID=""

# Program deploys a apache app
# ----------------------------
# Version 1: Install directly from managed apps without parameters- fast and simple
#managed_cat_id=`taikun catalog list -o "$oid" --no-decorate | grep " taikun-managed-apps " | cut -d ' ' -f1 | xargs`
#taikun catalog project bind "managed_cat_id" "$PROJECT_ID"
#taikun catalog application instance install "mock-apache" managed_cat_id "$PROJECT_ID"

# Version 2: Create catalog, bind application, install
echo "hello: world" > params.yaml # Prepare variables

# Create catalog, bind app and project to catalog from repository (in reality this can probably be some private imported repository), install
taikun catalog create "cli-mock-catalog" -o "$ORGANIZATION_ID" -d "cli-mock-catalog"
catalog_id=`taikun catalog list -o "$ORGANIZATION_ID" --no-decorate | grep " cli-mock-catalog " | cut -d ' ' -f1 | xargs`
taikun catalog project bind "$catalog_id" "$PROJECT_ID"
taikun catalog application bind "$catalog_id" apache taikun-managed-apps
taikun catalog application instance install "mock-apache" "$catalog_id" "$PROJECT_ID" --extra-values-file params.yaml

rm params.yaml # Cleanup params
