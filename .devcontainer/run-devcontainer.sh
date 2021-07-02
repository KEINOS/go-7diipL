#!/bin/bash
# =============================================================================
#  Bash Script to boot dev container for Non Remote Container + VSCode users.
# =============================================================================
#  This script

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_SCRIPT="$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)"
PATH_DIR_PARENT="$(dirname "$PATH_DIR_SCRIPT")"
FAILURE=1
NAME_IMG_DOCKER="qiitrans_img"
PATH_DIR_WORK="/workspace/QiiTrans"

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
if [ -f /.dockerenv ]; then
    echo 'You are already inside a Docker container.'
    /bin/bash "${PATH_DIR_SCRIPT}/postCreateCommand.sh"
    echo
    echo
    echo
    /bin/bash
    exit $?
fi

cd "$PATH_DIR_SCRIPT" || {
    echo >&2 "Failed to change directory. Path: ${PATH_DIR_SCRIPT}"
    exit $FAILURE
}
echo "- Current path: $(pwd)"

echo '- Pull base image'
docker pull golang:1.16-buster

echo '- Build container image'
docker build --tag "$NAME_IMG_DOCKER" .

cd "$PATH_DIR_PARENT" || {
    echo >&2 "Failed to change directory to the parent. Path: ${PATH_DIR_PARENT}"
    exit $FAILURE
}
echo "- Current path: $(pwd)"

echo '- Running container'
docker run \
    --rm \
    --interactive \
    --tty \
    --env LC_ALL \
    --env LANG \
    --volume "$(pwd):$PATH_DIR_WORK" \
    --entrypoint "${PATH_DIR_WORK}/.devcontainer/run-devcontainer.sh" \
    --workdir "$PATH_DIR_WORK" \
    "$NAME_IMG_DOCKER"
