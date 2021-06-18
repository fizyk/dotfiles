GOLANG_VERSION="1.16.5"
ZSHRC_FILE="$HOME/.zshrc"
echo "Installing and configuring golang $GOLANG_VERSION"

if ! [ -f "go$GOLANG_VERSION.linux-amd64.tar.gz" ]; then
    echo "Download golang $GOLANG_VERSION"
    wget https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
fi

if [ -d /usr/local/go ]; then
    sudo rm -rf /usr/local/go
fi

sudo tar -C /usr/local -xzf "go$GOLANG_VERSION.linux-amd64.tar.gz"

if ! grep -Fxq 'export PATH=$PATH:/usr/local/go/bin' "$ZSHRC_FILE"
then
    echo '# Golang configuration'
    echo 'export PATH=$PATH:/usr/local/go/bin' >> $ZSHRC_FILE
fi

/usr/local/go/bin/go version