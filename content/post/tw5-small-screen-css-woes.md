---
title: "TW5: CSS woes with smaller screens"
date: 2018-09-03
categories: ["digital"]
tags: ["tiddlywiki"]
author: Riz
---

We will keep this strictly to the ease o use of TW5 on smaller screens. I have been using TW5 for 4 years now, at the end of which I reached 2 conclusions.

1. It is not intended to handle large amount of data.
2. It is not well adapted for mobile platform. 

The first problem had solutions - categorize and split the data. The second problem took a long time for a satisfactory resolution. Before discussing the solutions that eventually turned up, let us indulge in a bit of rambling.

TW5 has not adapted itself well to the smaller screens. Given that it is a product supported by a community of quite a few crafty and creative gentlemen, it seems that its usage in smaller screens is not prevalent enough for it to be a pressing issue.


Let us point out the obvious one's here, and the most easily solvable ones. Anyone who has been using TW5 for any amount of time will acknowledge the importance of sidebar in the wiki. It houses the search, the tabs and page control buttons. So in a smaller screen, where will you place some of the most frequently used elements especially if the chances of long scrolls are a common occurrence? "At the top of the page" is definitely not the answer. However default TW5 themes does exactly that. Every time a new note ought to be created, you have to make the arduous journey to the top. Hey, at least my thumbs are getting stronger.


The next CSS faux pas is what makes me doubt whether anybody is using TW5 in mobile at all. When you try to access the more menu in view toolbar, the drop-down looks like this.

<img src="https://i.imgur.com/sSaufTT.jpg" title="source: imgur.com" class="center card" style="height: 500px;">

How is it that this escaped everyone?. Does nobody ever access those options at all? It could be that the order of options appearing in the drop-down menu has been memorized by everyone but me. Quite possible, I always took a little extra time to memorize the multiplication tables.

I am not touching the part where the color palettes look straight out of webzine museums or the mysterious reason there is a separate theme just to make text a little heavier. None of those concerns with small screens only and we decided to limit ourselves to that. Now for the final part of this cribbing about default CSS options has to do with how clunky an experience the actual note creation is. Between the overly large field for title, several rows of toolbar buttons and hardly touched fields like type, there is not much real estate left for actual note entry.

Admittedly all these are minor inconveniences. Community members have put in effort to create themes which will solve all these issues. However none of that has made its way to the core. Given that the desktop way of installing plugins and themes is drag 'n drop which is not available on small screens, if you want to start a wiki in smaller screens with a comfortable level of usability, you will have to jump through several hoops.

## Recommendations:
### [JD's material theme](http://j.d.material.tiddlyspot.com/)  
Using the latest [hammerjs widgets from BurningTreeC](http://hammerwidgets.tiddlyspot.com), JD has put together a modern theme for smaller screens which has behaviors expected from a webpage on this side of 2010, including swipe from edges to show the sidebar.

### [Multiling O Keyboard](https://play.google.com/store/apps/details?id=kl.ime.oh).  
Usual recommendation is to use [hacker's keyboard](https://play.google.com/store/apps/details?id=org.pocketworkstation.pckeyboard). However, multiling O keyboard is much more versatile.  

Pros
* It enables you to design your own keyboard layouts complete with Ctrl, Alt, Shift, and even Function keys. 
* User can design as many layouts as he needs and switch between them easily, essentially leaving him free to choose separate layouts fro seaparate applications. You might not need that escape key until you use vim. 

Cons
* Autocorrect and word predictions will leave you yearning for more.

## Ideas

<li> Separation of templates <br/>
Enable different filters load page-templates, view-templates, and toolbar buttons depending on screen width.<br/><br/> 
 It is not that difficult. It can even be backward compatible. Current solutions making use of css set to 'display:none' is ineffiecient as it loads the elements nonetheless.
</li><br/><br/>
<li> Fewer drop-downs <br/>
Awkwardly placed drop-down menus are rather difficult to handle in small screens. The problem is accentuated when the number of items in  the drop-down menus is high, eg; tag selection. May be tag selection in smaller screens can be done in a modal? The tag dropdown list could be activated after a minimum of 3 letters, like the search, so that number of items would be fewer?
</li><br/><br/>
<li> Better use of hammer-js.  
 There was a [plugin by danielo rodriguez](http://www.danielorodriguez.com/TW5-2click2edit) which enabled user to enter edit mode by double clicking on the view-template. This failed to work in mobile platforms when I tested, probably because it listens for clicks and not taps.<br/><br/>
Hammer widgets plugins are specifically built for mobile platforms. So I tried to replicate  the functionality using it. A view template with tap widget was created which would send "tm-edit-tiddler" message in response to 4 taps. It worked as intended, but with a major side effect. This makes the text in tiddler body unselectable from the view-template. How I wish !!.  <br/><br/> 
Nevertheless more use cases of hammer-js should be showcased. It would inspire better user interfaces hopefully.
</li>