---
title: "TW5: Options in Android Platform"
date: 2018-08-14
categories: ["digital"]
tags: ["tiddlywiki"]
author: Riz
---

Saving a TW5 wiki in android has 7 possible solutions. We shall take a quick look at each of them.

- [Android Apps](#android-apps)
  - [Tiddloid](#tiddloid)
  - [Andtidwiki](#andtidwiki)
  - [Webdav server](#webdav-server)
- [Servers using Termux](#servers-using-termux)
  - [Built in server.](#built-in-server)
  - [Tiddlyserver.](#tiddlyserver)
  - [Bob](#bob)
  - [And finally the winner: Widdly by Opennota](#and-finally-the-winner-widdly-by-opennota)


<br/><br/>
### Android Apps

#### Tiddloid
Tiddloid is the latest android app to save tiddlywiki in android. It is surprisingly good. It allows you to add TW5s stored anywhere in your internal memory, saves without glitches and even has dark mode. If you want to use TW5 in android, tiddloid is definitely the way to go. Developer is actively adding new features. Download the [apk from github](https://github.com/donmor/Tiddloid/releases).

#### Andtidwiki
 The first of them all. When I tested it 2 years ago it was severely limited in options, and had several bugs. However it was the only option available back then. When I was writing this, I went back to see why I stopped using Andtidwiki again. Unfortunately the free ad supported version of Andtidwiki is no longer available. Links pointing towards Andtidwiki from the official documentation is also dead.  I did not feel like giving money to product just to recall why I should not use it. In addition, last update to the app happened in 2013. If you are eager, go ahead and have a look [here](https://play.google.com/store/apps/details?id=de.mgsimon.android.andtidwikiplus).

#### Webdav server
I would consider this as the simplest and most beginner friendly way to save TW5. There are multiple applications in the google playstore that can launch a webdav server. I used the webdav-server listed [here](https://play.google.com/store/apps/details?id=com.theolivetree.webdavserver). The steps to get going could not have been simpler.

* Install the app and open it
* Go to Settings
* Network Interface > Loopback
* Set the port number you want
* Set the home directory to "Custom folder" and use the custom folder option to point it to a folder in your phone memory.
* Rest of the settings are optional.
* Now move your HTML TW5s to the folder you set in the app. Start the server from the app and open 127.0.0.1:8080 (or whatever port number you set) in your phone browser. Your TWs will be served. 

 This is a reasonably good solution. Unfortunately you are limited to using HTML TW5s, as it cannot serve a TW5 created using node server. Personally for me, this was a deal breaker as whatever TW5 I use, I need them in a format that lends itself to searching by external tools. 

### Servers using Termux

The rest of the options depend on [termux](https://f-droid.org/en/packages/com.termux/).

#### Built in server. 
It is a bare-bones solution which does what it is built for well. It can serve plain text wikis but not the HTML wikis. You can keep the wikis anywhere in your internal memory, as that is what is accessible to termux in a non rooted phone. Launching it is a matter of changing into your wiki directory and launching the server.

The glaring drawback of this is the start up time. In my android phone with 2GB RAM, it takes 20-30 seconds to load an empty wiki. With reasonable amount of data, it would go north of 80-90 secs. Unless the plan is to keep a termux instance dedicated to a wiki and keep it running all the time, this is not a viable solution to run a TW5 in a mobile phone.
![Imgur](https://i.imgur.com/xUwpmq1.jpg)

#### Tiddlyserver.
[Tiddlyserver by Arlen Beiler](https://github.com/Arlen22/TiddlyServer) can serve both node based wikis and HTML wikis. It requires to be configured beforehand, which throws a little spanner in the plans of syncing between different OS, but that is not the major issue. The major issue is the start up time once again. Tiddlyserver startup times are not much different from the built in server start up times. 

#### Bob
[Bob by Jed Carty](https://github.com/OokTech/TW5-Bob) is the new kid in the block. Its adds a lot of fancy options to the default server. The one that is most relevant in my use case was that it can reflect the changes made to the tid files in the browser in real time unlike the built in server which requires restart. This is a huge advantage for small screens.  Any minimal text editing android app can be used to edit the tid files, and you can see the results in the browser. You do not have to wade through the edit template anymore. There was a few hiccups here and there while using from termux, but the direction Jed has taken and his persistence with the project is definitely laudable. I think this would be the closest we can hope for short of a TW5 android application.
 
Nevertheless, even Bob is plagued by same start up timings that affect the other two.

#### And finally the winner: Widdly by Opennota
Boy who would have thought. [Widdly by opennota](https://gitlab.com/opennota/widdly) is one of the least known solutions for TW5. Despite being the most [starred](https://github.com/opennota/widdly) solution for serving TW5s, you cannot find a single mention about widdly in TW5 google groups.

Widdly by opennota is a minimal server written in golang for TW5. What is the biggest advantage? Well it started up and served my reasonably large wiki in 1.2 seconds. Let that sink in, and compare it with the node servers loading an empty wiki in 30 seconds. Now that is impressive, so what is the catch? The catch is that widdly saves your tiddlers neither as a HTML file, nor as tid files, but in a bolt-db database. This could make a few foreheads frown but hear this, the database automatically saves every revision you made to every tiddler. In addition, once compiled it can act as a portable server for that OS, running right from your pendrive. That nailed its position firmly at the top of my list fulfilling every criteria I set.

However it still sacrifices the ability to edit and reflect the changes made to tid files from an editor. Turns out, widdly had no intention to cast any doubts as to who should be the winner. The solution came as a surprise while browsing aimlessly through github. There is a [fork of widdly by Aaron Hutchinson](https://github.com/xarnze/widdly), in which he implemented both sqlite and flatfile storages. Better, you can choose any extension for the flatfile backend, and since the metadata is stored in a separate file, you can safely use anything from markdown editor android apps to notepad to edit your tiddlers and changes can be seen in browser with a simple refresh. 

Opennota recently moved the project to gitlab, indicating that he has not abandoned the project and hopefully, might merge  Aaron's mods. It requires documentation, and probably splitting different backends to separate branches so that users can compile by simply switching to the branch of desired backend. 
