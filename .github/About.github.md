# .github

This directory contains Bash script files to check: lint, unit test, coverage and etc. before merge.

Mostly used for CIs, such as GitHub Actions. But if you have the required commands installed in your local or in your dev environment, then it is very convenient to check ahead.

To check if the required commands are installed, run:

- `check-requirements.sh`

## PR (Pull Request) Regulation

Any `PR` (Pull Request) made must pass all the tests, which is the `run-tests-merge.sh` script, before any code review.

If you have all the required commands installed then run:

- `run-tests-merge.sh`

## Script File Description

- `check-documentation.sh` ... Checks if the auto-generated documents are up-to-date.
- `check-requirements.sh` ... Checks if required commands are installed for testing.
- `run-tests-coverage.sh` ... Runs all the Golang unit tests and checks if its coverage was 100%.
- `run-tests-lint.sh` ... Runs linters for Shell scripts and Golang source codes.
- `run-tests-merge.sh` ... Runs all the above scripts.
- `update-docs.sh` ... Updates all the auto-generated `README.md` documents of each package in the repo.

## TIPS

In the `../.devcontainer/` directory there's a docker environment for those who doesn't have a Go 1.16 dev environment.

Suitable for Docker + VSCode + Remote Container users as well.
