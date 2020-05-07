---
title: Workflow for updating documentation of tiddlywiki
tags: [TiddlyWiki5]
date: 2020-05-07 05:26:59 +0530
---

- Fork it and clone it your system

```bash
cd TiddlyWiki5
git checkout -b tiddlywiki-com origin/tiddlywiki-com
git checkout -b dev
cd editions/tw5.com-server
tiddlywiki --listen
```
- Make changes and commit

```bash
git checkout tiddlywiki-com
git merge dev
git push
```
