# dotfiles
Personal dotfiles and system tools installation scripts


## Prerequistes

1. Install taskfile:

   https://taskfile.dev/installation/

2. Install git

1. Install golang, or run:

        setup-go.sh

2. Install mage:

        go run main.go install:mage



For now majority of tools, and configurations I'm using are set up with bash scripts, but are slowly being moved into a golang scripts, to be run by [mage](https://magefile.org/)

To list existing mage targets run:

    mage