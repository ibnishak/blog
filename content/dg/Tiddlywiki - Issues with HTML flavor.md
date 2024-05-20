---
title: Tiddlywiki - Issues with HTML flavor
---

# Tiddlywiki - Issues with HTML flavor

In HTML flavor, all the data is saved in a single HTML file. This is a unique Tiddlywiki approach and a radical one. Every other note taking solution including markdown, requires some sort of parser. You might find yourselves in a position where you don't have the privileges to install a software in a computer, thereby cutting you off from your notes. However it is a safe bet that any of those machines will have a web browser installed, and that is all you need to read your Tiddlywiki notes in single HTML flavor.

The nitty-gritty details are not rosy though. You have all the data in a single file - essentially a database with no syncing API. So you need to sync the entire file between devices. Imagine you have a Tiddlywiki that is 25MB in size (remember h0p3's wiki as 15K notes and is 50MB in size). Each minor correction in any of those notes require re-upload of 25MB. By your 1000th modification, you have consumed 25GB bandwidth as opposed to 10MB for a 10KB file.
Consider a syncing conflict. It is certainly not impossible, but definitely harder to diff two large HTML files than two small text files. Git does not look kindly upon a >2MB HTML file modified in every commit. A partial solution to this issue is using Timimi - but it is more of a backup solution than version control.

For a user who is familiar with the ease tools like fzf, fd etc brings to workflow, single HTML file might not be palatable. Again it is not impossible to modify the Tiddlywiki HTML file using the familiar scripting tools, just harder.
