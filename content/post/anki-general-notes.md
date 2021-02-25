---
title: Anki - General Notes
date: 2021-02-21T06:57:13+05:30
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

## Card types
1. Image occlusion: The second coolest woz have greatly backed image occlusion in the 20 rules of information retention
2. ModifiedCloze: Title, Cloze, Extras
3. ModifiedBasic: Title, Front, Options, Back, Extras

## Justification for Title field
1. It helps to collect repeated text to a separate field and put frozen field addon to good use

```
Title: Osteomyelitis - Organism causeing the following type
Cloze:  Feathery - {{c1::Mycobacterium}}

Title: Osteomyelitis - Organism causing the following type
Cloze:  Ivory - {{c1::Syphilis}}
```
2. It ties up related cards well in card browser, at a more granular level than tags could afford.


## Recommended Addons
1. **[AutoEase Factor][6]**: Checking my stats, I realized that the frequency of by button presses is as follows: Good >> Again > Hard >> Easy. Anki documents claim that when you press hard, the ease factor drops down and you see the card more often, eventually you lean the card well enough to push the "Easy" button - which will bring back the ease factor up. For me there was marked hesitation to mark a normal card as Easy. Being that, marking a card hard card easy is practically never happening, no matter waht anki docs claim. This addon provides solution to this problem by raising ease factor automatically if you get a card right consecutively. It considers the entire history of a card in deciding its current ease factor. Check your stats and if your average ease factor is below 250, you need this addon.
2. **[Popup dictionary][7]**: To an extend, it helps to solve the fragmentation of knowledge that comes as a side effect of minimum information principle.
3. **[Image Occlusion][8]**: Because Mr.Wozniak said so.

## Link to your notes.
Create notes first, preferably as markdown files in a folder.  Then serve that folder using any one of the markdown folder servers available for free. Now you can view your notes in web browser with proper web addresses like `http:://127.0.0.1:8080/mynote`. When corresponding words appear in your cards, link to that note in the form `<a href="http:://127.0.0.1:8080/mynote">my word</a>`. Use [Add hyperlink addon][5] if you are serious about it.

## Add a splash of color to your cards
There is no law that says bold and italics has to remain the same color as the rest of the text or underline really should show an underline. Add the following styles to the style section of your cards (customised for night mode).

```css
b {
	color: #fabd2f !important;
}

u {
	text-decoration: none;
	color: #8ec07c;
	font-weight: bold
}

i {
	color: #fe8019;
	font-weight: bold;
	font-style: normal;
}

```
If you want more options regarding colors, consider the addon [Customise styles][3]

Also use a [font with ligatures like fira][4]. It will convert your `->` and `==>` to proper arrows. For more choices, see the comparison - [5 monospaced fonts][2]

## Use images
- Draw simple images, take photo and add it to your cards as explanations. Many a time it is easier than spending time looking for the perfect image in google search. 
- Use visual mnemonics as a help for memorising lists. Find an image corresponding to each item in the list, combine them using a collage maker, connect them with a story.

## Other
- Make MCQs. If you are preparing for an exam where MCQs are available, may be using the same in cards will reduce failure rate rapidly. This might have adverse effect on real world recalling, but with respect to exam prep, it is not a bad idea to make MCQs. Use [closet addon][1] to randomize order of options appearing while previewing.
- Limit information in extra field to minimum. If you exceed 4 bullet points, you will skip reading it while previewing.




[1]: https://ankiweb.net/shared/info/272311064
[2]: https://betterwebtype.com/articles/2020/02/13/5-monospaced-fonts-with-cool-coding-ligatures/
[3]: https://ankiweb.net/shared/info/1899278645
[4]: https://github.com/tonsky/FiraCode#whats-in-the-box
[5]: https://ankiweb.net/shared/info/318752047
[6]: https://ankiweb.net/shared/info/1672712021
[7]: https://ankiweb.net/shared/info/153625306
[8]: https://ankiweb.net/shared/info/1374772155