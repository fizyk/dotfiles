#!/bin/bash


if ! dpkg --get-selections | grep -q "^curl[[:space:]]*install$" >/dev/null; then
    echo "cURL not installed, installing..."
    sudo apt install curl
fi

./setup-zsh.sh
./setup-pyenv.sh
./setup-docker.sh
# TODO: install golang if not exists
echo "Install golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0
which golangci-lint


echo "Installed"


./setup-git.sh
