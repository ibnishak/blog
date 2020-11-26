---
title: "Awesome WM Popups"
date: 2019-10-12
tags: ["awesomewm"]
author: Riz
draft: true
---

* Pointless introduction.
[[https://awesomewm.org/][Awesome window manager]] is a highly configurable window manager for [[https://www.gnu.org/][GNU+Linux]] platform. One of its features is the ability to show a popup listing all the hotkeys we have configured and the descriptions of those hotkeys. In the default =rc.lua=, the popup only lists hotkeys configured in the rc.lua file or its linked lua files. However, AWM permits showing popups listing hotkeys not configured in =rc.lua= too, and show them only when particular clients are active on your screen. This is useful when you want to have a sneak peek on your extensive emacs keymaps or learning your way around firefox shortcut keys. 
* Steps
1. Open your AWM configuration directory. It is usually situated in =~/.config/awesome=
2. Create a file named =hkhelp.lua=. You can give any name you like to the file. Just remember to change the filename in all the steps where the filename here is mentioned.
3. Open your =rc.lua= and find the line line that says 
#+BEGIN_SRC lua
local hotkeys_popup = require("awful.hotkeys_popup").widget
#+END_SRC
#+BEGIN_alert
   If you cannot find the line, add it.
#+END_alert

4. Add the following line *after* the above line.
#+BEGIN_SRC lua
require("hkhelp.lua")
#+END_SRC

So far we have imported an empty file to be loaded whenever your =rc.lua= is loaded.

5. Open the file =hkhelp.lua=
6. Add the following code
#+BEGIN_SRC lua
local hk = require("awful.hotkeys_popup")
#+END_SRC
