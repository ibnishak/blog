---
title: Installation of PostgreSQL Database in Arch Linux
date: 2024-02-15T16:47:28+05:30
lastmod: 2024-02-18T08:00:04+05:30
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



1. Install the package `postgresql`
```
pacman -S postgresql
```

2. Installation also adds a new user called `postgres` to the system. Ensure this by running the following command
```
cat /etc/passwd
```

3. Set password for the user `postgres`
```
sudo passwd postgres
```
4. Login to `postgres` user
```
su postgres
```

5. Initiate DB
```
initdb -D /var/lib/postgres/data --data-checksums
```
You CANNOT use a sub-directory under your $USER home directory as postgres directory afaik.

6. Start PostgreSQL service by running
```
systemctl start postgresql.service
```