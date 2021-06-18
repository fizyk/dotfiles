
ZSHRC_FILE="$HOME/.zshrc"

if [ -d "$HOME/.pyenv/" ]; then
    pyenv update
    echo "Oh-My-Zsh installed, run omz update in terminal"
else
    echo "Install pyenv"
    curl https://pyenv.run | bash
fi

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
    echo "Configuring pyenv for zsh"
    echo 'eval "$(pyenv init -)"' >> $ZSHRC_FILE
fi
which pyenv
