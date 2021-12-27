#!/bin/bash
# =============================================================================
#  Bash shell script to build/compile the CLI app
# =============================================================================
#  Usage:
#    $ ./build-app.sh --help
#
#  This script will use Docker if installed. If Docker is not installed then it
#  will use the local Golang to build the application but may require some
#  dependencies.

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
# Name of the CLI app
NAME_FILE_BIN='qiitrans'

# Path info to export the app
PATH_DIR_RETURN="$(pwd)"
PATH_DIR_SCRIPT="$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)"
PATH_DIR_ROOT_REPO="$(dirname "$PATH_DIR_SCRIPT")"
PATH_DIR_BIN="${PATH_DIR_ROOT_REPO}/bin"

# Path to the main package
PATH_DIR_PKG_MAIN="${PATH_DIR_ROOT_REPO}/src/main"

# App version from git tag and rev-parse
TAG_APP=$(git describe --tag)
REV_APP=$(git rev-parse --short HEAD)

# Default values
GOOS_DEFAULT="${GOOS:-linux}"
GOARCH_DEFAULT="${GOARCH:-$(uname -m)}"
GOARM_DEFAULT="${GOARM:-}"

# Status
SUCCESS=0
FAILURE=1
TRUE=0
FALSE=1

# List of architectures when '--all' '-a'
LIST_ARCH=(
    "darwin/amd64"
    "darwin/arm64"
    "linux/386"
    "linux/amd64"
    "linux/arm/5"
    "linux/arm/6"
    "linux/arm/7"
    "linux/arm64"
    "windows/386"
    "windows/amd64"
    "windows/arm"
)

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------

archiveDirToZip() {
    path_file_target="${1:?"missing input path"}"
    path_dir_target="$(dirname "$path_file_target")"
    name_file_archive="$(basename "$path_dir_target")"

    cd "$PATH_DIR_BIN" || {
        echo >&2 "Failed to change direcotry. Path: ${PATH_DIR_BIN}"
    }

    # Archive with better compression and exclude OS dependent files.
    # Change dir to avoid including extra dir path in the archive as well.
    zip -9 "$name_file_archive" -r "$name_file_archive" -x "*.DS_Store" "*__MACOSX*" 1>/dev/null || {
        echo >&2 'Failed to archive directory'

        return $FAILURE
    }
    echo '  Archive ... OK'

    # Delete dir if success
    rm -rf "${path_dir_target:?"dir path not set to delete"}" || {
        echo >&2 'Failed to delete temp directory'

        return $FAILURE
    }
    echo '  Clean up ... OK'

    return $SUCCESS
}

buildBinary() {
    GOOS="${1:-"$GOOS_DEFAULT"}"
    GOARCH="${2:-"$GOARCH_DEFAULT"}"
    GOARM="${3:-"$GOARM_DEFAULT"}"
    GOARM_SUFFIX="${GOARM:+"-${GOARM}"}" # apply suffix only when GOARM is set

    if [ "$GOARCH" = "x86_64" ]; then
        GOARCH="amd64"
    fi

    VER_APP=""
    if [ "${TAG_APP:+undefined}" ]; then
        VER_APP="${TAG_APP} (${REV_APP})"
    fi

    NAME_DIR_BIN="${NAME_FILE_BIN}-${GOOS}-${GOARCH}${GOARM_SUFFIX}"
    PATH_DIR_OUT="${PATH_DIR_BIN}/${NAME_DIR_BIN}"
    PATH_FILE_BIN="${PATH_DIR_OUT}/${NAME_FILE_BIN}"

    if [ -e "${PATH_DIR_OUT}" ]; then
        rm -rf "${PATH_DIR_OUT:?"Output dir not set"}"
    fi

    # Make output dir
    mkdir -p "$PATH_DIR_OUT"

    # Build as static linked binary
    echo "- Building static linked binary to ... ${PATH_FILE_BIN}"
    echo "  Arch: ${GOOS} ${GOARCH}${GOARM}"

    if CGO_ENABLED=0 \
        GOOS="$GOOS" \
        GOARCH="$GOARCH" \
        GOARM="$GOARM" \
        go build \
        -installsuffix "$NAME_FILE_BIN" \
        -ldflags="-s -w -extldflags \"-static\" -X 'main.version=${VER_APP}'" \
        -o="$PATH_FILE_BIN" \
        "$PATH_DIR_PKG_MAIN"; then

        echo '  Build ... OK'

        archiveDirToZip "$PATH_FILE_BIN" || {
            echo >&2 'Failed to archive the built binary.'

            return $FAILURE
        }

        return $SUCCESS
    fi
    echo >&2 'Failed to build binary.'

    echo
    echoHelp

    return $FAILURE
}

echoHelp() {
    cat <<'HEREDOC'
About:
  This script builds the application binary and ZIP archives it under ./bin directory.

Usage:
  ./build-app.sh <GOOS>
  ./build-app.sh <GOOS> [<GOARCH>]
  ./build-app.sh <GOOS> [<GOARCH> <GOARM>]
  ./build-app.sh [-l] [--list] [-h] [--help]
  ./build-app.sh [-a] [--all]

GOOS:
  The fisrt argument is the OS platform. Such as:

    "linux", "darwin", "windows", etc.

  To see the supported platforms use '--list' option.

GOARCH:
  The 2nd argument is the architecture (CPU type). Such as:

    "amd64"(64bit, Intel/AMD/x86_64), "arm", "arm64", "386", etc. (Default: "amd64")

  To see the supported architectures use '--list' option.

GOARM:
  The 3rd argument is the ARM variant/version. Such as:

    "5", "6", "7". (Default: empty)

  To see the supported combination see: https://github.com/golang/go/wiki/GoArm

Options:
  -l --list ... Displays available platforms and architectures to build.
  -h --help ... Displays this help.
  -a --all .... Build major architectures at once.

Sample Usage:

  # Display available arcitectures
   $ ./build-app.sh --list
   $ ./build-app.sh -l

  # Build Linux (Intel) binary
   $ ./build-app.sh linux

  # Build macOS binary
   $ ./build-app.sh darwin        # Equivalent to: ./build-app.sh darwin amd64
   $ ./build-app.sh darwin arm64

  # Build Windows10 binary (32bit and 64bit)
   $ ./build-app.sh windows 386
   $ ./build-app.sh windows amd64

  # Build Raspberry Pi 3 binary (arm v7)
   $ ./build-app.sh linux arm 7

  # Build Raspberry Pi Zero binary (arm v6)
   $ ./build-app.sh linux arm 6

  # Build QNAP ARM5 binary (arm v5)
   $ ./build-app.sh linux arm 5

  # Build the above at once
   $ ./build-app.sh --all

HEREDOC

    exit $SUCCESS
}

isInsideDocker() {
    if [ -f /.dockerenv ]; then
        return $TRUE
    fi

    return $FALSE
}

isInstalled() {
    which "$1" 2>/dev/null 1>/dev/null || {
        return $FALSE
    }

    return $TRUE
}

indentSTDIN() {
    indent='    '
    while IFS= read -r line; do
        echo "${indent}${line}"
    done
    echo
}

listPlatforms() {
    list=$(go tool dist list) || {
        echo >&2 'Failed to get supported platforms.'
        echo "$list" | indentSTDIN 1>&2
        exit $FAILURE
    }
    echo 'List of available platforms to build. (GOOS/GOARCH)'
    echo "$list" | indentSTDIN
    exit $SUCCESS
}

# -----------------------------------------------------------------------------
#  Run Only If Not Inside Docker Container
# -----------------------------------------------------------------------------
cd "$PATH_DIR_ROOT_REPO" || {
    echo >&2 "Failed to change directory. Path: ${PATH_DIR_ROOT_REPO}"
    exit $FAILURE
}

if (! isInsideDocker) && (isInstalled docker); then
    echo -n '- Pulling Docker base image (golang:1.16-alpine) ... '
    result=$(docker pull golang:1.16-alpine 2>&1) || {
        echo 'NG'
        echo >&2 "$result"
        exit $FAILURE
    }
    echo 'OK'

    echo -n '- Building Docker image ... '
    result=$(docker build --file ./bin/Dockerfile --tag transqii_builder . 2>&1 3>&1) || {
        echo 'NG'
        echo >&2 "$result"
        exit $FAILURE
    }
    echo 'OK (Image tag: transqii_builder)'

    echo -n '- Building app binary in the container ... '
    result=$(docker run --rm -it -v "$(pwd)":/app transqii_builder "${@}" 2>&1) || {
        echo 'NG'
        echo >&2 "$result"
        exit $FAILURE
    }
    echo 'OK'

    echo "$result"
    exit $SUCCESS
fi

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
# Detect Golang
if ! isInstalled go; then
    echo >&2 'Golang not installed. Application build aborted.'
    exit $FAILURE
fi

case "$1" in
    "--help") echoHelp ;;
    "-h") echoHelp ;;
    "--list") listPlatforms ;;
    "-l") listPlatforms ;;
esac

# --all -a
if (echo "$@" | grep '\-all' 1>/dev/null 2>/dev/null) || (echo "$@" | grep '\-a' 1>/dev/null 2>/dev/null); then
    for list in "${LIST_ARCH[@]}"; do
        # shellcheck disable=SC2206
        arrList=(${list//// }) # Split delim="/"

        buildBinary "${arrList[0]}" "${arrList[1]}" "${arrList[2]}" || {
            exit 1
        }
    done

    exit
fi

# à la carte (individual designation)
buildBinary "${1:-"$GOOS_DEFAULT"}" "${2:-"$GOARCH_DEFAULT"}" "${3:-"$GOARM_DEFAULT"}"

# Change back to original
cd "$PATH_DIR_RETURN" || {
    echo >&2 "Failed to change directory. Path: ${PATH_DIR_RETURN}"
    exit $FAILURE
}
