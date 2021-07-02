#!/bin/bash
# -----------------------------------------------------------------------------
#  Aliases for Useful Shell Script
# -----------------------------------------------------------------------------

alias build-app='cd "$PATH_DIR_ROOT_REPO"; ./bin/build-app.sh'
alias run-all='cd "$PATH_DIR_ROOT_REPO"; ./.github/run-tests-merge.sh'
alias run-coverage='cd "$PATH_DIR_ROOT_REPO"; ./.github/run-tests-coverage.sh'
alias run-lint='cd "$PATH_DIR_ROOT_REPO"; ./.github/run-tests-lint.sh'
alias update-docs='cd "$PATH_DIR_ROOT_REPO"; ./.github/update-docs.sh'
alias update-readme='cd "$PATH_DIR_ROOT_REPO"; ./.github/update-docs.sh'
alias welcome='cd "$PATH_DIR_ROOT_REPO"; ${HOME}/.welcome.sh'
alias qiitrans='go run "$PATH_DIR_ROOT_REPO"/src/main '
