# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository purpose

Personal dotfiles repo that bootstraps a developer workstation. There is no application code here — the repo is a set of [Taskfile](https://taskfile.dev) definitions that drive `apt`, `curl`, and `git clone` to install and configure a fixed list of tools. It targets Ubuntu/Debian (uses `apt`, `dpkg`, `lsb_release`, `add-apt-repository`).

## Common commands

All commands run from the repo root via the `task` binary (must be installed first; see README).

- `task --list-all` — discover every task across all included Taskfiles.
- `task install` — run every tool's `install` (git, zsh, docker, mise, atuin) in order. This is also the default target.
- `task config` — run every tool's `config` (git, zsh, docker, mise, task). Atuin has no `config` task.
- `task <ns>:install` / `task <ns>:config` — operate on one tool only, where `<ns>` is `git`, `zsh`, `docker`, `mise`, `atuin`, or `task`.
- `task mise:tools` — install/pin the mise-managed runtimes (go, python, uv, prek, lazydocker) at the versions declared in `TaskfileMise.yaml` vars.
- `task atuin:update` — update an already-installed atuin.

Tasks are idempotent: each defines `status:` / `preconditions:` checks (e.g. `dpkg --get-selections | grep …`, `test -d …`, `mise current <tool> | grep …`) so re-running skips work that's already done. When editing or adding tasks, preserve this property — add a `status:` check that detects the post-condition rather than relying on the command itself being safe to re-run.

## Architecture

The top-level `Taskfile.yaml` is a thin aggregator. It `includes:` one Taskfile per tool and exposes only meta-targets (`install`, `config`, `default`). All real work lives in the per-tool files:

- `TaskfileGit.yaml` — adds the `git-core/ppa`, installs git, copies `.gitattributes` to `~/`, sets global aliases (`cleanup`, `cleanup-remote`) and config (default branch `main`, bitbucket SSH rewrite for Go).
- `TaskfileZsh.yaml` — installs zsh + oh-my-zsh (unattended curl installer), then in `config` clones powerlevel10k, zsh-syntax-highlighting, and zsh-autosuggestions into `${ZSH_CUSTOM:-~/.oh-my-zsh/custom}` and switches the login shell. Nerd Font installation is left as a manual step (printed to stdout).
- `TaskfileDocker.yaml` — adds Docker's official apt repository (keyring under `/etc/apt/keyrings`), installs `docker-ce docker-ce-cli containerd.io docker-compose-plugin`, then in `config` creates the `docker` group and adds `$USER` to it.
- `TaskfileMise.yaml` — installs mise via `sudo snap install mise --classic` (the `clean` dep first removes any older `/usr/local/bin/mise` left by the previous `mise.run` install method). The `config` task appends `eval "$(mise activate zsh)"` to `~/.zshrc`. The `tools` task parses `[tools]` out of `mise.toml` with awk and applies every pin in one `mise use -g` call — all tools come from mise's built-in registry, so no third-party plugins or apt build-deps are required (mise installs Python from python-build-standalone by default, and lazydocker comes from the `aqua:jesseduffield/lazydocker` registry entry).
- `TaskfileTask.yaml` — configures `task` line completion in zsh by appending `eval "$(task --completion zsh)"` to `~/.zshrc`.
- `TaskfileAtuin.yaml` — runs the upstream `setup.atuin.sh` installer; requires `~/.zshrc` to exist.
- `TaskfileDeps.yaml` — internal-only helper exposing a single `apt` task that takes a `DEP` var and idempotently installs it. Other Taskfiles include this as `deps:` (marked `internal: true`) and call `task: deps:apt` with `vars: { DEP: … }`. When adding a new apt-installable dependency, route it through this helper rather than calling `sudo apt install` directly so the `dpkg --get-selections` status check is consistent.

### Version pinning

Tool versions live in `mise.toml` at the repo root, not in the Taskfile. To bump one, edit that file — `task mise:tools` compares each pin against `mise current <tool>` and re-runs `mise use -g` for the whole set on any mismatch. `mise use -g` merges into `~/.config/mise/config.toml`, so tools pinned there outside this repo survive the run.

Two consequences of that layout:

- `mise.toml` is a *reference* copy, not a symlink or a mirror of the global config. The global config may legitimately hold more than this file does.
- Unpinned entries (`gh = "latest"`) can't be compared against a resolved version, so the `status:` check skips them; they are only refreshed when some other pin changes.

Renovate's built-in `mise` manager parses `mise.toml` with no custom-manager config, which is why the pins live in that format. Only the first version listed per tool is updated, so don't add fallback versions. `.github/dependabot.yml` declares `github-actions` and `gomod`, but neither has any files to scan in this repo.

### Conventions

- `.editorconfig` enforces 4-space indent and LF endings for `*.sh` only; YAML files are not covered.
- `.gitattributes` sets language-aware `diff=` drivers for `.go`, `.py`, `.md`, `.sh`. The git `config` task copies this file to `~/.gitattributes` and points global `core.attributesfile` at it, so the same diff drivers apply outside this repo.
