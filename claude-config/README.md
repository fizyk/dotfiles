# claude-config

Personal Claude Code rules and skills, shared by every machine and sandbox.

| Path | Installed as |
| --- | --- |
| `CLAUDE.md` | `~/.claude/CLAUDE.md`, loaded in every session |
| `skills/<name>/` | `~/.claude/skills/<name>/`, available in every repository |

## Install

```shell
./install.sh
```

From the dotfiles root, `task claude:config` (part of `task config`) runs the
same script and skips it when every link is already in place.

It symlinks each file into `~/.claude` (or `$CLAUDE_CONFIG_DIR`), so a
`git pull` updates every session. It refuses to replace a file that is not
already a symlink.

Run it once on the host and once inside each sandbox whose `~/.claude` does
not survive a rebuild. The sandbox must see this checkout at the path the
symlinks point to.

## Precedence

A repository's own `CLAUDE.md` loads after this one and wins where they
disagree.

A personal skill wins over a project skill with the same name, so every skill
here carries the `my-` prefix. A repository's own `awsome-skill` keeps running
as its team wrote it, and `my-awsome-skill` stays callable beside it.
For a repository you do not own, put personal additions in
`CLAUDE.local.md` and list it in `.git/info/exclude`, which keeps it out of git
without changing the repository's `.gitignore`.

## Other agents

Agents that read `AGENTS.md` (Codex, Cursor and others) can share the same
rules through a symlink. For Codex, from this directory:

```shell
ln -s "$PWD/CLAUDE.md" ~/.codex/AGENTS.md
```
