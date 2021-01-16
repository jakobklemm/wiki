---
title: Emacs hunspell setup
description: 
published: 1
date: 2021-01-16T21:08:51.193Z
tags: 
editor: markdown
dateCreated: 2021-01-16T21:08:51.193Z
---

# Linux
- Install hunspell using a package manager
- Download `.dict` and `.suff` files, store in folder
- Set `DICPATH` environment variable to the folder
(Use `.bashrc` if emacs is started using the terminal, or use `/etc/environment` if emacs is started graphically.)

# Emacs
(Replace language names)
```
(setq ispell-program-name "hunspell")

(setq ispell-local-dictionary "en_US")
(setq ispell-local-dictionary-alist
      '(("en_US" "[[:alpha:]]" "[^[:alpha:]]" "[']" nil ("-d" "en_US") nil utf-8)
        ("de_DE" "[[:alpha:]]" "[^[:alpha:]]" "[']" nil ("-d" "de_DE" "-a" "-i" "UTF-8") nil utf-8)))

(add-hook 'text-mode-hook #'flyspell-mode)
(add-hook 'org-mode-hook #'flyspell-mode)
```