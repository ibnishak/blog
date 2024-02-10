---
title: Fcron - Set up, run jobs and troubleshoot
date: 2022-10-03T14:50:05+05:30
lastmod: 
author: Riz

description: 
categories: ["digital"]
tags: ["linux", "fcron", "cron"]

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


## What to do in case of permission errors

Remember that the folder `/var/spool/fcron` should be owned by the user `fcron` while the fcrontab files inside that folder should be owned by respective users. So if you have permission issues, the following commands may resolve the issues

```bash
sudo chown fcron /var/spool/fcron
sudo chown $USER /var/spool/fcron/${USER}.orig
```