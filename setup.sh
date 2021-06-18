#!/bin/bash


if ! dpkg --get-selections | grep -q "^curl[[:space:]]*install$" >/dev/null; then
    echo "cURL not installed, installing..."
    sudo apt install curl
fi

./setup-zsh.sh
./setup-pyenv.sh
./setup-docker.sh

./setup-go.sh


echo "Installed"


./setup-git.sh
