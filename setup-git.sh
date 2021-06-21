echo "Configure git"
echo "Set git attributes"
# Based on https://tekin.co.uk/2020/10/better-git-diff-output-for-ruby-python-elixir-and-more
# Existing diffs https://github.com/git/git/blob/master/userdiff.c
cp .gitattributes ~/.gitattributes
git config --global core.attributesfile ~/.gitattributes
git config --global alias.cleanup "!git branch --merged | grep  -v '\\*\\|master\\|dev' | xargs -n 1 git branch -d"
git config --global alias.cleanup-remote "!git fetch origin -p && git branch -r --list 'origin/*' --merged origin/master | egrep -v '(^\*|master|dev)' | sed 's/origin\///' | xargs -n 1 git push origin --delete"

# Configure git for golang to treat repositories as git ones not mercurial:
git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
echo "Git configuration completed"
