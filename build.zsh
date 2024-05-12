#!/usr/bin/env zsh

hugo --minify # builds static pages to folder hugopublish
gatsby build # builds static pages to folder public

# ./gh-pages.js