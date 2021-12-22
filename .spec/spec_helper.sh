# shellcheck shell=sh

_rnd_name() {
  echo "tf-acc-test-$(tr -dc 'a-z' </dev/urandom | head -c 8)"
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
