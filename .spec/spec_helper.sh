# shellcheck shell=bash

_rnd_name() {
  echo "tk-cli-test-$(shuf --echo --repeat --head-count=8 {a..z} | tr -d '\n')"
}

# Radom number is created from numbers 1-9, because trailing zero can cause problems in prints
_rnd_number() {
  shuf --echo --repeat --head-count=6 $(seq 1 9) | tr -d '\n'
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
