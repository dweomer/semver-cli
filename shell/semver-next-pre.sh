#!/bin/sh
set -e
DIR=$(dirname $0)
if [ -z "${SEMVER_BRANCH}" ] || [ -z "${SEMVER_TARGET}" ] || [ -z "${SEMVER_VERSION}" ]; then
    . ${DIR}/semver-env.sh
fi

if [ -e .semver/${SEMVER_TARGET} ]; then
    SEMVER_BUMP=$1
    shift
    SEMVER_VERSION=$(${DIR}/semver-next.sh ${SEMVER_BUMP}) ./shell/semver-next.sh pre pre=${SEMVER_PRE:=pre} "$@"
else
    exit 1
fi
