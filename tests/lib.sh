#!/usr/bin/env bash
# Shared helpers for tests/*.test.sh. Source it; don't run it.

set -euo pipefail

# shellcheck disable=SC2034  # used by the scripts that source this file
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
FAILED=0

# Directory of stub binaries put first on PATH, so tasks' preconditions and
# status checks see fakes instead of the real system.
STUBS="$(mktemp -d)"
trap 'rm -rf "$STUBS"' EXIT
export PATH="$STUBS:$PATH"

# stub NAME SCRIPT: create an executable NAME in $STUBS running SCRIPT.
stub() {
    printf '#!/bin/sh\n%s\n' "$2" >"$STUBS/$1"
    chmod +x "$STUBS/$1"
}

# check NAME EXPECTED ACTUAL
check() {
    if [[ "$2" == "$3" ]]; then
        printf 'ok    %s\n' "$1"
    else
        printf 'FAIL  %s\n      expected: %q\n      actual:   %q\n' "$1" "$2" "$3"
        FAILED=1
    fi
}

finish() {
    exit "$FAILED"
}
