#!/bin/bash
# =============================================================================
#  Document creator.
# =============================================================================
#  This script will generate/update the README.md of each package in the repo.
#  Powered by gomarkdoc (https://github.com/princjef/gomarkdoc)

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
URL_REPO_DEFAULT='https://github.com/Qithub-BOT/QiiTrans'
NAME_BRANCH_DEFAULT='main'

# -----------------------------------------------------------------------------
#  Check Commit Status
# -----------------------------------------------------------------------------
git status | grep nothing\ to\ commit 1>/dev/null 2>/dev/null || {
    echo >&2 'Chenges detected.'
    echo >&2 'Before updating documents you need to commit the changes.'
    exit 1
}

# -----------------------------------------------------------------------------
#  Update Documents
# -----------------------------------------------------------------------------
cd "$PATH_DIR_PARENT" &&
    # Update all package's README document
    gomarkdoc -u --footer="------" --repository.url="$URL_REPO_DEFAULT" --repository.default-branch="$NAME_BRANCH_DEFAULT" --output '{{.Dir}}/README.md' ./...
