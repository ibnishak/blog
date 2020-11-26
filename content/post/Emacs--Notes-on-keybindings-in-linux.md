---
title: "Emacs: Notes on keybindings in GNU+Linux"
date: 2019-09-18
author: Riz
tags: ["emacs"]
description: "Using common shortcuts without breaking emacs"
---

# Pointless introduction

Emacs being a rather keyboard driven editor, tends to rely heavily on keyboard shortcuts. The non-conformity of keybindings in emacs is a rather tricky hurdle. While it is accepted that Emacs is practically an operating system and the true emacs-jedi can do everything in Emacs itself, the rest of mortals have to interact with other, normal softwares. Let us consider the simple tasks of cut, copy and paste.

Everyone and their grandmothers know that the shortcut for these are <kbd>Ctrl+x</kbd>, <kbd>Ctrl+c</kbd> and <kbd>Ctrl+v</kbd>. If you happen to flirt with terminals, then you need to remember that in most terminals the shortcuts are <kbd>Ctrl+Shift+c</kbd> and <kbd>Ctrl+Shift+v</kbd>. Now come to emacs, where the keybindings become <kbd>Ctrl+w</kbd>, <kbd>Alt+w</kbd> and <kbd>Ctrl+y</kbd>. The toll it takes on your muscle memory and frustration that ensues everytime you send SIGKILL instead of copying is no funny buisness.

# Plan

We are going to set up Emacs such that we can use the all the standard keybindings like: <kbd>Ctrl+a</kbd> to Select all <kbd>Ctrl+s</kbd> to Save <kbd>Ctrl+d</kbd> to add bookmark <kbd>Ctrl+f</kbd> to find <kbd>Ctrl+z</kbd> to undo <kbd>Ctrl+x</kbd> to cut <kbd>Ctrl+c</kbd> to copy <kbd>Ctrl+v</kbd> to paste

Since Emacs heavily uses Ctrl and Alt keys, we need to avoid messing with these modifier keys. So here is the plan

- Use `xmodmap` to convert <kbd>CapsLock</kbd> to <kbd>Hyper</kbd> key. So for every other software except Emacs, <kbd>CapsLock</kbd> acts as <kbd>Hyper</kbd>. All other modifier keys remain at their default position and usage.
- In Emacs and Emacs alone, swap the <kbd>Hyper</kbd> and <kbd>Ctrl</kbd> keys, so that when you press <kbd>Ctrl</kbd> button in keyboard, <kbd>Hyper</kbd> is sent to Emacs, while pressing <kbd>CapsLock</kbd> button sends <kbd>Ctrl</kbd> to Emacs.
- Finally set up keybindings like <kbd>Hyper-x</kbd> for cut, <kbd>Hyper-c</kbd> for copy in Emacs. Since Emacs is accepting <kbd>Ctrl</kbd> as <kbd>Hyper</kbd>, the actual keypress would be <kbd>Ctrl+x</kbd>, <kbd>Ctrl+y</kbd> etc

# Step 1: Xmodmap


-   Run the following command

```bash
    xmodmap -pke > ~/.Xmodmap
```

-   Open `~/.Xmodmap`. Add the following code to the file.


```bash
    clear      lock
    clear      mod1
    clear      mod2
    clear      mod3
    clear      mod4
    clear      mod5
    keycode      66 = Hyper_L
    add        mod1 = Alt_L Alt_R Meta_L
    add        mod2 = Num_Lock
    add        mod3 = Hyper_L
    add        mod4 = Super_L Super_R
    add        mod5 = Mode_switch ISO_Level3_Shift
```

- Reload the xmodmap by running

```bash
xmodmap ~/.Xmodmap
```

<div class="alert">
It is possible that keycodes vary in your keyboard. If the above codes
does not work after reloading, try using
<a href="https://linux.die.net/man/1/xev" class="alert-link">xev</a>
</div>

# Step 2: Emacs


Add the following snippet to your `init.el` in **Linux OS**. If you are using Windows/Mac, find appropriate snippet from
[ergoemacs](http://ergoemacs.org/emacs/emacs_hyper_super_keys.html).

``` {.commonlisp org-language="emacs-lisp" tangle="no"}
(setq x-ctrl-keysym 'hyper) ;;In Emacs, treat Control key as hyper
(setq x-hyper-keysym 'ctrl) ;; In Emacs, treat Hyper Key(Caps_Lock)  as Control
```

# My Keybindings

I am using [general.el](https://github.com/noctuid/general.el) for my keybindings. A subset of which looks as follows.

<div class="alert">
The following snippet depends on packages like general.el, helm,
undo-tree, helm-swoop, avy, anzu, yasnippet. You may modify the packages
as per your personal preferences
</div>

``` {.commonlisp org-language="emacs-lisp" tangle="no"}
(use-package general
  :defer 2
  :init
  (general-define-key
   :states '(normal insert visual)
   :keymaps 'override
   "M-x" 'helm-M-x
   "H-a" 'mark-whole-buffer
   "H-s" 'save-buffer
   "H-d" 'bookmark-set

   "H-z" 'undo-tree-undo
   Ctrl+Shift+z for Redo.
   "H-x" 'kill-region
   "H-c" 'copy-region-as-kill
   "H-v" 'yank

   "H-f" 'helm-swoop
   "H-F" 'helm-projectile-ag
   "H-o" 'fzf-projectile
   "H-n" 'helm-find-files ;; Ctrl+n for new files
   "H-w" 'kill-this-buffer  ;; Ctrl+w kills buffer
   "H-g" 'avy-goto-char-2
   "H-h" 'anzu-query-replace
   "H-H" 'anzu-query-replace-regexp

   "H-&lt;tab&gt;" 'helm-mini
   "H-*" 'bookmark-jump
   "H-y"  'yas-insert-snippet
   "H-`" 'helm-register
   "H-M-`" 'point-to-register
   ))
```

</div>

# 13<sup>th</sup> Function Key


Another key in the keyboard I hardly ever find using is the `Menu` key.
You know, [that
key](https://www.howtogeek.com/wp-content/uploads/2015/07/a-keyboard-shortcut-that-can-be-used-in-place-of-the-context-menu-key-01.jpg)
smugly sitting in between <kbd>Right AltGr</kbd> and <kbd>Right
Ctrl</kbd>. I configured the Xmodmap so that it will act as
<kbd>F13</kbd>. Since no normal keyboard has an <kbd>F13</kbd> key, no
shortcuts are assigned to this in any software. This gives us a array of
possibilities regarding keybindings devoid of headaches. Add the
following line to your `~/.Xmodmap`

``` {.bash org-language="sh" tangle="no"}
keycode 135 = F13
```

and reload your xmodmap by running

```
    xmodmap ~/.Xmodmap
```

# Right Alt as F13


If `Menu` key is hard to reach for you, the Right Alt key is another
option to be morphed into F13. The `~/.Xmodmap` will look like this:

``` {.bash org-language="sh" tangle="no"}
clear      lock
clear      mod1
clear      mod2
clear      mod3
clear      mod4
clear      mod5
keycode      66 = Hyper_L
keycode     108 = F13
add        mod1 = Alt_L Meta_L
add        mod2 = Num_Lock
add        mod3 = Hyper_L
add        mod4 = Super_L Super_R
add        mod5 = Mode_switch ISO_Level3_Shift
```

and reload your xmodmap by running

```bash
    xmodmap ~/.Xmodmap
```

# Unrelated tip


Since you are in the business of modifying `~/.Xmodmap`, you may want to
know that it can be also used to change the scrolling using mouse to
[natural
scrolling](https://jessequinnlee.com/2015/07/25/natural-scrolling-vs-reverse-scrolling/)
like those seen in Mac, as opposed to the default reverse scrolling.
Here is the line to add to `~/.Xmodmap`.

```bash
    pointer = 1 2 3 5 4 7 6 8 9 10 11 12
```

and reload your xmodmap by running

```bash
    xmodmap ~/.Xmodmap
```
