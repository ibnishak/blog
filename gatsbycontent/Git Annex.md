---
title: "Git Annex - Setup and Workflow"
tags:
up:
aliases:
status:
id: 20220222151700769
---

# Git Annex - Setup and Workflow

## What is git annex?
 - [Git Annex](https://git-annex.branchable.com/) is an open source software written in haskell. It provides a layer over Git to manage large and binary which are otherwise not checked into Git.
 - For me, git-annex is a software that enables me to see what all files are stored in all my hard-drives when they are not connected. It also acts as an interface to retrieve those files, add files, and create multiple copies stored in different hard-drives.
 
## How to install git annex in Arch Linux
- Install `git-annex-standalone`  from pacman or download linux version from its webpage and run install script. 
	 - DO NOT install `git-annex` - ONLY install `git-annex-standalone`. The former will install all haskell libraries which is heavy and reuqires frequent updates.

## How to set up folders in laptop and external-HDD for synchronizing?
- Create a source folder in laptop, and destination folder in external-HDD
 - Run init commands in each folder and add complementary folder as remote
```bash
cd source-folder
git init
git annex init "source"

cd destination-folder
git init
git annex init "destination"

cd source-folder
git remote add "destination" destination-folder

cd destination-folder
git remote add "source" source-folder
```
## How to move all files to destination-folder, leaving only metadata in source-folder
- Add files to git-annex in source-folder, run sync from each folder and move.
```bash 
cd source-folder
git annex add .
git commit -a -m "My commit message"
git annex sync
git annex move . --to "destination" 

cd destination-folder
git annex sync
```
## What special instructions does calibre-library folder require during setup and syncing
- During setup, add the following lines to `.gitignore` files in both source-folder and destination folder.
```
metadata_db_prefs_backup.json
metadata.db
-.jpg
-.opf
```
- Also during setup, run the following commands once from source-folder and once from destination folder.
```bash
git config annex.largefiles "include=* exclude=*.opf exclude=*.json exclude=*.db exclude=*.jpg"
```
- During synchronization, run the following command while in source-folder
```bash

rsync -av --delete --exclude=".git/" --include '*.jpg' --include '*.opf' --include 'metadata*' --include="*/" --exclude="*"  . destination-folder
```
## How to retrieve a file from external-HDD?
- Run `get` command from source folder
```bash
cd source-folder
git annex get "myfile"
```
## How to find the location where files is stored?
- Run whereis command from any of git-annex folders
```bash
git annex whereis "myfile"
```

[1]: https://git-annex.branchable.com/