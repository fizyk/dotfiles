# Add docker gpg key
DOCKER_APT_PGP_KEYFILE="/usr/share/keyrings/docker-archive-keyring.gpg"
DOCKER_APT_REPO_FILE="/etc/apt/sources.list.d/docker.list"
if ! [ -f "$DOCKER_APT_PGP_KEYFILE" ]; then
    echo "Add docker repository GPG key."
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o $DOCKER_APT_PGP_KEYFILE
fi
# create docker sources list

if ! [ -f "$DOCKER_APT_REPO_FILE" ]; then
    echo "Add doker repository"
    echo "deb [arch=amd64 signed-by=$DOCKER_APT_PGP_KEYFILE] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) stable" | sudo tee $DOCKER_APT_REPO_FILE > /dev/null
    sudo apt update -o Dir::Etc::sourcelist="$DOCKER_APT_REPO_FILE" -o Dir::Etc::sourceparts="-" -o APT::Get::List-Cleanup="0"
fi

if ! dpkg --get-selections | grep -q "^docker-ce[[:space:]]*install$" >/dev/null; then
    echo "Docker not installed, installing..."
    sudo apt install -y docker-ce docker-ce-cli containerd.io
fi

usermod -a -G docker $USER
newgroup docker
# Assume primary group is same as user's username
newgroup $USER

echo "Install lazydocker"
# https://github.com/jesseduffield/lazydocker
curl https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install_update_linux.sh | bash
which lazydocker

COMPOSE_VERSION="1.27.4"
echo "Installing docker-compose $COMPOSE_VERSION"

## Install docker compose
sudo curl -L "https://github.com/docker/compose/releases/download/$COMPOSE_VERSION/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
which docker-compose
