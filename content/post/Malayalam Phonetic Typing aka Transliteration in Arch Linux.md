---
title: Malayalam Phonetic Typing aka Transliteration in Arch Linux
date: 2022-09-04T19:52:43+05:30
lastmod: 
author: Riz

description: 
categories: []
tags: [malayalam, digital]

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

1. Install iBus and related libraries
```shell
yay -S ibus ibus-m17n m17n-db
``` 

2. Download Swanalekha inpput method lib and copy to to  `/usr/share/m17n/`

```bash
wget http://swanalekha.smc.org.in/m17n/ml-swanalekha.mim
sudo cp ml-swanalekha.mim /usr/share/m17n/ml-swanalekha.mim
```


3. Start ibus daemon
```bash
ibus-daemon -drxR
```

4. Open ibus preferences
```bash
ibus-setup
```

5. Switch to Input Method Tab, Click Add and search for `Malayalam - ml-swanalekha` and add it. If needed, change hotkey for switching input methods.
6. Restart ibus daemon
```bash
ibus-daemon -rdx
```

Also ensure that you have some malayalam fonts installed in your system. Fonts are available from:

- [Google Fonts](https://fonts.google.com/?subset=malayalam)
- [Swathanthra Malayalam Computing](https://smc.org.in/fonts/)