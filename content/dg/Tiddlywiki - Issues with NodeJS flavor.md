---
title: Tiddlywiki - Issues with NodeJS variant
---

# Tiddlywiki - Issues with NodeJS variant

Between NodeJS flavor and Single Page Application flavor of Tiddlywiki, NodeJS flavor of Tiddlywiki is the one that acts more in line with expectations from a plain text note taking application. Each note is saved to a separate plain-text file with the extension `.tid`. One should install NodeJS runtime and acquiring basic terminal skills in order to get this running. While it was not an issue for me, this should bother average Windows users who never had to interact with terminal before.

When it comes to mobile, things start to go bad though. The existing third-party android application for Tiddlywiki does not support NodeJS flavor of Tiddlywiki. You can run [Termux][1], install NodeJS and run the server in Android. However the [loading time for Tiddlywiki server in Android is atrocious][2]. One might not prefer to have multiple NodeJS servers - one for each wiki - running in background all the time. Given that, start-up times becomes a persistent problem rather than a one-time issue.

[1]: https://termux.com/
[2]: https://ibnishak.github.io/blog/post/tw5-options-in-android-platform/#built-in-server
