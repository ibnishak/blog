---
title: Various ways to Benchmarking neovim startup
date: 2022-09-29T01:38:09+05:30
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


## Benchmark the effect of colorscheme on your startup

Run the following. 

- Replace `neovide` with whatever your Neovim gui is
- Replace `onedark` in the second command, with the name of your color scheme


```
neovide -- -Nu NONE --startuptime startup.txt +colorscheme\ default +q
neovide -- -Nu NONE --startuptime startup.txt +colorscheme\ onedark +q
```
