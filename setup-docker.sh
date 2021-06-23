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
