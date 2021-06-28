

if ! dpkg --get-selections | grep -q "^zsh[[:space:]]*install$" >/dev/null; then
    echo "Zshell not installed, installing..."
    sudo apt install zsh
fi

if [ -d "$HOME/.oh-my-zsh/" ]; then
    echo "Oh-My-Zsh installed, run omz update in terminal"
else
    echo "Install Oh-My-Zsh"
    sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
    echo "Install powerlevel10k"
    git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k
    # Set ZSH_THEME="powerlevel10k/powerlevel10k"
    echo "Install zsh-syntax-highlighting"
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
    # https://github.com/zsh-users/zsh-syntax-highlighting/blob/master/INSTALL.md#oh-my-zsh
    echo "Install zsh-autosuggestions"
    git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
    # https://github.com/zsh-users/zsh-autosuggestions/blob/master/INSTALL.md#oh-my-zsh
    echo "Go to https://github.com/ryanoasis/nerd-fonts/releases and select proper fonts, install it in the system and configure your terminal to use it."
    # https://github.com/ryanoasis/nerd-fonts/releases/download/v2.1.0/UbuntuMono.zip
fi

# chsh -s /bin/zsh
