cd ~/Documents/blog
outfile=$(./obsidian-transformer/obsidian-transformer "$1")
if [ $? -eq 0 ]; then
    echo OK
else
    exit 1
fi
sed -ie "/^\#\s[a-zA-z0-9]/d" ${outfile}

# hugo --minify
# ./gh-pages.js
# git add .
# git commit -m "$(date)"
# git push