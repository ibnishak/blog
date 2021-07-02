---
title: Vscode Addons for Markdown based Personal knowledge base
date: 2021-07-01T21:04:15+05:30
lastmod:
author: Riz

description:
categories: []
tags: []
draft: false
enableDisqus: false
enableMathJax: falBette
disableToC: false
disableAutoCollapse: true
---

1. Markdown All in One
   - It provides basic shortcuts like Ctrl+B for bold, Ctrl+V for Italics, Marking/Un-marking checkboxes as done, increase/decrease heading levels etc.
   - It provides a single shortcut for toggling markdown preview, unlike VSCode default which maps opening and closing preview to 2 different shortcuts
   - When copy an external link, you simply need to select the target words in note and press Ctrl+V to get a properly formatted link.
   - It continues list bullets automatically and can switch between list formats easily.
2. Markdown Notes
   - Provides Wiki-links, back-links, hover peak of notes.
   - It creates non-existing notes automatically when the wiki-link is clicked.
   - In preview, wiki-links appear as normal links.
3. Markdown table formatter
   - Formats table.
   - Provide lenses for sorting the table by different columns.
   - Provides settings given below as customizable features.
     - Default justification.
     - Keep first and last pipes.
     - Space padding of columns.
     - Remove colons if same as default justification.
4. Markdown IDE
   - type `[` to trigger autocompletion for entering links to other Markdown files.
   - type `![` to trigger autocompletion for links to local image files.
   - updates links when renaming files.
   - removes links when deleting files.
5. MD-Graph
   - Faster graph showing links between files
   - ability to see graph of the current file only
6. Mzettel
   - Create a new note from given title, with zettelkasten IDs added and converted into slug
   - Copy link to current note to clipboard
7. Paste Image
   - Pastes image from clipboard as markdown links and stores the image in relative folders
8. Code-spell-checker
   - Spell checker. Worked better than `ban.spellright` extension.
9. Which-key
   - Brings Emacs hydra behavior to VSCode
10. Prettier
    - Beautifies the markdown file. Can be triggered on save.
11. Trigger task on save
    - Allows you to run tasks saved in `./.vscode/tasks.json` to be run on save. I use it to publish the current note whenever it is saved.
12. Markless
    - Enables to hide the markdown symbols used for markup. Also shows titles with larger font size according to hierarchy amd links as pretty links.
13. Markdown preview enhanced
    - Math-jax and PlantUML in preview
    - Automatic scroll sync with the editor
    - Show table of contents fixed on the side of preview
14. Themes: Mayukai, Cryptohack

## Personal extensions and modifications

1. Markdown All in One
   - added the ability to toggle checkboxes similar to the way it toggles list.
   - change the order of lists bullet point toggle. The numbered bullets appear right after the unordered lists.
2. Mzettel
   - Added the ability to open current note in a localhost server.
3. Markdown Linkify: Unpublished, developed by self
   - Converts selected text to markdown links where link part is a slug of created from selection.
   - Upon creating a link it checks for presence of said file in the current folder, and if not found it will offer to create it with title formed by selection.
4. Markdown Dated Notes: Unpublished, developed by self
   - The forward slash command of Foam plugin excised into a separate plugin. It can create wiki-links with dates as its contents

All together it has around 17 plugins. In order to ensure that VSCode does not have to bear the burden of these plugins when used for purposes other than markdown editing, I created an alias in my shell as

```zsh
alias mdcode="codium --extensions-dir ~/.local/share/mdcode/extensions --user-data-dir ~/.local/share/mdcode/usr-dir $@
```

So whenever I open any markdown folder in vscode, I open it from shell as `mdcode .`
