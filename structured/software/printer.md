---
title: Arch Linux HP-Printer Setup
description: 
published: 1
date: 2021-01-04T13:58:06.535Z
tags: 
editor: markdown
dateCreated: 2021-01-04T13:58:06.535Z
---

1. Install hplip & cups using pacman
2. `sudo systemctl enable cups`
3. `sudo hp-setup -i`, go through the setup.
4. Install system-printer-config using pacman
5. Execute system-printer-config as root, it requires permissions.