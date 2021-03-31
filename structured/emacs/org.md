---
title: Org-mode snippets
description: Collection of useful snippets for emacs org-mode.
published: 1
date: 2021-03-31T12:07:50.467Z
tags: 
editor: markdown
dateCreated: 2021-03-08T12:55:12.587Z
---

# Images & Tables
## Caption
```
#+CAPTION: Übersicht aller durchgeführten Durchläufe mit 
```
## Position
```
#+ATTR_LATEX: :float nil
```
(float nil => Position fixed in text.)
## Borders
```
|   | A | B | C  |
| / | < | < | <> |
|   | 1 | 2 | 3  |
```
(< left, > right)

## Language
```
#+LANGUAGE: de
#+LATEX_HEADER: \usepackage[]{babel}
```