---
title: Dotfile Management - Criteria and Candidates
date: 2020-12-10T22:00:43+05:30
author: Riz
categories: ["review"]
tags: ["linux"]
---

## Essential
- NOT symlinking. Some programs (eg: XFCE) will overwrite symlinks.
- NOT copying. The config files itself should be the source of truth. ie, If the config files are edited, the changes should automatically reflect in the dotfiles vault.
- Version control. If the config files ought be edited without worry, they need be reversible to previous states
- Scriptable

## Desirable
- Simple command structure
- Ability to encrypt files
- Ability to ignore files based on OS or other parameters
- Bootstrapping - Ability to run some scripts automatically on first run.

## Ideal
- Automated synchronisation to cloud targets on regular intervals (like an inbuilt cron)
- Part of core packages so that it is immediately available upon OS reinstall

## Candidates
- Gnu stow: Symlinks
- [Chezmoi](https://github.com/twpayne/chezmoi): Copying
- [RCM](https://github.com/thoughtbot/rcm)
- [Comtrya](https://github.com/comtrya/comtrya)
- [dotfiles](https://github.com/rhysd/dotfiles): Symlinks
- [dotbot](https://github.com/sigmonsays/dotbot): Symlinks
- [Yadm](https://yadm.io/): CURRENT