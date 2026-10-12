#!/usr/bin/env bash
# `deps:apt` decides whether a package is installed by grepping
# `dpkg --get-selections`. Feed it canned output through a dpkg stub and see
# whether a dry run would call apt.

# shellcheck source=lib.sh
source "$(dirname "$0")/lib.sh"

stub sudo 'exit 0'

# installs DEP SELECTIONS: "apt" if deps:apt would install DEP, "skip" if not.
installs() {
    stub dpkg "printf '%b' '$2'"
    if task --taskfile "$ROOT/tests/Taskfile.yaml" --dry apt DEP="$1" 2>&1 |
        grep -q 'apt install'; then
        echo apt
    else
        echo skip
    fi
}

check "installed package is skipped" \
    skip "$(installs zsh 'curl\t\t\t\t\tinstall\nzsh\t\t\t\t\t\tinstall\n')"
check "installed :amd64 package is skipped" \
    skip "$(installs containerd.io 'containerd.io:amd64\t\t\tinstall\n')"
check "missing package is installed" \
    apt "$(installs zsh 'curl\t\t\t\t\tinstall\n')"
check "package marked deinstall is installed" \
    apt "$(installs zsh 'zsh\t\t\t\t\t\tdeinstall\n')"
check "package that only shares a prefix is installed" \
    apt "$(installs git 'git-man\t\t\t\t\tinstall\n')"

finish
