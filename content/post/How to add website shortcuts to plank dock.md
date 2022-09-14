---
title: How to Add Website Shortcuts to Plank Dock
date: 2022-09-14T20:24:33+05:30
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

- Create a file named `~/.local/share/applications/Youtube.desktop` with following content. Remember to replace `myusername` appropriately.
```
[Desktop Entry]
Type=Application
Icon=/home/myusername/.local/share/applications/youtube.svg
Name=Youtube
Exec=firefox-developer-edition https://www.youtube.com/
Terminal=false
Hidden=false
Keywords=youtube;video;entertainment
Categories=Internet
Comment=Youtube
```
- Quit plank.
- Create a file named `~/.config/plank/dock1/launchers/Youtube.dockitem` with following content. Remember to replace `myusername` appropriately.
```
[PlankDockItemPreferences]
Launcher=file:///home/myusername/.local/share/applications/Youtube.desktop

```
- Restart plank

Things to note:
- Use a helper program to make desktop files. I recomment Arronax
- In the desktop entry,  `Keywords`, `Categories` and `Comments` are optional
- If you are adding `Keywords`, separate them with semicolon