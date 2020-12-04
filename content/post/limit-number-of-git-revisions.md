---
title: Limiting Number of Git Revisions-How
date: 2020-10-10T22:42:40+05:30
lastmod: 
author: Riz

description: 
categories: []
tags: []

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

## What
Whenever a git commit is made, you are essentially creating a snapshot of the directory tree with some optimisations. This is great for software development workflow. However if you are using git to version control something like a personal knowledge base, keeping snapshots all the way back to origin is unnecessary. 

Git provides facilities to prune old commits. We make use of this facility to keep only last 5 commits. This means whenver a new commit is made, the fifth oldest commit is removed. You can verify this with `git log`.

## How

1. In your folder under git version control, create/open the file `./.git/hooks/post-commit`
2. Add the following code to it and save the file
```git
git rev-parse HEAD~1 > .git/shallow && git fsck --unreachable && git gc --prune=now
```
3. Make the post-commit hook executable
```
chmod +x .git/hooks/post-commit
```
