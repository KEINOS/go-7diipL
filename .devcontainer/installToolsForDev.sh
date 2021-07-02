#!/bin/bash
echo '==============================================================================='
echo ' Install Helper Commands for Development'
echo '==============================================================================='

GoVersion=$(go version | grep -o -E "([0-9]+\.){1}[0-9]+(\.[0-9]+)?" | head -n1)
GoVersionMinor=$(echo "$GoVersion" | awk -F. '{printf "%d", $2}')
GoVersionRequireMin=16

NameRepo="QiiTrans"

echo "Current Go version is: ${GoVersion} (Minor ver is: ${GoVersionMinor})"
echo "Go path: $(which go)"

if [ "$GoVersionMinor" -lt "$GoVersionRequireMin" ]; then
    echo >&2 "Minimum Go version does not satisfy. Required minimum Go version: 1.${GoVersionRequireMin}"
    exit 1
fi

set -eu

echo
echo '* Installing additional apt packages'
# Install apt-utils to avoid debconf warning then verify presence of other common developer tools and dependencies
apt-get -y install --no-install-recommends \
    apt-utils \
    \
    apt-transport-https \
    ca-certificates \
    curl \
    dialog \
    iproute2 \
    git \
    gnupg2 \
    htop \
    jq \
    less \
    libc6 \
    libgcc1 \
    libgssapi-krb5-2 \
    libicu[0-9][0-9] \
    liblttng-ust0 \
    libstdc++6 \
    locales \
    lsb-release \
    lsof \
    man-db \
    nano \
    ncdu \
    net-tools \
    openssh-client \
    procps \
    psmisc \
    ripgrep \
    rsync \
    sudo \
    tree \
    unzip \
    vim-tiny \
    wget \
    xz-utils \
    zip \
    zlib1g

# Install libssl1.1 if available
if [[ -n $(apt-cache --names-only search ^libssl1.1$) ]]; then
    apt-get -y install libssl1.1
fi

# Install appropriate version of libssl1.0.x if available
LIBSSL=$(dpkg-query -f '${db:Status-Abbrev}\t${binary:Package}\n' -W 'libssl1\.0\.?' 2>&1 || echo '')
if [ "$(echo "$LIBSSL" | grep -o 'libssl1\.0\.[0-9]:' | uniq | sort | wc -l)" -eq 0 ]; then
    if [[ -n $(apt-cache --names-only search ^libssl1.0.2$) ]]; then
        # Debian 9
        apt-get -y install libssl1.0.2
    elif [[ -n $(apt-cache --names-only search ^libssl1.0.0$) ]]; then
        # Ubuntu 18.04, 16.04, earlier
        apt-get -y install libssl1.0.0
    fi
fi

echo
echo '* Installing tools in go'
go install "github.com/msoap/go-carpet@latest"
go install "mvdan.cc/sh/v3/cmd/shfmt@latest"
go install "github.com/tenntenn/goplayground/cmd/gp@latest"
go install "github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest"
go install "github.com/nicksnyder/go-i18n/v2/goi18n@latest"
go install "mvdan.cc/gofumpt@latest"
go install "golang.org/x/tools/gopls@latest"
go install "github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest"
go install "github.com/ramya-rao-a/go-outline@latest"
go install "github.com/go-delve/delve/cmd/dlv@latest"
go install "honnef.co/go/tools/cmd/staticcheck@latest"

echo
echo '* Installing shellcheck'
# ShellCheck - Static analyzer and formatter for shell script
# Note: Install the latest shellcheck. See: https://github.com/koalaman/shellcheck/issues/704
url_download="https://github.com/koalaman/shellcheck/releases/download/latest/shellcheck-latest.linux.$(uname -m).tar.xz" &&
    timestamp="$(date +%Y%m%d%H%M%S)" &&
    path_tmp_dir=$(mktemp "/tmp/${NameRepo}-${timestamp}.tmp.XXXXXX") &&
    echo "TEMP PATH: ${path_tmp_dir}" &&
    wget -P "${path_tmp_dir}/" "$url_download" &&
    tar xvf "${path_tmp_dir}"/shellcheck* -C "${path_tmp_dir}/" &&
    cp "${path_tmp_dir}/shellcheck-latest/shellcheck" "${GOPATH:?Undefined}/bin/shellcheck" &&
    shellcheck --version &&
    rm -r "$path_tmp_dir"

echo
echo '* Installing golangci-lint'
# golangci-lint - The fast Go linters runner. Version=latest
# binary will be installed under: $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" &&
    golangci-lint --version
