---
title: How to set up and run jobs with Fcron
date: 2022-10-03T14:50:05+05:30
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

- Install Fcron
- Ensure that an Fcron group is created
```bash
cat /etc/group | grep fcron
# Should output something like this: fcron:x:23
```
- Add current user to Fcron group (Need Root)
```bash
sudo usermod -a -G fcron $USER
```

You might have to re-login after the last command.

- Open the fcrontab

```bash
fcrontab -e
```

- An empty text file should open. This is where we define what tasks to run and when to run them. However, first we add the necessary environment variables

```bash
SHELL=/usr/bin/zsh
DISPLAY=:0

```
- Now we can tasks one per line. For eg, below given task will change wallpaper to a random picture in wallpapers folder every 5 mins (Need `feh` installed)

```bash
@ 5 /usr/bin/feh --no-fehbg --bg-fill --randomize ~/Collection/Pictures/wallpapers
```

