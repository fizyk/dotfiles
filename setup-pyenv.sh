
ZSHRC_FILE="$HOME/.zshrc"

if [ -d "$HOME/.pyenv/" ]; then
    echo "Updating pyenv"
    $HOME/.pyenv/bin/pyenv update
else
    echo "Install pyenv"
    curl https://pyenv.run | bash
fi

echo "Checking and configuring pyenv for zsh"
if ! grep -Fxq 'export PYENV_ROOT="$HOME/.pyenv"' "$ZSHRC_FILE"
then
    echo 'export PYENV_ROOT="$HOME/.pyenv"' >> $ZSHRC_FILE
fi

if ! grep -Fxq 'export PATH="$PYENV_ROOT/bin:$PATH"' "$ZSHRC_FILE"
then
    echo 'export PATH="$PYENV_ROOT/bin:$PATH"' >> $ZSHRC_FILE
fi

if ! grep -Fxq 'eval "$(pyenv init --path)"' "$ZSHRC_FILE"
then
    echo 'eval "$(pyenv init --path)"' >> $ZSHRC_FILE
fi

if ! grep -Fxq 'eval "$(pyenv init -)"' "$ZSHRC_FILE"
then
    echo 'eval "$(pyenv init -)"' >> $ZSHRC_FILE
fi

$HOME/.pyenv/bin/pyenv --version
