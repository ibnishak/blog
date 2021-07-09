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

```zsh
#!/usr/bin/env zsh

##########  VARIABLES  ################
AREAS=$(echo 'Acad!Programming!Reading!Calories')
DIR=/home/richie/i/04-compass

#######################################

cd $DIR

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