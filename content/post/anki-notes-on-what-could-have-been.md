---
title: Anki - Reimagining
date: 2021-02-28T09:37:55+05:30
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

Anki SRS software's major functions can be listed as
1. Create Notes
2. Generate cards from Notes
3. Schedule cards according to spaced repetition algorithms
4. Present cards to user according to schedule
5. Provide interface for searching and reasonable organization of notes
6. Synchronise notes and scheduling information to cloud

Anki desktop application does all of these very well.
- Qt based App for input, search, organization and review.
- Notes, cards and scheduling information are stored in SQlite
- Free 1 GB cloud space for synchronisation

If I was rebuilding Anki from scratch, what would I have done?

## Notes
- Notes written in plain markdown files, and organized in folders. The program will track these folders and generate cards from them. 
  - Notes could be structured and written for the sole purpose of card generation. However it would be bettern if notes were written for general purpose and there is some markup to indicate which part of notes should be used for card generation. This would help combat knowledge fragmentation
- Use git for version control and regeneration of cards when notes are modified. The card generation is hooked to git commit (When you make a git commit/trigger card generation, git checks for modified files and generates cards from them)
- Later support can be expanded to card generation from CSV files (Frozen field type questions), YAML, XML, HTML and plain text.
  - Cloze deletion on XML based mindmaps could be used for studying for long form exams. 

## Cards
- Generated cards, models, scheduling information, revlog are stored either as individual YAML/JSON files or in a portable NoSQL database file like boltDB or badger.
  - It is better to store in both formats (YAML and NoSQL). Text will allow version control and easy backup, while databases will allow fine grained search queries, stats generation etc.
  - Sync number of due cards with calendar?
- Linked cards, which will be shown in predefined order whenever its review is scheduled. 
  - Combining this with mind-palaces would be a good way to memorise lists and procedures
- MCQ mode. A large list of options, and a set of questions with answers. During review, user will be shown the question, three randomly chosen items from the list combined with correct answer.
 

## Program
- Backend written in Golang or Rust for cross-platform compilation.
  - Golang is preferred for its faster compilation times and better availability of libraries. Majority bloggers have opined that Rust programs are faster. However remember that Anki was written in python and speed was never a complaint even then.
- Front-end is written in HTML-CSS-JS and served at a localhost port which can be viewed using any web-browser.
  - Default front-end can be replaced by anything that can ping the back-end server endpoints, icluding terminal applications.
- APIs for addon development. 
- Scheduling algorithms that can be configured and even completely replaced.
  - In order to be completely replacable, scheduling algorithms should have no side effect other than returning next due date and ease factor.
- Templating engines for creating models.
- One or more folders can be served statically for viewing related informations.
- For ease of doing backups, a 1 level folder structure could be implemented. Even that could be theoretically configured.
```
MySRSFolder
----------------
notes
    Any number and levels of subfolder containing files and media.
system
    cards
        00000000001.yaml
        00000000002.yaml
        00000000003.yaml
    metadata
        00000000001.yaml
            uuid, created, modified, easefactor, interval, model, path-to-note
            revision-log
        00000000002.yaml
        00000000003.yaml
    schedule
        2021-01-01.yaml
        2021-01-02.yaml
    models
        model1.gohtml
        model2.gohtml // Consider alternate template engines
    public
        webpages, css, javascript
```
## Advantages
- Since all notes are plain text files that can be grepped and organized in folders, there need not be separate UI for note input, search and organization
  - Especially for markdown there are several great note taking applications available including Obsidian, Foam (VS Code), Notability, Markor (Android) etc.
  - Xmind can export mindmaps as XML.
- Since git is used for version control, all git servers (Github, Gitlab, Bitbucket) become available for cloud backups. Even self hostable servers are available (Gitea, Gitlab).
- Github and similar sites enable creating notes collaboratively.
- Several other unrelated resources that can be put to use:
  - [NPM packages](https://www.npmjs.com/search?q=markdown) that extend markdown.
  - Javascript libraries for charting and thus. 
  - Browser addons
- Compare the efficacy of different algorithms (SM2, ebisu, SM-15 etc)
  
## Note
Several, if not all, of these advantages can be achieved in anki. However there are several hoops to jump through, and some are not documented. I also believe that such monolithic applications are restrictive. I understand the convienience they provide to the end user as opposed to the modular approach I described above. However when you build an application, it should have maximum possible openness to workflows and usage habits. In course of evolution of an application,  monolithic branches may appear due to the very nature of its openness, but never as its main main stem. 