---
title: The Simplest Time Logger I Could Have Written
date: 2021-07-09T20:37:38+05:30
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

A zsh script to take input from a form, and store it in a sqlite database.

Dependecies: YAD, Sqlite

```zsh
##########  VARIABLES  ################
AREAS=(Acad Programming Reading Calories) #Single words separated by spaces

#######################################

JA=$(echo ${(j:\!:)AREAS}) # Join Areas with ! because that is how yad wants it.

data=$(yad --form --columns=2 --date-format="%Y-%m-%d"  \
--field="Date::DT" "$(date +%Y-%m-%d)" \
--field="Area::CBE" "${JA}" \
--field="Units::NUM" "25")

DTARR=("${(@s/|/)data}") # Split yad output to array

if [ -z "$DTARR[1]" ] # If user had cancelled the yad window, exit
then
   echo "User cancellation"
   exit 1
fi

sqlite3 tracker.db <<END_SQL
CREATE TABLE IF NOT EXISTS data(
   ID INTEGER PRIMARY KEY AUTOINCREMENT,
   DATE           TEXT    NOT NULL,
   AREA           TEXT    NOT NULL,
   UNITS        	INT 		NOT NULL
);
INSERT INTO data (DATE,AREA,UNITS)
VALUES ( "$DTARR[1]", "$DTARR[2]", "$DTARR[3]" );
END_SQL
```

Create a desktop launcher for it by saving the following file as `Tracker.desktop` in your desktop. Remember to substitute paths appropriately.

```
[Desktop Entry]
Type=Application
Icon=/path/to/any/icon
Name=Tracker
Exec=/path/to/the/script
Path=/dir/where/database/should/be/kept
Terminal=false
Hidden=false
```

Next stop - creating graphs from the tracker database.