#!/bin/bash


if ! dpkg --get-selections | grep -q "^curl[[:space:]]*install$" >/dev/null; then
    echo "cURL not installed, installing..."
    sudo apt install curl
fi

./setup-zsh.sh

echo "Install pyenv"
curl https://pyenv.run | bash
which pyenv
pyenv update

echo "Install lazydocker"
# https://github.com/jesseduffield/lazydocker
curl https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install_update_linux.sh | bash
which lazydocker

echo "Install golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0
which golangci-lint

COMPOSE_VERSION="1.27.4"
echo "Installing docker-compose $COMPOSE_VERSION"

## Install docker compose
sudo curl -L "https://github.com/docker/compose/releases/download/$COMPOSE_VERSION/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
which docker-composet

echo "Installed"


./setup-git.sh
