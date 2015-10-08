#!/bin/bash

VER="$1"
BRANCH="${2:-develop}"

echo Bumping to $VER on $BRANCH

git fetch
git checkout $BRANCH
git reset --hard origin/$BRANCH

npm version $VER

git push origin $BRANCH
