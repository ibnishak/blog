---
title: Ways to Benchmark neovim startup
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

## Use builtin statup time profiler

Run the following snippet in your terminal. Replace neovide with the nvim gui you are using.
```bash
neovide -- --startuptime vim-startup-times.txt
```

It will create a file calle vim-startup-times.txt in the current path detailing the effect of each plugin, file read and other things. 

## Using vim-startup time

Download and install vim-startuptime from [Github](https://github.com/rhysd/vim-startuptime) and run it as 

```bash
vim-startuptime -vimpath nvim
```

It will automatically start and close vim 10 times (configurable) and measure statup times each time. In the end it will show you average times, slowest startup time and fastest startup time in those 10 attempts.

There are multiple other repos doing the same in github. I do not have experience with them however you are welcome to try

1. https://github.com/dstein64/vim-startuptime
2. https://github.com/szkny/vim_startuptime
3. https://github.com/edwardsnjd/slow-vim


## Use packer profiler to measure impact of plugins

If you are using packer as plugin manager, open neovim and run the following command in normal mode

```bash
:PackerCompile profile=true
```

Now close neovim and reopen it. After that run the following command in normal mode

```bash
:PackerProfile
```

It will show the impact of each plugin on startup. Consider lazy loading the slow ones.

## Benchmark the effect of colorscheme on your startup

Run the following. 

- Replace `neovide` with whatever your Neovim gui is
- Replace `onedark` in the second command, with the name of your color scheme


```
neovide -- -Nu NONE --startuptime startup.txt +colorscheme\ default +q
neovide -- -Nu NONE --startuptime startup.txt +colorscheme\ onedark +q
```

