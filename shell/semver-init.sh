#!/bin/sh
set -e
DIR=$(dirname $0)
. ${DIR}/semver-env.sh

if ! (git rev-parse --abbrev-ref ${SEMVER_BRANCH} 2>/dev/null 1>/dev/null); then
    git checkout --force --orphan ${SEMVER_BRANCH}
    git reset --hard
    git clean -fdx --quiet
    echo "${SEMVER_VERSION}" > ${SEMVER_TARGET}
    git add ${SEMVER_TARGET}
    git commit --message "semver(${SEMVER_TARGET}): $(cat ${SEMVER_TARGET})"
    git checkout --force ${SEMVER_TARGET}
    git clean -fdx --quiet
    git reset --hard
fi

git worktree add .semver semver