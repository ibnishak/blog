cd ~/Documents/blog
./obsidian-transformer/obsidian-transformer "$1" > content/post/01.md
hugo --minify
./gh-pages.js
git add .
git commit -m "$(date)"
git push