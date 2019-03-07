#!/bin/sh
set -e
DIR=$(dirname $0)

if [ -z "${SEMVER_BRANCH}" ] || [ -z "${SEMVER_TARGET}" ] || [ -z "${SEMVER_VERSION}" ]; then
    . ${DIR}/semver-env.sh
fi

if [ -e .semver/${SEMVER_TARGET} ]; then
    ${DIR}/semver-next.sh "$@" > .semver/${SEMVER_TARGET}
    ${DIR}/semver-commit.sh
else
    exit 1
fi
