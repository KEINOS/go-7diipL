# =============================================================================
#  Multi-Stage Builder
# =============================================================================
# Default Language
ARG LC_ALL_DEFAULT='ja_JP.utf8'
ARG LANG_DEFAULT='ja_JP.utf8'
ARG LANG="${LANG:-$LANG_DEFAULT}"
ARG LC_ALL="${LC_ALL:-$LC_ALL_DEFAULT}"

# -----------------------------------------------------------------------------
#  First Stage
# -----------------------------------------------------------------------------
FROM golang:alpine AS build-env

ARG NAME_FILE_ARCH='qiitask-linux-'

COPY . /app

WORKDIR /app

RUN \
    # Install Deps
    apk add --no-cache \
    alpine-sdk \
    build-base \
    bash \
    zip \
    tzdata \
    musl-locales \
    && go mod tidy \
    # Build app binary
    && ./bin/build-app.sh \
    # Extract the built app binary
    && path_file_archive="$(ls /app/bin/qiitask-linux*.zip)" \
    && unzip -q "$path_file_archive" -d "/app/bin" \
    && rm -rf "$path_file_archive" \
    && path_dir_archive="$(cd /app/bin/qiitask-linux*; pwd)" \
    && cd "$path_dir_archive" \
    # Install app
    && mv qiitask /go/bin \
    # TimeZone
    && echo "Asia/Tokyo" >  /etc/timezone

# -----------------------------------------------------------------------------
#  Final Stage
# -----------------------------------------------------------------------------
FROM alpine

ARG LC_ALL_DEFAULT='ja_JP.utf8'
ARG LANG_DEFAULT='ja_JP.utf8'
ARG LANG="${LANG:-$LANG_DEFAULT}"
ARG LC_ALL="${LC_ALL:-$LC_ALL_DEFAULT}"

ENV \
    LANG="$LANG" \
    LANGUAGE="$LANG" \
    LC_ALL="$LC_ALL"

RUN \
    apk add --no-cache \
    musl-locales

COPY --from=build-env /go/bin/qiitask /usr/local/bin/qiitask
COPY --from=build-env /usr/share/zoneinfo/Japan /etc/localtime
COPY --from=build-env /etc/timezone /etc/timezone

WORKDIR /tasks

ENTRYPOINT ["/usr/local/bin/qiitask"]
