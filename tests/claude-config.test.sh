#!/usr/bin/env bash
# `claude:config` runs claude-config/install.sh, and its `status:` must agree
# with what the script leaves behind. Both are pointed at a scratch
# CLAUDE_CONFIG_DIR, so the real ~/.claude is never touched.

# shellcheck source=lib.sh
source "$(dirname "$0")/lib.sh"

CLAUDE_CONFIG_DIR="$(mktemp -d)"
export CLAUDE_CONFIG_DIR
trap 'rm -rf "$STUBS" "$CLAUDE_CONFIG_DIR"' EXIT

# up_to_date: "yes" if claude:config's status: says there is nothing to do.
up_to_date() {
    if task --dir "$ROOT" --status claude:config >/dev/null 2>&1; then
        echo yes
    else
        echo no
    fi
}

check "empty config dir needs linking" no "$(up_to_date)"

"$ROOT/claude-config/install.sh" >/dev/null
check "CLAUDE.md is linked" \
    "$ROOT/claude-config/CLAUDE.md" "$(readlink "$CLAUDE_CONFIG_DIR/CLAUDE.md")"
check "status sees the links install.sh made" yes "$(up_to_date)"

"$ROOT/claude-config/install.sh" >/dev/null
check "second run is harmless" yes "$(up_to_date)"

skill="$(basename "$(find "$ROOT/claude-config/skills" -mindepth 1 -maxdepth 1 -type d | head -n 1)")"
rm "$CLAUDE_CONFIG_DIR/skills/$skill"
check "missing skill link needs linking" no "$(up_to_date)"
"$ROOT/claude-config/install.sh" >/dev/null

rm "$CLAUDE_CONFIG_DIR/CLAUDE.md"
echo "# hand-written" >"$CLAUDE_CONFIG_DIR/CLAUDE.md"
if "$ROOT/claude-config/install.sh" >/dev/null 2>&1; then refused=no; else refused=yes; fi
check "a regular CLAUDE.md is refused, not replaced" yes "$refused"
check "the regular CLAUDE.md is left alone" \
    "# hand-written" "$(cat "$CLAUDE_CONFIG_DIR/CLAUDE.md")"

finish
