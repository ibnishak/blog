---
title: "zsh: Parameter expansion cheatsheet"
date: 2018-09-26
categories: ["digital"]
tags: ["tip", "linux"]
author: Riz
description: "String and path manipulation using zsh alone"
---

## Parent directory
```bash
echo $PWD
# /home/learning/pf/zsh/zsh2.00/src
echo $PWD:h
# /home/learning/pf/zsh/zsh2.00
```
## Grandparent directry
```bash
echo $PWD:h:h
# /home/learning/pf/zsh
```
## Current dir without path
```bash
echo $PWD:t
# src
```
## Extracting file name and extension
```bash
name=foo.c
echo $name:r
# foo
echo $name:e
# c
```

## Substituting parts of variable
```bash
name=foo.c
echo $name:s/foo/bar/ 
# bar.c

```

Mass renaming of files
```bash
ls
# foo.c    foo.h    foo.o    foo.pro
for i in foo.*; mv $i $i:s/foo/bar/
ls
# bar.c    bar.h    bar.o    bar.pro

```

## Uppercasing and Lower casing
```bash
xx=Test
echo $xx:u
# TEST
echo $xx:l
# test
```

## Cutting patterns from beginning and end

- Cut from end
```bash
echo $PWD
# /home/learning/pf/zsh/zsh2.00/src
echo ${PWD%/*}
# /home/learning/pf/zsh/zsh2.00
echo ${PWD%/*/*}
#/home/learning/pf/zsh

name=foo.c
echo ${name%.*}
# foo
```
- Cut from beginning
```bash

echo ${PWD##*/}
# src

name=foo.c
echo ${name#*.}
# c
```


