---
title: Alias for updating YADM
date: 2020-12-07T21:15:38+05:30
lastmod: 
author: Riz

description: 
categories: ["digital"]
tags: ["linux"]

draft: false
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

## Yadm update alias

Ref: [YADM][3]

Open `~/.config/yadm/repo.git/config` and add the following

```
[alias]
  update = !git add -u && git commit -m \"$(date)\" && git push
```

Now when you give the command `yadm update`, yadm will update all the exisiting files in repo, commit it with date and try to push to orgin. For bonus, you can add ssh keys to github/[gitlab][2] accounts and set this on a [cron job][1] to completely automate backing up of your configuration. 

[1]: https://crontab.guru/#0_*/3_*_*_*
[2]: https://docs.gitlab.com/ee/gitlab-basics/create-your-ssh-keys.html
[3]: https://yadm.io/