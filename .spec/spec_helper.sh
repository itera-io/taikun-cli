# shellcheck shell=sh

_rnd_name() {
  echo $(shuf --echo --repeat --head-count=8 {a..z} | tr -d '\n')
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
