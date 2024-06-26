---
title: Org mode - Tangle on file save
date: 2019-05-16T11:19:07+05:30
lastmod: 2024-04-09T11:49:55+05:30
author: Riz

description: ""
categories: []
tags: ["emacs", "org-mode"]

draft: true
enableDisqus : false
enableMathJax: false
disableToC: false
disableAutoCollapse: true
---

If you use org file to maintain your =init.el=, it is possible that many a times you need to tangle codes frequently. Below given is a couple of methods which will enable you to automatically tangle code snippets out of org mode file everytime the file is saved. 

# Via toggle function
 A toggle function is one which you can enable/disable from M-x menu (call once to enable, call again to disable). If enabled, `ri-org-toggle-tangle-on-save`  will tangle src blocks automatically everytime the org file is saved.

```elisp 
  (defun ri-tangle-on-save-org-mode-file()
    (when (equal major-mode 'org-mode)
      (org-babel-tangle)))
  (defun ri-org-toggle-tangle-on-save ()
    "Toggle tangling code blocks on save"
    (interactive)
    (if (get 'ri-org-toggle-tangle-on-save 'state)
        (progn
          (message "%s" "Org tangle on save Disabled")
          (remove-hook 'after-save-hook 'ri-tangle-on-save-org-mode-file)
          (put 'ri-org-toggle-tangle-on-save 'state nil))
      (progn
        (message "%s" "Org tangle on save Enabled")
        (add-hook 'after-save-hook 'ri-tangle-on-save-org-mode-file)
        (put 'ri-org-toggle-tangle-on-save 'state t))))

```


You can substitute ~org-babel-tangle~ part with functions other than tangling. For eg, those who use org mode to publish to html as a part of project can set it as `org-publish-current-file` to update the html file everytime you save the  org file. When you are out of the project, do not forget to toggle it to off.

# Via minor mode
An alternative is to define a minor mode that will do the same. Add the following to the =init.el= 

```elisp 
 (define-minor-mode opos-mode
      "Org-publish-on-save mode"
      :global t
      (if zo-mode
          (add-hook 'after-save-hook 'publish-on-save-org-mode-file)
        (remove-hook 'after-save-hook 'publish-on-save-org-mode-file)))

```

   You can set the opos-mode in `.dir-local-variable` in a folder to non-nil as given below, so that all org files in that folder will update the html files automatically upon file save.

```
  ((nil .  ((eval . (opos-mode 1)))))
```
