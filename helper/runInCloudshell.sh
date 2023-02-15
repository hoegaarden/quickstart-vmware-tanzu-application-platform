#!/usr/bin/env bash

set -e
set -u
set -o pipefail

: "${REPO:=https://github.com/vmware-tap-on-public-cloud/quickstart-vmware-tanzu-application-platform}"
: "${BRANCH:=main}"
: "${DIR:=qs}"
readonly REPO BRANCH DIR

main() {
  echo >&2 '## >---- setup ----'

  echo >&2 '# installing golang'
  sudo yum install -y golang

  echo >&2 '# setting up the repo'
  [[ -d "$DIR" ]] || {
    git clone "${REPO}" --depth 1 --branch "${BRANCH}" "${DIR}"
  }
  cd "${DIR}"
  git pull
  echo >&2 '## <---- setup ----'

  echo >&2
  echo >&2 '## >---- running the helper thingamajg ----'
  cd helper
  go run main.go
}

main "$@"
