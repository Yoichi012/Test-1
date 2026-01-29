#!/bin/bash

# Github Quick-Pull Script
echo "=============================="
echo "Github Quick-Pull"
echo "=============================="
echo ""

# Print the current branch
echo "Pulling from branch:"
git branch | grep \* | cut -d ' ' -f2
echo ""

# Pull latest changes
git pull

echo ""
echo "Pull completed!"
