---
title: "Adding custom language highlight and fonts to markor"
date: 2019-12-24
categories: ["digital"]
tags: ["android"]
author: Riz
---

## Custom Language Highlight

Markor has inbuilt support for syntax highlighting of a select few languages. However markor comes bundled with [prism js](https://prismjs.com/). This means adding syntax highlighting for any of [199 languages](https://prismjs.com/#supported-languages) supported by prism is a matter of two steps
- Go to [prism CDN](https://cdnjs.com/libraries/prism) and copy the link corresponding to your language.
- In the markor app, open `Settings>View mode` and add the the following snippet to `Inject -> head` after changing the link.
```html
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.17.1/components/prism-toml.js"></script>
```

## Custom Fonts.

* Copy font to one of these locations:
  - Markor-Notebook/.app/fonts/
  - /storage/emulated/0/Fonts/
  - /sdcard/Fonts/
* In the markor app, open `Settings>Edit mode>Fonts` and select your font.