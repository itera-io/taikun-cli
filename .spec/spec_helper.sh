# shellcheck shell=bash

_rnd_name() {
  echo "tk-cli-test-$(shuf --echo --repeat --head-count=8 {a..z} | tr -d '\n')"
}

_rnd_between() {
  local min=$1
  local max=$2
  echo $((RANDOM % (max - min + 1) + min))
}

# Autoscaler flavor
AUTOSCALER_FLAVOR="m1.extra_tiny"

spec_helper_precheck() {
  : minimum_version "0.28.1"
}

spec_helper_loaded() {
  :
}

spec_helper_configure() {
  : import 'support/custom_matcher'
}
