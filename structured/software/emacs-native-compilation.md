---
title: Emacs Native Compilation
description: 
published: true
date: 2021-04-30T21:10:05.404Z
tags: 
editor: markdown
dateCreated: 2021-04-30T21:10:05.404Z
---

```
yay -S libgccjit
git clone https://github.com/emacs-mirror/emacs
cd emacs
./autogen
./configure --with-native-compilation --with-mailutils --with-sound=yes --with-xwidgets CFLAGS='-O3 -march=native'
make NATIVE_FULL_AOT=1 -j$(nproc)
```