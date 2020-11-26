---
title: "Terminal theme: Changes"
date: 2019-10-18
draft: true
tags: ["misc"]
author: Riz
---

In themes/terminal/css/main.css the class affecting the blogs viewing area is defined as `.container`. I added `margin-left: auto;` to the container class and changed `border-right `to `border-left`. To utilize a little more space than default, I slighlty increased the blog area to 100px from 840px. My container class looks like this

```css
.container {
  display: flex;
  flex-direction: column;
  padding: 40px;
  max-width: 1000px;
  min-height: 100vh;
  border-left: 1px solid rgba(255, 255, 255, 0.1);
  margin-left: auto;

  &.full,
  &.center {
    border: none;
    margin: 0 auto;
  }
```
I also added a class for kbd as

```css
kbd {
  color: #ffe1cc
}
```