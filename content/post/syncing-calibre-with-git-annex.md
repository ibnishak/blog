---
title: Syncing Calibre With Git Annex
date: 2021-07-14T02:21:07+05:30
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

Assuming that you want to keep only metadata and covers in your laptop while the books itself in external hard drive.

Dependencies: Rsync, Git-Annex

- Place script in PATH
- Edit the USER VARIABLES
- Run the `init` command once
- Run the `sync` command in all consequent runs
- See help with `--help`

```zsh
#!/usr/bin/env zsh
set -e # Exit on error

#---------USER VARIABLES---------------#
SRCDIR="${HOME}/Documents/Library"
TARGETDIR="${HOME}/External/Library"
#--------------------------------------#

declare -r _VERSION=2.3.4
declare -r _SCRIPT=$0

# Init function takes 4 params dir1, name1, name2, dir2 in that order
function setup(){
	if [ ! -d $1 ] 
	then
			mkdir $1
	fi
	pushd $1
	git init
	git annex init $2
<<END > .gitignore
metadata_db_prefs_backup.json
metadata.db
*.jpg
*.opf
END
	git config annex.largefiles "include=* exclude=*.opf exclude=*.json exclude=*.db exclude=*.jpg"
	git remote add $3 $4 
	git annex add . 
	git commit -a -m "$(date)" 
	popd
}


func init() {
	setup $SRCDIR "SRC" "TARGET" $TARGETDIR
	setup $TARGETDIR "TARGET" "SRC" $SRCDIR
}

func sync() {
	pushd $SRCDIR
	git annex add . 
	git commit -a -m "$(date)" 
	git annex sync 
	git annex move . --to TARGET 
	rsync -av --delete --exclude=".git/" --include '*.jpg' --include '*.opf' --include 'metadata*' --include="*/" --exclude="*"  . $TARGETDIR 
	popd

	pushd $TARGETDIR
	git annex sync 
	popd
}

func helpmenu(){
	echo "This script has 2 commands"
	echo "${_SCRIPT} init"
	echo "${_SCRIPT} sync"
}

if [ "$1" = "--help" ]; then
    helpmenu
    exit 0
elif declare -f "$1" > /dev/null; then
    "$@"
else
    echo "'$1' is not a known command" >&2
    helpmenu
    exit 1
fi
```