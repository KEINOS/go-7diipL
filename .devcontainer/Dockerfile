# =============================================================================
#  Pre Defined Variables
# =============================================================================
# User provided Go version of Docker image
ARG VARIANT
# Default Go version. Choices: 1, 1.16, 1.15, 1.14
ARG VARIANT_DEFAULT="1.16"
# Default Language
ARG LC_ALL_DEFAULT='en_US.utf8'
ARG LANG_DEFAULT='en_US.utf8'
# Default Node.js settings
ARG INSTALL_NODE_DEFAULT="false"
ARG NODE_VERSION_DEFAULT="lts/*"
ARG LANG="${LANG:-$LANG_DEFAULT}"
ARG LC_ALL="${LC_ALL:-$LC_ALL_DEFAULT}"

# =============================================================================
#  Main
# =============================================================================
FROM golang:${VARIANT:-$VARIANT_DEFAULT}-buster

# Declare user args to receive while building an image.
ARG LANG
ARG LC_ALL
ARG INSTALL_NODE="${INSTALL_NODE:-INSTALL_NODE_DEFAULT}"
ARG NODE_VERSION="${NODE_VERSION:-NODE_VERSION_DEFAULT}"

ENV \
    LANG="$LANG" \
    LANGUAGE="$LANG" \
    LC_ALL="$LC_ALL" \
    PATH="/usr/local/go/bin:${PATH}" \
    # Enforce go module mode
    GO111MODULE='on' \
    # Fix: https://github.com/microsoft/vscode-dev-containers/issues/51
    SHELL="/bin/bash" \
    DEBIAN_FRONTEND=noninteractive

#RUN localedef -f UTF-8

# [Option] Install Node.js
RUN if [ "${INSTALL_NODE}" = "true" ]; then \
    echo 'Installing Node.js'; \
    su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; \
    fi

# Setup scripts
COPY postCreateCommand.sh /postCreateCommand.sh
COPY installToolsForDev.sh /installToolsForDev.sh
COPY createVSCodeUser.sh /createVSCodeUser.sh

#  Pre-setup
# -----------------------------------------------------------------------------
RUN \
    echo 'ENV INFO:'; env \
    # Get to latest versions of all current packages
    && apt-get -y update \
    && apt-get -y upgrade --no-install-recommends \
    && apt-get autoremove -y \
    # Set env path for go to .bashrc
    && echo 'export PATH="/go/bin:/usr/local/go/bin:${PATH}"' >> "${HOME}/.bashrc" \
    # Install tools to help dev
    && /bin/bash /installToolsForDev.sh \
    # Create vscode user for non-root-user usage
    && /bin/bash /createVSCodeUser.sh

RUN \
    echo "${LC_ALL} UTF-8" >>/etc/locale.gen \
    && locale-gen "$LC_ALL" \
    && update-locale LANG="$LANG"

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
