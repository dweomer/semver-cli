#!/bin/sh
set -e
DIR=$(dirname $0)

if [ -z "${SEMVER_BRANCH}" ] || [ -z "${SEMVER_TARGET}" ] || [ -z "${SEMVER_VERSION}" ]; then
    . ${DIR}/semver-env.sh
fi

if [ -e .semver/${SEMVER_TARGET} ]; then
    semver -from ${SEMVER_VERSION} -bump "$@"
else
    exit 1
fi
