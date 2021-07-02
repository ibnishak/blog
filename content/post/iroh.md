---
title: Project Iroh
date: 2021-06-22T08:33:25+05:30
lastmod:
author: Riz

description:
categories: []
tags: []

draft: true
enableDisqus: false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

Project Iroh is a collection of scripts, programs and extensions which is used to convert a bunch of markdown files to this website. While the code is still messy, needs reorganization and optimization, I think it would be a good idea to document various parts of it.

## Rationality and Justification of choices

## The Core

At the core of Iroh is a command line interface written in Golang. It provides 3 commands

1. init
   Init command takes user input regarding where the markdown notes are stored, where do you want to keep the published notes etc. It would then create a `.meta` folder
2. parse
   Parse co

As mentioned above, it is a bit too messy for public consumption right now.

## VSCode and Plugins

Contributions of various vscode plugins to Iroh

- Markdown Linkify

  - Markdown: Linkify command, Ctrl+L shortcut, Creat the newly linked file if they do not exist

- Mayukai dark theme

## Scripts and Other components

At its core is a golang program I wrote.
