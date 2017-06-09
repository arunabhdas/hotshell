#!/bin/bash
IFS=$'\n\t'
set -oeu pipefail

. ./scripts/common-vars.sh

./scripts/generate.sh

./scripts/format.sh

./scripts/vet.sh

#todo find-out if '-cover' low percentages are due to bindata.go files
go test $ALL_BUT_VENDORS -timeout 10s