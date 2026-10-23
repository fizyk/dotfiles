#!/usr/bin/env bash
set -euo pipefail

repo_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
claude_dir="${CLAUDE_CONFIG_DIR:-$HOME/.claude}"

link() {
    local source="$1" target="$2"
    if [[ -e "$target" && ! -L "$target" ]]; then
        echo "Refusing to replace $target: not a symlink. Move it aside first." >&2
        exit 1
    fi
    ln -sfn "$source" "$target"
    echo "$target -> $source"
}

mkdir -p "$claude_dir/skills"
link "$repo_dir/CLAUDE.md" "$claude_dir/CLAUDE.md"
for skill in "$repo_dir"/skills/*/; do
    [[ -d "$skill" ]] || continue
    link "${skill%/}" "$claude_dir/skills/$(basename "$skill")"
done
