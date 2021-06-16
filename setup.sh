#!/bin/bash


if ! dpkg --get-selections | grep -q "^curl[[:space:]]*install$" >/dev/null; then
    echo "cURL not installed, installing..."
    sudo apt install curl
fi

./setup-zsh.sh

if [ -d "$HOME/.pyenv/" ]; then
    pyenv update
    echo "Oh-My-Zsh installed, run omz update in terminal"
else
    echo "Install pyenv"
    curl https://pyenv.run | bash
fi

if ! grep -Fxq 'eval "$(pyenv init -)"' "$HOME/.zshrc"
then
    echo "Configuring pyenv for zsh"
    echo 'eval "$(pyenv init -)"' >> ~/.zshrc
fi
which pyenv

./setup-docker.sh
# TODO: install golang if not exists
echo "Install golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0
which golangci-lint


echo "Installed"


./setup-git.sh
