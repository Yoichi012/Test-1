#!/bin/bash

# Github Quick-Push Script
echo "=============================="
echo "Github Quick-Push"
echo "=============================="
echo ""

# Print the current branch
echo "Pushing to branch:"
git branch | grep \* | cut -d ' ' -f2
echo ""

# Get commit message from user
read -p "Enter Commit title (pushes with you as author): " commit_title

echo ""
echo "Pulling latest changes..."
git pull

echo "Adding files..."
git add .

echo "Committing changes..."
git commit -m "$commit_title"

echo "Pushing to remote..."
git push

echo ""
echo "Push completed!"
