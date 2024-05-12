---
title: Tiddlywiki - Extensibility issues
---

# Tiddlywiki - Extensibility issues

Tiddlywiki has an expressive internal language in which most of plugins are written[^1]. This reminds me of Emacs/Vim attitude to extensibility. You have all the power, but only if you learn ELisp/Vim script/WikiText.

None of these languages has any application outside their restricted environments. The knowledge regarding these can be neither transferred in or transferred out. There is hardly any meta tools for this internal language - say syntax-highlighting or prettification tools. In addition, unlike Emacs/Vim, the community size of Tiddlywiki is rather small. The effect of such barriers easy to gauge when you compare the speed at which [plugin ecosystem is expanding for apps like Obsidian][1](which has well defined API in Typescript) compared to that of Tiddlywiki.

Overtime I have gained enough mastery over Tiddlywiki widgets and other quirks. However greater number of community built plugins have 2 fold effect

1. Higher chances of finding something ready-made.
2. You have more templates to learn from and build something anew and personalized.

So less tinkering, more actually writing.

[^1]: While it is also possible to write plugins in Javascript, you will be hard-pressed to find any kind of documentation regarding this.

[1]: https://obsidian-plugin-stats.vercel.app/
