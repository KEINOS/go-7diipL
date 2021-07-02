#!/bin/bash
# Place any script here which you want to run after creating the container.
echo '==============================================================================='
echo ' Post Create Command'
echo '==============================================================================='

echo "- WHOAMI: $(whoami)"
echo "- PWD: $(pwd)" # This path should be the root directory of the repo

HOME="${HOME:-"/root"}"
LANG="${LANG:-"en_US.UTF-8"}"
LC_ALL="${LC_ALL:-"$LANG"}"

set -eu

# -----------------------------------------------------------------------------
#  Simbolic-link the helper files
# -----------------------------------------------------------------------------

# Sim-link Welcome message for bash
ln -s "$(pwd)/.devcontainer/welcome.sh" "${HOME}/.welcome.sh"
# Sim-link alias definition file
ln -s "$(pwd)/.devcontainer/aliases.sh" "${HOME}/.aliases.sh"
# Sim-link Cobra configuration file
ln -s "$(pwd)/.devcontainer/cobra.yaml" "${HOME}/.cobra.yaml"

# -----------------------------------------------------------------------------
#  Append lines to ~/.bashrc
# -----------------------------------------------------------------------------
{
    # Set welcome message for bash to display useful commands to help develop on every
    # new terminal open.
    #
    # Single quotes are intensional to not expand the expressions.
    # shellcheck disable=SC2016
    echo '"${HOME}/.welcome.sh"'

    # Set aliases of the helper scripts
    #
    # Single quotes are intensional to not expand the expressions.
    # shellcheck disable=SC2016
    echo 'source "${HOME}/.aliases.sh"'

    # Set directory path of the repository
    echo 'export PATH_DIR_ROOT_REPO='"$(pwd)"

    # Set language
    #
    # Single quotes are intensional to not expand the expressions.
    # shellcheck disable=SC2016
    echo 'export LANGUAGE="$LANG"'

} >>"${HOME}/.bashrc"

echo "${LC_ALL} UTF-8" | sudo tee -a /etc/locale.gen >/dev/null

sudo /usr/sbin/locale-gen "$LC_ALL"
sudo /usr/sbin/update-locale LANG="$LANG"

# Make sure go.mod matches the source code in the module.
sudo /usr/local/go/bin/go mod tidy
