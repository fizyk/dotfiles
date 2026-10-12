#!/usr/bin/env bash
# The awk in `mise:tools` turns the [tools] table of a mise.toml into
# `mise use -g` arguments. A dry, forced run prints the command it would run
# without evaluating `status:` or touching the real mise.

# shellcheck source=lib.sh
source "$(dirname "$0")/lib.sh"

FIXTURES="$ROOT/tests/fixtures"

stub mise 'exit 0'

# tools_cmd FILE: the `mise use -g ...` line mise:tools would run for FILE.
tools_cmd() {
    task --dir "$ROOT" --color=false --dry --force mise:tools MISE_FILE="$1" 2>&1 |
        sed -n 's/^task: \[mise:tools\] //p' | sed 's/ *$//'
}

# Expected pins come from a real TOML parser, so this keeps passing as
# Renovate bumps versions.
check "repo mise.toml parses to one pin per tool" \
    "mise use -g $(python3 -c '
import sys, tomllib
with open(sys.argv[1], "rb") as f:
    print(" ".join(f"{k}@{v}" for k, v in tomllib.load(f)["tools"].items()))
' "$ROOT/mise.toml")" \
    "$(tools_cmd "$ROOT/mise.toml")"

check "only [tools] is read; comments, blanks and spacing are ignored" \
    "mise use -g go@1.2.3 python@3.14.0 aqua:jesseduffield/lazydocker@0.25.2 node@latest" \
    "$(tools_cmd "$FIXTURES/mise-messy.toml")"

finish
