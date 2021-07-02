#!/bin/bash
# =============================================================================
#  Test Script to Run GolangCI-Lint
# =============================================================================
#
#  Requirements:
#    - golangci-lint: https://github.com/golangci/golangci-lint

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
PATH_FILE_CONF_GOLANGCILINT="${PATH_DIR_PARENT}/.github/golangci.yml"
SUCCESS=0
FAILURE=1

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
set -eu
set -o pipefail

cd "${PATH_DIR_PARENT}" || {
    echo >&2 "Failed to change directory: ${PATH_DIR_PARENT}"
    exit $FAILURE
}

go mod tidy

echo "${@}" | grep -e "-l" -e "--linters" >/dev/null && {
    echo "List of enabled linters:"
    golangci-lint linters --config "$PATH_FILE_CONF_GOLANGCILINT" ./...
    exit $?
}

printf "%s" 'Running golangci-lint ... '
if ! golangci-lint run --config "$PATH_FILE_CONF_GOLANGCILINT" ./...; then
    echo >&2
    echo >&2 'References for debugging lint error:'
    echo >&2 '  wsl     : https://github.com/bombsimon/wsl/blob/master/doc/rules.md'
    echo >&2 '  gofumpt : https://github.com/mvdan/gofumpt#added-rules'
    echo >&2 '  * for other lint errors see: https://golangci-lint.run/usage/linters/'
    # shellcheck disable=SC2016
    echo >&2 'To suppress lint use `//nolint:<lint type>`:'
    echo >&2 '  //nolint:wsl // Your comment of reason to suppress'
    exit $FAILURE
fi

echo 'OK'
exit $SUCCESS
