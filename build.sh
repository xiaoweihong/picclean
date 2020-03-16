#!/bin/bash

if [[ $# -eq 0 ]]; then
  echo "relase num is null"
  # shellcheck disable=SC2242
  exit 500
fi

# shellcheck disable=SC2034
tag=$1

RELEASENAME=picclean-$1

mkdir release/"${RELEASENAME}"

GOOS=linux ARCH=amd64 go build -v -o picclean main.go
cp picclean config.yaml readme.md release_note.md release/"${RELEASENAME}"
rm -f picclean
cd release
tar zcvf "${RELEASENAME}".tar.gz "${RELEASENAME}"