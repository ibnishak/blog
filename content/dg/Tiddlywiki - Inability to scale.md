---
title: Tiddlywiki - Inability to scale
---

# Tiddlywiki - Inability to scale

See [[Scalability Limits for Note taking Applications]]

There are [claims in the official website][1] about how much data can a single Tiddlywiki can hold. However, no actual benchmarks are given. I looked around and found the biggest public examples of Tiddlywiki I could find:

- [Kenan's wiki][2]: It has close to 30K notes, 85.6 MB in size as of 13th March 2022.
- [H0p3's wiki][3]. It has above 15K notes, 50.4 MB in size as of 13th March 2022.

I downloaded them for further testing. H0p3 seems to be managing all lists in the wiki statically[^1]. That is fine by me, I am a firm believer of not using app specific markup in the notes. However, an [audit note by the author][4] in which he talks about data loss, issues with NodeJS and other TW5 saver mechanisms and so on gives some serious points to ponder over.

Kenan's wiki loaded without any issues, but the good news ended there. Saving a note takes 5+ seconds. 5 seconds is not the end of the world, however the mild irritation that comes up with each wait is something I do not want in my note-taking workflow. Imagine each spelling mistake correction taking 5+ seconds.

The RAM usage that comes with loading and editing such a large single page application is also quite significant. Adding and editing notes in these wikis crashed my browser multiple times.

One of the most veteran users of Tiddlywiki had recently [raised the issue of scalability][5] which the community seems to have largely ignored.

Usual advice from the community is to split the data among several wikis. However, that will

- Undercut the ability to see back-links.
- Make it unable to search the text content of your notes.
- Adds link-rot upon modification of note titles.

[^1]: If you are familiar with the Tiddlywiki community, you will be hard-pressed to find a wiki where a list is created statically. Usage of [list widget][6] is the bread and butter of Tiddlywiki users. Yet here is the biggest Tiddlywiki that seems to have been extra careful in avoiding the same. 

[1]: https://tiddlywiki.com/#Scalability
[2]: https://onetwo.ren/CDDA-Kenan-Modpack-Chinese/#Index:Index%20README.md
[3]: https://philosopher.life/
[4]: https://philosopher.life/#2019.03.26%20-%20Wiki%20Audit%3A%20Cake
[5]: https://groups.google.com/g/tiddlywiki/c/dEhq2V4ZYzg/m/PZVkt5hLFAAJ
[6]: https://tiddlywiki.com/#ListWidget