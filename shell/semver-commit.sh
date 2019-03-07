#!/bin/sh
set -e
DIR=$(dirname $0)

if [ -z "${SEMVER_BRANCH}" ] || [ -z "${SEMVER_TARGET}" ] || [ -z "${SEMVER_VERSION}" ]; then
    . ${DIR}/semver-env.sh
fi

if [ -e .semver/${SEMVER_TARGET} ]; then
    cd .semver
    git add ${SEMVER_TARGET}
    git commit --message "semver(${SEMVER_TARGET}): $(cat ${SEMVER_TARGET})"
else
    exit 1
fi
