---
title: Converting CapsLock to Hyper key with XKB
date: 2020-12-26T22:54:47+05:30
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

Warning: While doing this, my cat decided to take a walk over my keyboard because I ignored her attempts to gain my attention for sometime. I saved the file not realizing that there were some syntax errors due to the cat walk. Next time I tried to login, my keyboard wouldn't work. I couldn't switch to console. The arch linux fall back did not help either. Finally I had to do a clean install of my OS. So please do not ignore your cat and ensure you do not have any syntax errors.


## Why
I needed to set a bunch of keyboard shortcuts that will not conflict with other softwares or my window manager. So control key and alt key was out of question. Super key is in an awkward position in the keyboard. Converting capslock to a hyper key seemed like the logical thing to do.

You can achieve this fairly easily with Xmodmap. However it is an outdated method. It also has the disadvantage that any new USB keyboard will completely reset it. For me personally, the complete reset was being triggered by something else too, randomly. After some hair pulling attempts to figure out what, I decided it is time to ditch the xmodmap.

## How
- Make a copy of your existing layout and open the copy with your editor - or nano which is always reliable.
    - NOTE: Ideally you should be able to add just the required code to a separate file and load that file in addition to the usual layout. However all the tutorials on this subject did not work for me.

```
cd /usr/share/X11/xkb/symbols
sudo cp us hyper
sudo nano hyper
```

- Add the following code within the `{ }`s of `xkb_symbols "basic"`.
```
    key <CAPS> {[ Hyper_L ]};
    key <HYPR> {[ Hyper_R , Hyper_L	]};
    modifier_map Mod3   { <HYPR> };
```
- Ensure that there are no syntax errors
```
 setxkbmap -layout hyper
```
If there are syntax errors, the commnad will refuse to load the layout - a lesson I learned painfully as stated above.

The Capslock should be converted to Hyper now. You can test it by setting shortcuts to something in your window manager.

- `setxkbmap` only loads the layout for current session. To make the changes permanent, use the following command
```
localectl --no-convert set-x11-keymap hyper
```