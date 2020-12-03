---
title: Jakku
description: Browser based emacs task viewer.
published: 1
date: 2020-12-03T21:02:59.916Z
tags: 
editor: markdown
dateCreated: 2020-12-03T21:02:59.916Z
---

Easy online viewing and editing of Emacs org-mode TODOs, by Jakob Klemm.
## Setup
** Architecture
Owncloud + Emacs
### Variables
- letter, string.
  Letter used in =org-agenda= selection menu for the agenda, default is =a=. Special
  letters allow special filters to be applied within Emacs.
- init, string.
  Path to init file for Emacs config, default is ~/.emacs.d/init.el
## Extraction
Get all TODO's with the following command, formatted as CSV.
``` sh
emacs -batch -l ~/.emacs.d/init.el -eval '(org-batch-agenda-csv "a" org-agenda-span (quote month))'
```
The following fields exist:
``` csv
category,head,type,todo,tags,date,time,extra,priority-l,priority-n
```
https://orgmode.org/manual/Extracting-Agenda-Information.html
The shell script fetches that information and writes it to =current.csv=. From
there the data is parsed. =extract.sh= can simple be run by a CRON-Job.