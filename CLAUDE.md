# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository purpose

Personal dotfiles repo that bootstraps a developer workstation. There is no application code here — the repo is a set of [Taskfile](https://taskfile.dev) definitions that drive `apt`, `curl`, and `git clone` to install and configure a fixed list of tools. It targets Ubuntu/Debian (uses `apt`, `dpkg`, `lsb_release`, `add-apt-repository`).

## Common commands

All commands run from the repo root via the `task` binary (must be installed first; see README).

- `task --list-all` — discover every task across all included Taskfiles.
- `task install` — run every tool's `install` (git, zsh, docker, asdf, atuin) in order. This is also the default target.
- `task config` — run every tool's `config` (git, zsh, docker, asdf). Atuin has no `config` task.
- `task <ns>:install` / `task <ns>:config` — operate on one tool only, where `<ns>` is `git`, `zsh`, `docker`, `asdf`, or `atuin`.
- `task asdf:plugins` — install/pin the asdf-managed runtimes (golang, python, pre-commit, lazydocker) at the versions declared in `TaskfileAsdf.yaml` vars.
- `task atuin:update` — update an already-installed atuin.

Tasks are idempotent: each defines `status:` / `preconditions:` checks (e.g. `dpkg --get-selections | grep …`, `test -d …`, `asdf current | grep …`) so re-running skips work that's already done. When editing or adding tasks, preserve this property — add a `status:` check that detects the post-condition rather than relying on the command itself being safe to re-run.

## Architecture

The top-level `Taskfile.yaml` is a thin aggregator. It `includes:` one Taskfile per tool and exposes only meta-targets (`install`, `config`, `default`). All real work lives in the per-tool files:

- `TaskfileGit.yaml` — adds the `git-core/ppa`, installs git, copies `.gitattributes` to `~/`, sets global aliases (`cleanup`, `cleanup-remote`) and config (default branch `main`, bitbucket SSH rewrite for Go).
- `TaskfileZsh.yaml` — installs zsh + oh-my-zsh (unattended curl installer), then in `config` clones powerlevel10k, zsh-syntax-highlighting, and zsh-autosuggestions into `${ZSH_CUSTOM:-~/.oh-my-zsh/custom}` and switches the login shell. Nerd Font installation is left as a manual step (printed to stdout).
- `TaskfileDocker.yaml` — adds Docker's official apt repository (keyring under `/etc/apt/keyrings`), installs `docker-ce docker-ce-cli containerd.io docker-compose-plugin`, then in `config` creates the `docker` group and adds `$USER` to it.
- `TaskfileAsdf.yaml` — downloads the asdf binary tarball at a pinned `ASDF_VERSION`, appends the shims `PATH` line to `~/.zshrc`, and via the `plugins` task installs pinned versions of golang, python, pre-commit, and lazydocker. Python pulls in a long list of build-deps because asdf compiles CPython from source. Lazydocker uses a third-party plugin repo (`comdotlinux/asdf-lazydocker`) — noted as unmaintained but working.
- `TaskfileAtuin.yaml` — runs the upstream `setup.atuin.sh` installer; requires `~/.zshrc` to exist.
- `TaskfileDeps.yaml` — internal-only helper exposing a single `apt` task that takes a `DEP` var and idempotently installs it. Other Taskfiles include this as `deps:` (marked `internal: true`) and call `task: deps:apt` with `vars: { DEP: … }`. When adding a new apt-installable dependency, route it through this helper rather than calling `sudo apt install` directly so the `dpkg --get-selections` status check is consistent.

### Version pinning

Tool versions live as `vars:` at the top of `TaskfileAsdf.yaml` (`ASDF_VERSION`, `PYTHON_VERSION`, `GOLANG_VERSION`, `LAZYDOCKER_VERSION`, `PRE_COMMIT_VERSION`). To bump a version, change the var — the `status:` checks (`asdf current | grep …`) will then detect the mismatch and re-run the install. Dependabot is configured (`.github/dependabot.yml`) for GitHub Actions and Go modules only; asdf-managed versions are bumped manually.

### Conventions

- `.editorconfig` enforces 4-space indent and LF endings for `*.sh` only; YAML files are not covered.
- `.gitattributes` sets language-aware `diff=` drivers for `.go`, `.py`, `.md`, `.sh`. The git `config` task copies this file to `~/.gitattributes` and points global `core.attributesfile` at it, so the same diff drivers apply outside this repo.
