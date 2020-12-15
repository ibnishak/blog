---
title: Arch Linux Sudoers
date: 2020-12-15T20:53:19+05:30
lastmod: 
author: Riz

description: 
categories: ["digital"]
tags: ["linux"]

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

# Arch Linux: Sudoers

Sudoer files is the configuration file for sudo command. It is located at `/etc/sudoers`.

## Editing Sudoer
It **MUST** be edited using the command `sudo visudo`. Behind the scenes this command creates a temporary file, checks for any errors once you save and finally replace the original file.

However editing in vi can be a pain. You can change it to any terminal editors using the command in the form below.

```
su
VISUAL=nano visudo
```

In case that is not working, you can symlink nano to vi

```
sudo ln -s /usr/bin/nano /usr/bin/vi
sudo visudo
```

Some changes made to configuration are as follows:

```
Defaults passwd_timeout=0
Defaults passprompt="^G[sudo] password for %p: "
Defaults insults
```

1. Do not timeout when asking for password. Especially helpful when running commands like `yay -Syu` where password is asked only after  downloads are done and if you do not type password within some time, it will timeout. 
1. Prompt Template for password - ^G is a bell symbol so it will alert you.
1. Throw insults if the password is wrong

For more tricks, See [Archwiki][1]

[1]: https://wiki.archlinux.org/index.php/sudo#Tips_and_tricks