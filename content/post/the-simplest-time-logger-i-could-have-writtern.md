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
#!/usr/bin/env zsh

##########  VARIABLES  ################
AREAS=$(echo 'Acad!Programming!Reading!Calories')

#######################################

data=$(yad --form --columns=2 --date-format="%Y-%m-%d"  \
--field="Date::DT" "$(date +%Y-%m-%d)" \
--field="Area::CBE" "${AREAS}" \
--field="Units::NUM" "25") \

dt=$(echo $data | awk -F'|' '{print $1}')
ar=$(echo $data | awk -F'|' '{print $2}')
un=$(echo $data | awk -F'|' '{print $3}')

echo $dt $ar $un

sqlite3 tracker.db <<END_SQL
CREATE TABLE IF NOT EXISTS data(
   ID INTEGER PRIMARY KEY AUTOINCREMENT,
   DATE           TEXT    NOT NULL,
   AREA           TEXT    NOT NULL,
   UNITS        	INT 		NOT NULL
);
INSERT INTO data (DATE,AREA,UNITS)
VALUES ( "$dt", "$ar", "$un" );
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