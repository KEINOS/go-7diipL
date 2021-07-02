#!/bin/bash
echo '==============================================================================='
echo ' Create VSCode User'
echo '==============================================================================='

set -eu

echo "ENV INFO:"
env

USERNAME="vscode"
USERGROUP="vscode"
USER_UID=1000
USER_GID=1000

if ! id -u "${USERNAME}" >/dev/null 2>&1; then
    # Create group
    groupadd -f --gid "$USER_GID" "$USERGROUP"

    # Create user
    useradd -s /bin/bash -d "/home/${USERNAME}" --uid "$USER_UID" --gid "$USER_GID" -m "$USERNAME"

    # Add add sudo support for non-root user
    echo "$USERNAME" ALL=\(root\) NOPASSWD:ALL >/etc/sudoers.d/"$USERNAME"
    chmod 0440 /etc/sudoers.d/"$USERNAME"
    export EXISTING_NON_ROOT_USER="${USERNAME}"
fi
