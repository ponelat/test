# Bail on first error...
set -e

# $GITHUB_TOKEN=
# $GITHUB_URL=https://ponelat:$GITHUB_TOKEN@github.com/ponelat/test
# globals: $TRAVIS_BRANCH

# git remote rm origin
# git remote add origin "$GITHUB_URL"

git checkout -f master

echo `date` >> touched

# Delete bot_tested branch
git branch -D --quiet bot_tested
git push origin :bot_tested

# Recreate bot_tested branch
git checkout -b bot_tested
# Add files and commmit
git add touched
git commit -m 'Just a touched commit'

# Push it back to origin
git push -u origin bot_tested
