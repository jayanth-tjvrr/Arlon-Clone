#!/bin/bash
# vim:filetype=sh:tabstop=4:shiftwidth=4:expandtab:

set -eo pipefail
[ ! -z "${BASH_DEBUG}" ] && set -x
PS4='${BASH_SOURCE}.${LINENO}+ '

echo "This build script serves as a place to run anything before make is invoked"

export REPO_ROOT=$(cd $(dirname $0); git rev-parse --show-toplevel)
export WORKSPACE=$(cd ${REPO_ROOT}/..; pwd)
export REPO_NAME=$(basename ${REPO_ROOT})

pushd ${REPO_ROOT}
    # lame check to see if we are running via TeamCity
    # intent is not to avoid pushing to ecr, but to avoid unnecessary pushes

    if [ -z ${TEAMCITY_BUILD_ID} ]; then
        make --max-load=$(nproc)
    else
        make --max-load=$(nproc) push-images
    fi
popd
