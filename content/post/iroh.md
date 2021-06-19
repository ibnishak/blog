---
title: Project Iroh
date: 2021-06-18T08:33:25+05:30
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

- Markdown IDE

  - type `[` to trigger autocompletion for entering links to other Markdown files
  - type `![` to trigger autocompletion for links to local image files.
  - updates links when renaming files
  - removes links when deleting files

- Mzettel

  - Mzettel - Note: Asks for a title. Creates a new note with **slugified title + extension** as filename and adds title as content.
  - Mzettel - Note Link to Clipboard: Copies markdown link to clipboard in the format `[Title](filename)`
  - Mzettel - Open in Browser: Opens current note in browser localhost.

- Markdown All in One

  - Ctrl+B, Ctrl+I shortcuts for bold
  - Paste over as link
  - List management [Changed order of list shortcuts]
  - Toggle checkboxes, Toggle checkbox status
  - Toggle preivew and close it with same shortcut
  - Shortcuts for increasing and decreasing heading levels

- Markdown notes

  - Wikilinks, Create new files from wikilinks, Wikilinks in preview
  - Hover peek

- Markdown table formatter

  - Formats table
  - Provide lenses for sorting the table by different columns
  - Provides below settings as customizable features
    - Default justification
    - Keep first and last pipes
    - Space padding of columns
    - Remove colons if same as default

- Paste image
- Md-graph
- Spell check
- Which-key
- prettier
- Mayukai dark theme

## Scripts and Other components

At its core is a golang program I wrote.
