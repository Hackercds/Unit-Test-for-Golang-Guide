#!/bin/sh
# Enable pre-commit hooks by setting the hooks path
cd "$(dirname "$0")/.."
git config core.hooksPath .githooks
echo "Hooks enabled. Pre-commit checks will run on git commit."
echo "To skip hooks for a single commit: git commit --no-verify"
