# shellcheck shell=sh

RSC_PREFIX="tf-acc-test-"

_rnd_name() {
  cat /dev/urandom | tr -dc 'a-z' | head -c 8
}

spec_helper_precheck() {
  : minimum_version "0.28.1"
}

spec_helper_loaded() {
  :
}

spec_helper_configure() {
  : import 'support/custom_matcher'
}
