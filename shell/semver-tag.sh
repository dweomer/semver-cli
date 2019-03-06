#!/bin/sh
set -e
DIR=$(dirname $0)
. ${DIR}/semver-env.sh

if [ -e .semver/${SEMVER_TARGET} ]; then
    git tag v$(semver -from ${SEMVER_VERSION}) "$@"
else
    exit 1
fi
