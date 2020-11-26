---
title: "Vim: Copy and paste to termux clipboard"
date: 2019-12-12
tags: ["android"]
author: Riz
---

Copying to system clipboard from vim is usually done by using the "+ register. Unfortunately this will not work in termux. Here is a workaround to cut, copy and paste from android clipboard. Remember that you need termux-apk tools for this to work.

- Add the following to vimrc
```vim
vnoremap <C-x> :!termux-clipboard-set<CR>
vnoremap <C-c> :w !termux-clipboard-set<CR><CR>
inoremap <C-v> <ESC>:read !termux-clipboard-get<CR>i
```
- Source the vimrc or reload vim.
- In visual mode, you can now use "Ctrl-x" to cut, "Ctrl-c" to copy to termux clipboard.
- In insert mode, you can use "Ctrl+p" to paste from termux clipboard.
- You might need hackers keyboard or Multilingo keyboard for control keys
