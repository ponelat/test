# Bail on first error...
set -e

# globals: $TRAVIS_BRANCH, $GH_TOKEN
# PS the following is a plain text password, delete before commiting or pushing
GITHUB_URL=https://ponelat:$GH_TOKEN@github.com/ponelat/test.git

git remote remove origin
git remote add origin "$GITHUB_URL"

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
