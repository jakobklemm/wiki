---
title: Microsoft Teams
description: Issues and debugging for Microsoft Teams.
published: 1
date: 2021-01-03T19:54:42.672Z
tags: 
editor: markdown
dateCreated: 2021-01-03T19:54:42.672Z
---

# Login Issues
If all Microsoft Products are available and useable but the login at https://teams.microsoft.com fails after being stuck in a reconnect loop it might be an issue with the local system time. 

If it is just a timezone issue or a local config error doesn't matter but in order to solve this issue the local time needs to be set to (more or less) UTC.

(It is irelevant whether the installed local app or the browser version is used, the local systemtime is used either way.)
