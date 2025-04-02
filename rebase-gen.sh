#!/bin/bash
# Automated rebase script for seaqt-gen that regenerates gen submodules.
# To be used after other conflicts related to the upgrade have been resolved.
#
# Usage:
#   ./rebase-gen.sh <upstream>

set -euo pipefail

die() { echo "ERROR: $1" >&2; exit 1; }

UPSTREAM="${1:?Usage: $0 <upstream>}"
git rev-parse --verify "$UPSTREAM" >/dev/null 2>&1 || die "upstream '$UPSTREAM' does not exist"

BRANCH="$(git symbolic-ref --short HEAD)"
[ -n "$BRANCH" ] || die "not on a branch (detached HEAD)"

echo "=== Rebase $BRANCH onto $UPSTREAM ==="
echo ""

# Convert all 'pick' to 'exec' so each commit runs in the loop below
GIT_SEQUENCE_EDITOR="sed -i -re 's/^pick /e /'" git rebase -i "$UPSTREAM"  >/dev/null 2>&1

GENCOMMITS=false

while [ -d .git/rebase-merge ] || [ -d .git/rebase-apply ]; do
    COMMIT_MSG="$(tail -1 .git/rebase-merge/done)"

    if [[ "$COMMIT_MSG" == *"Introduce \`gen\` submodules"* ]]; then
        rm -rf gen/nim-seaqt-* gen/seaqt-*
        echo "[$(date +%H:%M:%S)] Structural: $COMMIT_MSG"
        make reset-gen >/dev/null 2>&1 || die "make reset-gen failed"
        GENCOMMITS=true

    elif [[ "$COMMIT_MSG" == *"Add nim-seaqt submodules"* ]]; then
        echo "[$(date +%H:%M:%S)] Structural: $COMMIT_MSG"
        make reset-gen-nim >/dev/null 2>&1 || die "make reset-gen-nim failed"
        GENCOMMITS=true

    elif $GENCOMMITS; then
        echo "[$(date +%H:%M:%S)] $COMMIT_MSG"

        UNMERGED="$(git diff --name-only --diff-filter=U 2>/dev/null || true)"
        if [ -n "$UNMERGED" ] && echo "$UNMERGED" | grep -qv '^gen/'; then
            echo "  ERROR: Conflicts in non-gen files!" >&2
        fi

        make gencommits >/dev/null 2>&1 || die "make gencommits failed"
        git add gen/* >/dev/null 2>&1 || true
    fi

    GIT_EDITOR=true git rebase --continue >/dev/null 2>&1 || true
done

echo ""
echo "=== Rebase complete! ==="
