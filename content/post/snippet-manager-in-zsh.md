---
title: Snippet Manager in Zsh
date: 2021-07-21T01:13:42+05:30
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

1. Save it as `snippets.zsh` in your `$PATH`
2. Make it executable

Usage
1. To create a new golang snippet
```
snippets.zsh -d go -n
```

2. To create a new zsh snippet
```
snippets.zsh -d zsh -n
```

3. To copy a golang snippet to clipboard
```
snippets.zsh -d go -l
```

```zsh
#!/usr/bin/env zsh

##------------USER VARIABLES---------------##
SNIPDIR=/path/to/snippet/dir
EDITOR=(leafpad) # Do not remove the brackets
##-----------------------------------------##

slugify() {
echo $1 |                                                                                                                              
        iconv -t ascii//TRANSLIT |                                                                                
        sed -E 's/[~\^]+//g' |                                                         
        sed -E 's/[^a-zA-Z0-9]+/-/g' |                          
        sed -E 's/^-+\|-+$//g' | 
        sed -E 's/^-+//g' | 
        sed -E 's/-+$//g' | 
        tr A-Z a-z
}


new() {
	TMP=$(mktemp /tmp/quick-note.XXXXXX)
	${EDITOR} ${TMP} 
	TITLE=$(head -n 1 ${TMP})
	fn=$(slugify "${TITLE}").${LANG}
	tail -n +2 ${TMP} > $fn
	echo $TITLE >> index
	rm ${TMP}
}

launch() {
	TITLE=$(cat index | fzf)
	fn=$(slugify "${TITLE}").${LANG}
	cat ${fn} | xsel --clipboard 
	if [ $? -eq 0 ]; then
    notify-send -i "notification-message-IM" 'Snippet Copied!' "$TITLE"
	else
			notify-send -i "notification-message-IM" 'ERROR' "Unknown error"
	fi
}


while getopts 'nld:' z
do
  case $z in
    n) ACTION=NEW ;;
    l) ACTION=LAUNCH ;;
    d) LANG=$OPTARG ;;
  esac
done

pushd ${SNIPDIR}

    if [ -z ${LANG} ]
    then
          notify-send -i "notification-message-IM" 'ERROR' "Language is empty"
          exit 1
    fi

    if [ ! -d ${LANG} ] 
    then
        mkdir ${LANG}
    fi

    pushd ${LANG}
    case $ACTION in
        NEW)    new    ;;
        LAUNCH) launch ;;
      esac
    popd

popd
```