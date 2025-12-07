#!/bin/bash

# Script to run GitHub Actions locally using act
# This allows you to test your workflows before pushing

set -e

echo "🚀 Running GitHub Actions locally with act..."
echo ""

# Check if act is installed
if ! command -v act &> /dev/null; then
    echo "❌ act is not installed. Install it with: brew install act"
    exit 1
fi

# List available workflows
echo "📋 Available workflows:"
act -l

echo ""
echo "💡 To run a specific workflow, use:"
echo "   act -W .github/workflows/ping-ci.yml"
echo ""
echo "💡 To run without Docker push (recommended for local testing):"
echo "   act -W .github/workflows/ping-ci.yml --skip-push"
echo ""
echo "💡 To run a specific job:"
echo "   act -j build-and-deploy"
echo ""
echo "💡 To run with secrets (create .secrets file first):"
echo "   act --secret-file .secrets"
echo ""
echo "🔧 Running workflow now (skipping push step)..."
echo ""

# Run the workflow, skipping the push step for local testing
# Use linux/amd64 architecture for Apple Silicon compatibility
if [[ $(uname -m) == "arm64" ]]; then
    echo "🍎 Detected Apple Silicon - using linux/amd64 architecture"
    act push -W .github/workflows/ping-ci.yml --skip-push --container-architecture linux/amd64
else
    act push -W .github/workflows/ping-ci.yml --skip-push
fi

