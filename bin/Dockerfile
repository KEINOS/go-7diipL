# =============================================================================
#  Application Binary Builder
# =============================================================================
# Dockerfile to build app binary for your architecture.
#
# For your convenience run the "./build-app.sh" script in this directory which
# does everything you need to build the app.
FROM golang:1.16-alpine

WORKDIR /app

ENV \
    # Enforce go module mode
    GO111MODULE='on' \
    # Fix: https://github.com/microsoft/vscode-dev-containers/issues/51
    SHELL="/bin/bash"

RUN \
    apk add --no-cache \
    alpine-sdk \
    build-base \
    bash

ENTRYPOINT [ "/app/bin/build-app.sh" ]
