#!/bin/sh
set -e

echo -n "# " 1>&2
type semver 1>&2

: ${SEMVER_BRANCH:=semver}
: ${SEMVER_TARGET:=$(git rev-parse --abbrev-ref HEAD)}
: ${SEMVER_VERSION:=$(cat .semver/${SEMVER_TARGET} 2>/dev/null || echo 0.0.0)}

export SEMVER_BRANCH SEMVER_TARGET SEMVER_VERSION