---
title: "Notes on Accessories Folder"
tags:
up:
aliases:
status:
id: 20220317152036516
---

# Notes on Accessories Folder

Accessories folder is served over a localhost port using a basic HTTP server. Images and Videos are then embedded in notes using the URL.
This affords me the following conveniences:

1. Keep heavy files out of obsidian vaults. This leaves just text in obsidian vaults which I can encrypt and store in any number of free cloud storage services. See [[Rclone]]
2. No need to worry about a bunch of ignore rules for backup and version control of vaults. The Accessories folder is backed up weekly to external hard drives weekly.
3. Use images and videos in other software (e.g.: Anki) without concern about their exact location in the file system.

Obsidian config that is shared between vaults is also kept in Accessories folder and symlinked to each vault. Shared configs include plugins, hotkeys, CSS snippets, and themes. Vault specific config, which remains within vault, controls things like which plugin/themes etc. are active in that specific vault.