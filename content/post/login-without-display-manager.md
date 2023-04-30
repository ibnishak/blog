---
title: Login Without Display Manager
date: 2023-04-30T14:10:18+05:30
lastmod: 
author: Riz

description: 
categories: []
tags: [startx, sddm, gdm, light-dm]

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

Add the following to `/etc/zsh/zprofile`

```bash
if [[ -z "$DISPLAY" ]] && [[ $(tty) = /dev/tty1 ]]; then
	exec startx
fi
```