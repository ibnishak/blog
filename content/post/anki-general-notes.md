---
title: Anki - Notes on things as they are now
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
1. Image occlusion: Piotr Wozniak has backed image occlusion in the 20 rules of information retention
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
1. **[AutoEase Factor][6]**: Frequency of by button presses is as follows: Good >> Again > Hard >> Easy. Anki documents claim that when you press hard, the ease factor drops down and you see the card more often. Eventually you learn the card well enough to mark it as easy - which will bring back the ease factor up. 
   
   Given the hesitation to press "easy" button as evidenced from stats, the dropped ease factors do not recover. Anki users call this situation "ease hell". AutoEase Factor addon provides solution to this problem by raising ease factor automatically if you get a card right consecutively. It considers the entire history of a card in deciding its current ease factor. If your average ease factor is below 250, give this addon a try.
2. **[Popup dictionary][7]**: To an extend, it helps to solve the fragmentation of knowledge that comes as a side effect of minimum information principle.
3. **[Image Occlusion][8]**:

## Link to your notes.
- When you learn, write notes. Prefer markdown for ease of writing and backuo.
- Serve your notes folder over loaclhost. If notes are in markdown, you may even serve it as web pages.
- Link cards in anki to notes.
  - Use [Add hyperlink addon][5] makes adding links easy.

## Make emphasis colorful.
Changing colors of bold, italics and other empasis markers will make them more noticeable.

Eg: add the following styles to the style section of your cards (customised for night mode).

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

- For more options regarding colors, consider the addon [Customise styles][3]
- Use a [font with ligatures like fira][4]. It will convert your `->` and `==>` to proper arrows. 
  - For more monospaced fonts, see - [5 monospaced fonts][2]

## Use images
- Draw simple images, take its photo and add it to your cards as explanations. It is easier than spending time searching for the perfect image. 
- Use visual mnemonics for memorising lists. 
  - Find an image corresponding to each item in the list.
  - Combine them using a collage maker
  - Connect them with a story.

## Other
- Make MCQs. If you are preparing for an exam where MCQs are available, may be using the same in cards will reduce failure rate rapidly. 
  - This might have adverse effect on real world recalling. However, for exam preparation, it is acceptable. Use [closet addon][1] to randomize order of options appearing while previewing.
- Limit information in extra field to minimum. If you exceed 4 bullet points, you will skip reading it while previewing.
- If you have 2 or 3 categories and a long list of items belonging to each category, prepare a table. It is easier to remember which column the items are in, rather than remember a list. If you have > 3 categories, prepare 2 tables of 2 columns each.



[1]: https://ankiweb.net/shared/info/272311064
[2]: https://betterwebtype.com/articles/2020/02/13/5-monospaced-fonts-with-cool-coding-ligatures/
[3]: https://ankiweb.net/shared/info/1899278645
[4]: https://github.com/tonsky/FiraCode#whats-in-the-box
[5]: https://ankiweb.net/shared/info/318752047
[6]: https://ankiweb.net/shared/info/1672712021
[7]: https://ankiweb.net/shared/info/153625306
[8]: https://ankiweb.net/shared/info/1374772155