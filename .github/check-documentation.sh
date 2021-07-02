#!/bin/bash
# =============================================================================
#  This script checks if the package documentation is up-to-date.
# =============================================================================
#  To update documents (the README.md) of each package, run:
#   $ ../.github/update-docs.sh

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
SUCCESS=0
FAILURE=1
URL_REPO_DEFAULT='https://github.com/Qithub-BOT/QiiTrans'
NAME_BRANCH_DEFAULT='main'

# -----------------------------------------------------------------------------
#  Main script
# -----------------------------------------------------------------------------
cd "$PATH_DIR_PARENT" || {
    echo >&2 'Failed to change directory'

    exit $FAILURE
}

result=$(gomarkdoc -u -vv --check --footer="------" --repository.url="$URL_REPO_DEFAULT" --repository.default-branch="$NAME_BRANCH_DEFAULT" --output '{{.Dir}}/README.md' ./... 2>&1) || {
    echo "$result"
    echo >&2 'To update docs, run:'
    echo >&2 '  $ ./.github/update-docs.sh'
    echo >&2 'Note: Do not forget to check if the changes were commited as well.'

    exit $FAILURE
}

echo 'All documents are up-to-date.'

exit $SUCCESS
