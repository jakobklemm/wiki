---
author: Jakob Klemm
db: true
editor: markdown
email: jakob.klemm\@protonmail.com
title: A clean emacs config is an empty config file.
---

```{=org}
#+HUGO_SECTION: posts
```
```{=org}
#+HUGO_BASE_DIR:~/documents/projects/knowledge
```
Emacs Config - Jakob Klemm

Setup
=====

`use-package`
-------------

Verify the `use-package` installation and ensures the packages.

``` {.commonlisp org-language="emacs-lisp"}
(unless (package-installed-p 'use-package)
(package-refresh-contents)
(package-install 'use-package))

(require 'use-package)
(setq use-package-always-ensure t)
```

Update and compile
------------------

Update and compile all packages.

``` {.commonlisp org-language="emacs-lisp"}
(use-package auto-compile
  :config (auto-compile-on-load-mode))
(setq load-prefer-newer t)
```

Defaults
--------

Default settings cloned from [Harry R.
Schwartz](https://github.com/hrs/sensible-defaults.el). Functions:

-   Ensuring that files end with newlines,
-   Always enabling syntax highlighting,
-   Increasing the garbage collection threshold,
-   Defaulting line-length to 80 characters,
-   Creating parent directories after saving a deeply nested file.

``` {.commonlisp org-language="emacs-lisp"}
(load-file "~/.emacs.d/sensible-defaults.el/sensible-defaults.el")
(sensible-defaults/use-all-settings)
(sensible-defaults/use-all-keybindings)
;; Disable all backups.
(setq make-backup-files nil)
```

Resources
---------

Add `resources` to the path

``` {.commonlisp org-language="emacs-lisp"}
(add-to-list 'load-path "~/.emacs.d/resources/")
```

Interface
=========

Disable
-------

### Scroll bar

Disable the scroll bar in the entire window and mini buffers.

``` {.commonlisp org-language="emacs-lisp"}
;; Menu bar
(tool-bar-mode 0)
(menu-bar-mode 0)
(scroll-bar-mode -1)
;; Minibuffer
(set-window-scroll-bars (minibuffer-window) nil nil)
```

### Minor modes

Hide all minor modes with `minions`.

``` {.commonlisp org-language="emacs-lisp"}
(use-package minions
:config
(setq minions-mode-line-lighter ""
minions-mode-line-delimiters '("" . ""))
(minions-mode 1))
```

### Scrolling

Don\'t skip to center of page when at bottom / top, *normal* smooth
scrolling.

``` {.commonlisp org-language="emacs-lisp"}
(setq scroll-conservatively 100)
```

Symbols
-------

Prettify symobls.

``` {.commonlisp org-language="emacs-lisp"}
(global-prettify-symbols-mode 1)
(defun add-pretty-lambda ()
"Make some word or string show as pretty Unicode symbols.  See https://unicodelookup.com for more."
(setq prettify-symbols-alist
      '(
        ("lambda" . 955)
        ("->" . 8594)
        ("<-" . 8592)
        ("<=" . 8804)
        (">=" . 8805)
        ("=~" . 8771)
        ("!=" . 8800)
        (":=" . 8788)
        )))
(add-hook 'prog-mode-hook 'add-pretty-lambda)
(add-hook 'org-mode-hook 'add-pretty-lambda)
```

Mod line
--------

Doom mod line.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (use-package doom-modeline
    :ensure t
    :init (doom-modeline-mode 1)))
```

Config and settings for mod line, from
[doom-modline](https:github.com/seagle0128/doom-modline)

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (setq doom-modeline-icon (display-graphic-p))
  (setq doom-modeline-major-mode-icon t)
  (setq doom-modeline-buffer-state-icon t)
  (setq doom-modeline-buffer-modification-icon t)
  (setq doom-modeline-unicode-fallback nil)
  (setq doom-modeline-minor-modes nil)
  (setq doom-modeline-enable-word-count nil)
  (setq doom-modeline-buffer-encoding t)
  (setq doom-modeline-lsp t)

  ;; Whether display the environment version.
  (setq doom-modeline-env-version t)
  ;; Or for individual languages
  (setq doom-modeline-env-enable-python t)
  (setq doom-modeline-env-enable-ruby t)
  (setq doom-modeline-env-enable-perl t)
  (setq doom-modeline-env-enable-go t)
  (setq doom-modeline-env-enable-elixir t)
  (setq doom-modeline-env-enable-rust t)
  (display-time-mode 1))
```

Fullscreen
----------

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
   (add-to-list 'default-frame-alist '(fullscreen . maximized)))
```

Current line
------------

Hightlight the current line.

``` {.commonlisp org-language="emacs-lisp"}
(global-hl-line-mode)
    ;; Marked number
    (set-face-background hl-line-face "#090405")
```

Line numbers
------------

``` {.commonlisp org-language="emacs-lisp"}
;; Toggle with =C-c j=
     ;;(global-display-line-numbers-mode)
```

Kill and close
--------------

Kill the current buffer and close the window in one command.

``` {.commonlisp org-language="emacs-lisp"}
(global-set-key (kbd "C-x j") 'kill-buffer-and-window)
```

Cursor
------

Set the cursor color to the same as beacon.

``` {.commonlisp org-language="emacs-lisp"}
(set-cursor-color "#c678dd")
```

Images
------

Display images inline. Toggle with `C-c C-x C-v`

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (add-hook 'org-mode-hook 'org-toggle-inline-images)
  (setq org-image-actual-width '(600)))
```

Theme
-----

Used themes:

-   Elixify - AstonJ (elixirforum)
-   Doom-nord
-   Doom-material
-   Doom-spacegrey

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (use-package doom-themes)
  (load-theme `doom-nord t))
  ;; (load-theme `doom-material t)
  ;; (load-theme `doom-spacegrey t)
```

Dashboard
---------

Setup the dashboard with come modifications and configs.

### Dependancies

Page-break-lines

``` {.commonlisp org-language="emacs-lisp"}
(use-package page-break-lines)
(turn-on-page-break-lines-mode)
```

Install icons. Not only used by `dashboard` but its the main dependancy.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
     (use-package all-the-icons))
```

### Setup

Setup the dashboard.

``` {.commonlisp org-language="emacs-lisp"}
;; Disable *scratch*
(when (display-graphic-p)
  (setq inhibit-startup-screen t
        initial-buffer-choice  nil)

  (use-package dashboard
    :ensure t
    :config
    (dashboard-setup-startup-hook)))
```

### Config

Options and configuration for dashboard following the readme.

``` {.commonlisp org-language="emacs-lisp"}
;; Set the banner
(setq dashboard-startup-banner 2)
;; Content is not centered by default. To center, set
(setq dashboard-center-content t)
;; Icons
(setq dashboard-set-heading-icons t)
(setq dashboard-set-file-icons t)
;; Navigator
(setq dashboard-set-navigator t)
;; Init info
(setq dashboard-set-init-info t)
;; Message
(setq dashboard-footer-messages '("Every time I see this package I think to myself \"People exit Emacs?\""))

(setq dashboard-items '((recents  . 10)
                        (agenda . 15)))
```

Font
----

Use Fira Code as default font.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
       (set-face-attribute
        'default nil
        :font "Fira Code"
        :weight 'normal
        :width 'normal
        ))
```

New window
----------

Directly switch to new window after opening. (Credit: hrs)

``` {.commonlisp org-language="emacs-lisp"}
(defun hrs/split-window-below-and-switch ()
"Split the window horizontally, then switch to the new pane."
(interactive)
(split-window-below)
(balance-windows)
(other-window 1))

(defun hrs/split-window-right-and-switch ()
"Split the window vertically, then switch to the new pane."
(interactive)
(split-window-right)
(balance-windows)
(other-window 1))

;; Keys
(global-set-key (kbd "C-x 2") 'hrs/split-window-below-and-switch)
(global-set-key (kbd "C-x 3") 'hrs/split-window-right-and-switch)
```

Ace
---

Use `ace-windows`, mainly as a dependency for `org-roam`

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (global-set-key (kbd "M-o") 'ace-window))
```

Beacon
------

Beacon for highlighting the cursor when switching buffers.

``` {.commonlisp org-language="emacs-lisp"}
(use-package beacon
  :custom
  (beacon-color "#c678dd")
  :hook (after-init . beacon-mode))
```

Title
-----

Set the window title to the current file.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (setq-default frame-title-format
                '(:eval
                  (format "%s@%s: %s %s"
                          (or (file-remote-p default-directory 'user)
                              user-real-login-name)
                          (or (file-remote-p default-directory 'host)
                              system-name)
                          (buffer-name)
                          (cond
                           (buffer-file-truename
                            (concat "(" buffer-file-truename ")"))
                           (dired-directory
                            (concat "{" dired-directory "}"))
                           (t
                            "[no file]"))))))
```

Resize
------

Easy zoom in & out.

``` {.commonlisp org-language="emacs-lisp"}
(defun zoom-in ()
  (interactive)
  (let ((x (+ (face-attribute 'default :height)
              10)))
    (set-face-attribute 'default nil :height x)))

(defun zoom-out ()
  (interactive)
  (let ((x (- (face-attribute 'default :height)
              10)))
    (set-face-attribute 'default nil :height x)))

(define-key global-map (kbd "C-1") 'zoom-in)
(define-key global-map (kbd "C-0") 'zoom-out)
```

Move
----

Move between multiple open windows.

``` {.commonlisp org-language="emacs-lisp"}
(when (fboundp 'windmove-default-keybindings)
(windmove-default-keybindings))
```

Previous buffer
---------------

Quickly switch to the previous buffer, useful for `org-agenda`

``` {.commonlisp org-language="emacs-lisp"}
(define-key global-map (kbd "C-x x") 'previous-buffer)
```

Projects
========

Editing
-------

### Indentation

Copied from [.emacs.d](https://github.com/MatthewZMD/.emacs.d)

``` {.commonlisp org-language="emacs-lisp"}
(use-package highlight-indent-guides
:if (display-graphic-p)
:diminish
;; Enable manually if needed, it a severe bug which potentially core-dumps Emacs
;; https://github.com/DarthFennec/highlight-indent-guides/issues/76
:commands (highlight-indent-guides-mode)
:custom
(highlight-indent-guides-method 'character)
(highlight-indent-guides-responsive 'top)
(highlight-indent-guides-delay 0)
(highlight-indent-guides-auto-character-face-perc 7))
```

Indent config

``` {.commonlisp org-language="emacs-lisp"}
(setq-default indent-tabs-mode nil)
(setq-default indent-line-function 'insert-tab)
(setq-default tab-width 4)
(setq-default c-basic-offset 4)
(setq-default js-switch-indent-offset 4)
(c-set-offset 'comment-intro 0)
(c-set-offset 'innamespace 0)
(c-set-offset 'case-label '+)
(c-set-offset 'access-label 0)
(c-set-offset (quote cpp-macro) 0 nil)
(defun smart-electric-indent-mode ()
"Disable 'electric-indent-mode in certain buffers and enable otherwise."
(cond ((and (eq electric-indent-mode t)
(member major-mode '(erc-mode text-mode)))
(electric-indent-mode 0))
((eq electric-indent-mode nil) (electric-indent-mode 1))))
(add-hook 'post-command-hook #'smart-electric-indent-mode)
```

### CamelCase

Treat camel casing (the best and only right variable naming system) as
multiple words.

``` {.commonlisp org-language="emacs-lisp"}
(use-package subword
:config (global-subword-mode 1))
```

### UTF-8

Treat every file as UTF-8 by default.

``` {.commonlisp org-language="emacs-lisp"}
(set-language-environment "UTF-8")
```

### Wrap

Auto wrap paragraphs. Or use `M-q`.

``` {.commonlisp org-language="emacs-lisp"}
(add-hook 'text-mode-hook 'auto-fill-mode)
(add-hook 'gfm-mode-hook 'auto-fill-mode)
(add-hook 'org-mode-hook 'auto-fill-mode)
```

### Spacing

Cycle spacing options.

``` {.commonlisp org-language="emacs-lisp"}
(global-set-key (kbd "M-SPC") 'cycle-spacing)
```

### Modes

Other *cool* default modes.

``` {.commonlisp org-language="emacs-lisp"}
(show-paren-mode 1)
(column-number-mode 1)
(size-indication-mode 1)
(transient-mark-mode 1)
(delete-selection-mode 1)
```

### Undo tree

Visual undo tree

``` {.commonlisp org-language="emacs-lisp"}
(use-package undo-tree
:defer t
  :diminish undo-tree-mode
  ;;:init (global-undo-tree-mode)
  :custom
  (undo-tree-visualizer-diff t)
  (undo-tree-visualizer-timestamps t))
  (global-set-key (kbd "C-x C-u") 'global-undo-tree-mode)
```

Helper
------

### Kill current

Kill the current buffer instead of asking.

``` {.commonlisp org-language="emacs-lisp"}
    (defun kill-current-buffer ()
(interactive)
(kill-buffer (current-buffer)))

    ;; Keybind
    (global-set-key (kbd "C-x k") 'kill-current-buffer)
```

### Save

Save the location within a file.

``` {.commonlisp org-language="emacs-lisp"}
(save-place-mode t)
```

### Which key

Helpful with long keybinds.

``` {.commonlisp org-language="emacs-lisp"}
(use-package which-key
:config (which-key-mode))
```

### Jump

Jump to function definitions. (Works with elixir)

``` {.commonlisp org-language="emacs-lisp"}
(use-package dumb-jump
:ensure t
:bind (("M-g o" . dumb-jump-go-other-window)
("M-g j" . dumb-jump-go))
:config (setq dumb-jump-selector 'ivy))
```

### Refresh

Auto refresh updated files to avoid overwriting changes.

``` {.emasc-lisp}
(global-auto-revert-mode t)
```

Correction
----------

### Spell check

Enable spellcheck for both English and German in all `org-mode` and
`text-mode` buffers. Select the current spellcheck with
`ispell-change-directory`, then use `C-.` to see suggestions and see
issues with `flyspell-buffer`.

``` {.commonlisp org-language="emacs-lisp"}
(setq ispell-program-name "hunspell")
(setq ispell-hunspell-dict-paths-alist
      '(("en_US" "~/.emacs.d/dict/en_US.aff")
        ("de_DE" "~/.emacs.d/dict/de_DE.aff")))

(setq ispell-local-dictionary "en_US")
(setq ispell-local-dictionary-alist
      '(("en_US" "[[:alpha:]]" "[^[:alpha:]]" "[']" nil ("-d" "en_US") nil utf-8)
        ("de_DE" "[[:alpha:]]" "[^[:alpha:]]" "[']" nil ("-d" "de_DE" "-a" "-i" "UTF-8") nil utf-8)))

(flyspell-mode 1)
(add-hook 'text-mode-hook #'flyspell-mode)
(add-hook 'org-mode-hook #'flyspell-mode)

(global-set-key (kbd "C-.") 'ispell-word)
```

### Completion

Use package `company` as a dependency of lsp-mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package company)
(add-hook 'after-init-hook 'global-company-mode)
(use-package lsp-mode
:commands lsp
:ensure t
:diminish lsp-mode
:hook
(elixir-mode . lsp)
:init
(add-to-list 'exec-path "~/.emacs.d/elixir-ls"))
```

Fly check mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package flycheck)
(global-flycheck-mode)
```

Configure `lsp-mode`

``` {.commonlisp org-language="emacs-lisp"}
(use-package lsp-ui :commands lsp-ui-mode)
(use-package lsp-ivy :commands lsp-ivy-workspace-symbol)
```

Shell
-----

Bind `C-x t` to `eshell`.

``` {.commonlisp org-language="emacs-lisp"}
(global-set-key (kbd "C-x t") 'eshell)
```

Ivy - Swiper
------------

Use Ivy and Swiper over Helm.

``` {.commonlisp org-language="emacs-lisp"}
(use-package swiper)
(use-package ivy)
(ivy-mode 1)
(setq ivy-use-virtual-buffers t)
(setq enable-recursive-minibuffers t)
(setq search-default-mode #'char-fold-to-regexp)
(global-set-key "\C-s" 'swiper)

(use-package amx
  :ensure t
  :after ivy
  :custom
  (amx-backend 'auto)
  (amx-save-file "~/.emacs.d/amx-items")
  (amx-history-length 50)
  (amx-show-key-bindings nil)
  :config (amx-mode 1))

(use-package all-the-icons-ivy-rich
  :ensure t
  :init (all-the-icons-ivy-rich-mode 1))

(use-package ivy-rich
  :ensure t
  :after ivy
  :custom (setcdr (assq t ivy-format-functions-alist) #'ivy-format-function-line)
  :config (ivy-rich-mode 1))
```

Posframe
--------

Use posframe for Ivy & Swiper.

``` {.commonlisp org-language="emacs-lisp"}
(use-package ivy-posframe
  :after ivy
  :init (ivy-posframe-mode 1)
  :config (setq ivy-posframe-display-functions-alist '((t . ivy-posframe-display-at-frame-center))
        ivy-posframe-border-width 8))

(setq ivy-posframe-parameters
      '((left-fringe . 4)
        (right-fringe . 4)))
(setq ivy-posframe-border-width 2)
```

Snippets
--------

Use yasnippets and the snippets from github.com/hrs/dotfiles

``` {.commonlisp org-language="emacs-lisp"}
(use-package yasnippet)
(setq yas-snippet-dirs '("~/.emacs.d/snippets/text-mode"))
(yas-global-mode 1)
(setq yas-indent-line 'auto)
(define-key yas-minor-mode-map (kbd "<tab>") nil)
(define-key yas-minor-mode-map (kbd "TAB") nil)
(define-key yas-minor-mode-map (kbd "<C-tab>") 'yas-expand)
```

Management
----------

Projectile for project management.

``` {.commonlisp org-language="emacs-lisp"}
(use-package projectile)
(projectile-mode +1)
(define-key projectile-mode-map (kbd "C-c p") 'projectile-command-map)
```

Magit
-----

Magit keybinds.

``` {.commonlisp org-language="emacs-lisp"}
(use-package magit)
(global-set-key (kbd "C-x g") 'magit-status)
(global-set-key (kbd "C-x p") 'magit-init)
```

Highlight keywords like todo in files and in magit.

``` {.commonlisp org-language="emacs-lisp"}
(use-package magit-todos)
(magit-todos-mode t)
```

GREP
----

Plain text search using xah-find. Bound to `C-x C-j`.

``` {.commonlisp org-language="emacs-lisp"}
(use-package xah-find)
(global-set-key (kbd "C-x C-j") 'xah-find-text)
```

Search
------

For plain text document search with use of `recoll`, instead of `GREP`.

``` {.commonlisp org-language="emacs-lisp"}
(require 'org-recoll)
;;(load "org-recoll")
(global-set-key (kbd "C-c g") 'org-recoll-search)
(global-set-key (kbd "C-c u") 'org-recoll-update-index)
```

Programming
===========

Elixir
------

Elixir major mode with syntax highlighting etc.

``` {.commonlisp org-language="emacs-lisp"}
(unless (package-installed-p 'elixir-mode)
(package-install 'elixir-mode))
```

Commands:\
Use

-   M-x elixir-format

to format the document following mix styleguide.

Web mode
--------

Web mode and enable rainbow mode for hex colors.

``` {.commonlisp org-language="emacs-lisp"}
(use-package web-mode)
(add-hook 'web-mode-hook
(lambda ()
(rainbow-mode)
(rspec-mode)
(setq web-mode-markup-indent-offset 2)))
```

Golang
------

Golang major mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package go-mode)
(use-package go-errcheck)
```

JavaScript
----------

JavaScript major mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package coffee-mode)
```

Rust
----

Rust major mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package rust-mode)
```

Scala
-----

Scala major mode.

``` {.commonlisp org-language="emacs-lisp"}
(use-package scala-mode
:interpreter
("scala" . scala-mode))
(use-package sbt-mode)
```

Markdown
--------

Github markdown.

``` {.commonlisp org-language="emacs-lisp"}
(use-package markdown-mode
:commands gfm-mode
:mode (("\\.md$" . gfm-mode))
:config
(setq markdown-command "pandoc --standalone --mathjax --from=markdown")
(custom-set-faces
'(markdown-code-face ((t nil)))))
```

More modes
----------

``` {.commonlisp org-language="emacs-lisp"}
(use-package dockerfile-mode)
```

Org mode
========

Bullets
-------

Use org-bullets whenever possible.

``` {.commonlisp org-language="emacs-lisp"}
(use-package org-superstar)
(add-hook 'org-mode-hook (lambda () (org-superstar-mode 1)))
(org-superstar-configure-like-org-bullets)
(setq org-superstar-special-todo-items t)
(setq org-superstar-todo-bullet-alist '(("TODO" . 9744) ("BLOCKED" . 9202) ("PROGRESS" . 8885) ("DONE" . 9745) ("PAL" . 9745) ("IDEA" . 1422) ("NOTE" . 9738) ("INTAKE" . 8227)))
```

Folded
------

Instead of \"...\" show a downward pointing arrow at the end of title.
TODO Change symbol or something.

``` {.commonlisp org-language="emacs-lisp"}
;;(setq org-ellipsis " ➘ ")
(setq org-ellipsis "  ") ;; folding symbol
```

Visibility
----------

Set initial startup visibility to folded.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-startup-folded t)
```

Table of content
----------------

Org-mode toc

``` {.commonlisp org-language="emacs-lisp"}
(use-package toc-org)
(add-hook 'org-mode-hook 'toc-org-enable)
```

Indent mode
-----------

Globally enable `org-indent-mode`

``` {.commonlisp org-language="emacs-lisp"}
(add-hook 'org-mode-hook 'org-indent-mode)
```

Code block
----------

Highlight the entire code block when editing.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-src-fontify-natively t)
```

Todos
-----

Differentiate between org-agenda for calendar events and org-todos for
general todos. General binds.

``` {.commonlisp org-language="emacs-lisp"}
(define-key org-mode-map (kbd "C-c t") 'org-todo)
(define-key org-mode-map (kbd "C-c x") 'todo/done)
```

### Location

Org document storage location for archive and other documents.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-directory "~/documents/supervisors/")

(defun org-file-path (filename)
"Return the absolute address of an org file, given its relative name."
(concat (file-name-as-directory org-directory) filename))

(setq org-archive-location "~/documents/vaults/archive/archive.org::* From %s")
```

### Archive

Hitting `C-c C-x C-s` will mark a todo as done and move it to an
appropriate place in the archive.

``` {.commonlisp org-language="emacs-lisp"}
(defun hrs/mark-done-and-archive ()
"Mark the state of an org-mode item as DONE and archive it."
(interactive)
(org-todo 'done)
(org-archive-subtree))
;; Shortcut to archive
(define-key org-mode-map (kbd "C-c C-x C-s") 'hrs/mark-done-and-archive)
```

### Time

Record the time that a todo was archived.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-log-done 'time)
```

### States

Add new states to the todo cycle to extend the basic TODO and DONE
states that org mode normally provides.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-todo-keywords
      '((sequence "TODO(t)" "PROGRESS(p)" "BLOCKED(b)" "|" "DONE(d)" "PAL(a)")
        (sequence "IDEA(i)" "NOTE(n)" "INTAKE(o)" "|" "DONE(f)")))
```

### Faces

Color for the different states.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-todo-keyword-faces
      '(("IDEA" . (:foreground "#facb20" :weight bold))
        ("TODO" . (:foreground "#af1212" :weight bold))
        ("PROGRESS" . (:foreground "#a8fa80" :weight bold))
        ("BLOCKED" . (:foreground "#b213c4" :weight bold))
        ("PAL" . (:foreground "#30bb03" :weight bold))
        ("NOTE" . (:foreground "#eaa222" :weight bold))
        ("INTAKE" . (:foreground "#bba494" :weight bold))
        ("DONE" . (:foreground "#ffffff" :weight bold))
        ))
```

Tags
----

Default tags with keys.

``` {.commonlisp org-language="emacs-lisp"}
  (setq org-tag-persistent-alist
        '((:startgroup . nil)
          ("HOME" . ?h)
          ("CONTENT" . ?c)
          ("PROJECT" . ?p)
          ("PRÜFUNG" . ?t)
          ("SCHULE" . ?s)
          (:endgroup . nil)
          )
        )
(define-key org-mode-map (kbd "C-c q") 'org-set-tags-command)
(use-package counsel)
(define-key org-mode-map (kbd "C-c C-q") 'counsel-org-tag)

(setq org-tag-faces
      '(
        ("HOME" . (:foreground "GoldenRod" :weight bold))
        ("CONTENT" . (:foreground "LimeGreen" :weight bold))
        ("PROJECT" . (:foreground "OrangeRed" :weight bold))

        ("SCHULE" . (:foreground "DarkMagenta" :weight bold))
        ("PRÜFUNG" . (:foreground "Magenta" :weight bold))
        )
      )
```

Priority
--------

Colored priorities.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-priority-faces '((?A . (:foreground "red" :weight 'bold))
                           (?B . (:foreground "yellow"))
                           (?C . (:foreground "green")))
)
```

Agenda
------

### Agenda

``` {.commonlisp org-language="emacs-lisp"}
(defun todo/done ()
  (interactive)
  (org-todo 'done))

(defun agenda/super (&optional arg)
  (interactive "P")
  (org-agenda arg "c"))

;; Make agenda a global keybind.
(global-set-key (kbd "C-c  a") 'org-agenda)
;; Direct access to super-agenda
(global-set-key [f1] 'agenda/super)

(define-key org-mode-map (kbd "C-c d") 'org-deadline)
(define-key org-mode-map (kbd "C-c s") 'org-schedule)
(define-key org-mode-map (kbd "C-c e") 'org-set-effort)
```

### Commands

`org-super-agenda`

``` {.commonlisp org-language="emacs-lisp"}
(use-package org-super-agenda
  :ensure t
  :config (org-super-agenda-mode)
  )

(setq org-agenda-custom-commands
      '(("c" "Super view"
         ((agenda "" ((org-agenda-span 1)
                      (org-super-agenda-groups
                       '((:name "Today"
                                :time-grid t
                                :date today
                                :order 1)
                         ))))
          (alltodo "" ((org-agenda-overriding-header "Tasks")
                       (org-super-agenda-groups
                        '(
                          (:name "Sapin"
                                 :file-path "tasks.org"
                                 :order 10)
                          (:name "Home"
                                 :tag ("HOME")
                                 :order 20)
                          (:name "Projects & Home"
                                 :tag ("PROJECT")
                                 :order 30)
                          (:name "Content"
                                 :tag ("CONTENT")
                                 :order 40)
                          (:name "Schule"
                                 :tag ("SCHULE")
                                 :order 50)
                          (:name "Prüfung"
                                 :tag ("PRÜFUNG")
                                 :order 60)
                          (:name "Stack"
                                 :file-path "stack.org"
                                 :order 70)
                          (:name "Future"
                                 :file-path "guide.org"
                                 :deadline future
                                 :order 80)
                          ))))))))
```

### Path

Use the entire home directory for agenda files.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-agenda-files (directory-files-recursively "~/documents/supervisor/" "\\.org$"))
```

Add file to the list of included agenda files, bound `C-c v`

``` {.commonlisp org-language="emacs-lisp"}
(define-key org-mode-map (kbd "C-c v") 'org-agenda-file-to-front)
```

Refile
------

Org-refile setup for multi-file task management.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-refile-targets
      '(("~/documents/supervisor/tasks.org" :maxlevel . 1)
        ("~/documents/supervisor/week.org" :maxlevel . 1)))
(define-key org-mode-map (kbd "C-c f") 'org-refile)
(define-key org-mode-map (kbd "C-c C-f") 'org-agenda-refile)
```

Capture
-------

Org-gtd setup.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-default-notes-file "~/documents/supervisors/tasks.org")

(setq org-capture-templates
      '(("c" "Stack TODO" entry (file "~/documents/supervisor/stack.org")
         "* TODO %?\n  %i\n  %a")
        ("v" "Quick Task" entry (file+headline "~/documents/supervisor/tasks.org" "Sapin")
         "* TODO %?\n  %i\n  %a")
        ("e" "New event" entry (file "~/documents/supervisor/events.org")
         "* TODO %?\n  %i\n  %a")
        ("n" "Note" entry (file+datetree "~/documents/supervisor/notes.org")
         "* %?\nEntered on %U\n  %i\n  %a")))

(global-set-key (kbd "C-c c") 'org-capture)
```

Export
------

Allow export to markdown and beamer (for presentations).

``` {.commonlisp org-language="emacs-lisp"}
(eval-after-load "org" '(require 'ox-odt nil t))
```

### Code

Allow `babel` to evaluate Emacs lisp, Ruby, dot, or Gnuplot code.

``` {.commonlisp org-language="emacs-lisp"}
(use-package ob-go)
(use-package ob-elixir)


(use-package gnuplot)
(org-babel-do-load-languages
 'org-babel-load-languages
 '((emacs-lisp . t)
   (ruby . t)
   (dot . t)
   (gnuplot . t)
   (python . t)
      (go . t)
      (sql . t)
      (elixir . t)
      ))
```

Don\'t ask before evaluating code blocks.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-confirm-babel-evaluate nil)

;; inside .emacs file
(setq org-latex-listings 'minted
      org-latex-packages-alist '(("" "minted"))
      org-latex-pdf-process
      '("pdflatex -shell-escape -interaction nonstopmode -output-directory %o %f"
        "pdflatex -shell-escape -interaction nonstopmode -output-directory %o %f"
        "pdflatex -shell-escape -interaction nonstopmode -output-directory %o %f"))

```

### HTML

Disable footer.

``` {.commonlisp org-language="emacs-lisp"}
(setq org-html-postamble nil)
```

Tex
---

### Parse

Parse file after loading it.

``` {.commonlisp org-language="emacs-lisp"}
(setq TeX-parse-self t)
```

### PDF-Latex

``` {.commonlisp org-language="emacs-lisp"}
(setq TeX-PDF-mode t)
```

### Math mode

``` {.commonlisp org-language="emacs-lisp"}
(add-hook 'LaTeX-mode-hook
(lambda ()
(LaTeX-math-mode)
(setq TeX-master t)))
```

Links
-----

Use `C-c C-l` to save the current file for linking. Then use `C-c l` to
insert a new link or write the saved one.

``` {.commonlisp org-language="emacs-lisp"}
(global-set-key (kbd "C-c l") 'org-insert-link)
(define-key org-mode-map (kbd "C-c l") 'org-insert-link)
(global-set-key (kbd "C-c C-l") 'org-store-link)
(define-key org-mode-map (kbd "C-c C-l") 'org-store-link)
```

Design
------

General improvements and design changes for `org-mode`, currently under
the `org-mode` heading instead of `Interface`. Following
[this](/database/https://zzamboni.org/post/beautifying-org-mode-in-emacs/) blog.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (setq org-hide-emphasis-markers t)
  (let* ((variable-tuple
          (cond ((x-list-fonts "ETBembo")         '(:font "ETBembo"))
                ((x-list-fonts "Fira Code") '(:font "Fira Code"))
                (nil (warn "Cannot find a Sans Serif Font.  Install Source Sans Pro."))))
         (base-font-color     (face-foreground 'default nil 'default))
         (headline           `(:inherit default :weight bold :foreground ,base-font-color)))
    (custom-theme-set-faces
     'user
     `(org-level-8 ((t (,@headline ,@variable-tuple))))
     `(org-level-7 ((t (,@headline ,@variable-tuple))))
     `(org-level-6 ((t (,@headline ,@variable-tuple))))
     `(org-level-5 ((t (,@headline ,@variable-tuple))))
     `(org-level-4 ((t (,@headline ,@variable-tuple :height 1.1))))
     `(org-level-3 ((t (,@headline ,@variable-tuple :height 1.25))))
     `(org-level-2 ((t (,@headline ,@variable-tuple :height 1.40))))
     `(org-level-1 ((t (,@headline ,@variable-tuple :height 1.55))))
     `(org-document-title ((t (,@headline ,@variable-tuple :height 2.0 :underline nil))))))

  (add-hook 'org-mode-hook 'visual-line-mode)

  (custom-theme-set-faces
   'user
   '(org-block ((t (:inherit fixed-pitch))))
   '(org-code ((t (:inherit (shadow fixed-pitch)))))
   '(org-document-info ((t (:foreground "dark orange"))))
   '(org-document-info-keyword ((t (:inherit (shadow fixed-pitch)))))
   '(org-indent ((t (:inherit (org-hide fixed-pitch)))))
   '(org-link ((t (:foreground "dark orange" :underline t))))
   '(org-meta-line ((t (:inherit (font-lock-comment-face fixed-pitch)))))
   '(org-property-value ((t (:inherit fixed-pitch))) t)
   '(org-special-keyword ((t (:inherit (font-lock-comment-face fixed-pitch)))))
   '(org-table ((t (:inherit fixed-pitch :foreground "#83a598"))))
   '(org-tag ((t (:inherit (shadow fixed-pitch) :weight bold :height 0.8))))
   '(org-verbatim ((t (:inherit (shadow fixed-pitch)))))))
```

Apps
----

Default apps for file extensions.

``` {.commonlisp org-language="emacs-lisp"}
(eval-after-load "org"
  '(progn
     (setcdr (assoc "\\.pdf\\'" org-file-apps) "google-chrome-stable %s")))
```

Table
-----

Use pretty tables. <http:https://github.com/Fuco1/org-pretty-table>

``` {.commonlisp org-language="emacs-lisp"}
(require 'org-pretty-table)
(add-hook 'org-mode-hook (lambda () (org-pretty-table-mode 1)))
```

Navigation
----------

``` {.commonlisp org-language="emacs-lisp"}
(setq org-src-tab-acts-natively t)
```

Content
=======

org-roam
--------

(Configured under `/content` instead of `/org-mode`, might change in the
future) Org-roam setup for the entire `documents` directory.

``` {.commonlisp org-language="emacs-lisp"}
  (use-package org-roam
    :ensure t
    :hook
    (after-init . org-roam-mode)
    :custom
    (org-roam-directory "~/documents/vaults/knowledge/")
    :bind (:map org-roam-mode-map
              (("C-c r l" . org-roam)
               ("C-c r f" . org-roam-find-file)
               ("C-c r u" . org-roam-unliked-references)
               ("C-c r g" . org-roam-graph-show)
               ("C-c r s" . org-roam-server-mode))
              :map org-mode-map
              (("C-c r i" . org-roam-insert))
              (("C-c r I" . org-roam-insert-immediate))))

(setq org-roam-db-location "~/documents/vaults/org-roam.db")
(add-hook 'after-init-hook 'org-roam-mode)

(use-package company-org-roam
  :ensure t
  :config
  (push 'company-org-roam company-backends))
```

org-roam-server
---------------

Dedicated config for `org-roam-server`

``` {.commonlisp org-language="emacs-lisp"}
(use-package org-roam-server
  :ensure t
  :config
  (setq org-roam-server-host "127.0.0.1"
        org-roam-server-port 8080
        org-roam-server-authenticate nil
        org-roam-server-export-inline-images t
        org-roam-server-serve-files nil
        org-roam-server-served-file-extensions '("pdf" "mp4" "ogv" "jpg" "png")
        org-roam-server-network-poll t
        org-roam-server-network-arrows nil
        org-roam-server-network-label-truncate t
        org-roam-server-network-label-truncate-length 60
        org-roam-server-network-label-wrap-length 20))
```

org-roam-config
---------------

Some configurations for `org-roam` inspired by [Care to share configs
for how you use
org-roam?](https://www.reddit.com/r/emacs/comments/gsv5np/care_to_share_configs_for_how_you_use_orgroam/)

``` {.commonlisp org-language="emacs-lisp"}
  (defun org-force-open-current-window ()
    (interactive)
    (let ((org-link-frame-setup (quote
                                 ((vm . vm-visit-folder)
                                  (vm-imap . vm-visit-imap-folder)
                                  (gnus . gnus)
                                  (file . find-file)
                                  (wl . wl)))
                                ))
      (org-open-at-point)))
  ;; Depending on universal argument try opening link
  (defun org-open-maybe (&optional arg)
    (interactive "P")
    (if arg
        (org-open-at-point)
      (org-force-open-current-window)
      )
    )

  ;; redefined to use org-open-maybe
  (defun ace-link-org (&optional arg)
    "Open a visible link in an `org-mode' buffer."
    (interactive "P")
    (require 'org)
    (let ((pt (avy-with ace-link-org
                        (avy-process
                         (mapcar #'cdr (ace-link--org-collect))
                         (avy--style-fn avy-style)))))
      (when (numberp pt)
        (goto-char pt)
        (org-open-maybe arg))
      ))

  (defun mmr/org-roam-insert-replace-region-with-link-and-follow ()
    (interactive )
    (let ((title (buffer-substring (mark) (point)) )
          (top (current-buffer)))
      (org-roam-find-file title)
      (let ((target-file (buffer-file-name (buffer-base-buffer)))
          (note-buffer (current-buffer)))
        (switch-to-buffer top nil t)
        (kill-region (mark) (point))
        (insert (concat "[[" target-file "][" title "]]"))
        (switch-to-buffer note-buffer nil t)
        (save-buffer))))

  (define-key org-mode-map (kbd "C-c o") 'org-open-maybe)
  (define-key org-mode-map (kbd "C-c r r") 'mmr/org-roam-insert-replace-region-with-link-and-follow)

(setq org-roam-graph-exclude-matcher '("tasks" "archive"))
```

Org-roam-capture
----------------

``` {.commonlisp org-language="emacs-lisp"}
  (setq org-roam-capture-templates
          '(("i" "With fixed filename -> For Export." plain (function org-roam--capture-get-point)
             "%?"
             :file-name "${slug}"
             :head "#+title: ${title}
#+SETUPFILE:~/.hugo.org\n"
             :unnarrowed t)
            ("o" "Unchanged org-roam insert." plain (function org-roam--capture-get-point)
             "%?"
             :head "#+title: ${title}\n")))
```

Writing
-------

Turn off all distractions.

``` {.commonlisp org-language="emacs-lisp"}
(when (display-graphic-p)
  (setq left-margin-width 2)
  (setq right-margin-width 2)

  (use-package darkroom)
  (define-key org-mode-map (kbd "C-c m") 'darkroom-mode)

  (global-set-key (kbd "C-c j") 'display-line-numbers-mode)


  (setq org-startup-indented t
        org-bullets-bullet-list '(" ") ;; no bullets, needs org-bullets package
        org-ellipsis " > " ;; folding symbol
        org-pretty-entities t
        org-hide-emphasis-markers t
        ;; show actually italicized text instead of /italicized text/
        org-agenda-block-separator ""
        org-fontify-whole-heading-line t
        org-fontify-done-headline t
        org-fontify-quote-and-verse-blocks t))

```

Hoth
----

``` {.commonlisp org-language="emacs-lisp"}
(load-file "~/.emacs.d/hoth/hoth.el")
(require 'hoth)
(define-key org-mode-map (kbd "C-c r h") 'hoth-total)
```

Deft
----

Deft for taking quick notes and storing them in plain text.

``` {.commonlisp org-language="emacs-lisp"}
(use-package deft)
(setq deft-default-extension "org")
(setq deft-extensions '("org"))
(setq deft-directory "~/documents/supervisor/inbox/")
(setq deft-recursive t)
(setq deft-text-mode 'org-mode)
(setq deft-use-filename-as-title t)

(define-key org-mode-map (kbd "C-c n") 'deft)
(define-key org-mode-map (kbd "M-c n") 'deft-find-file)
```

Special tags
------------

Special tags used for content inside of a table.

``` {.commonlisp org-language="emacs-lisp"}
(defun org-table-tag (choice)
  "Select one of the two available modes as CHOICE."
  (interactive)
  (let ((completion-ignore-case  t))
  (list (completing-read "File type: " '(":fun:" ":wiki:" ":rec:") nil t)))
  )
  (require 'org-inlinetask)
```

Browser
-------

Use `chrome` as the default browser for links.

``` {.commonlisp org-language="emacs-lisp"}
(setq browse-url-browser-function 'browse-url-generic browse-url-generic-program "google-chrome-stable")
```

Games
-----

Malyon: Text adventure interface, games located under `/games`

``` {.commonlisp org-language="emacs-lisp"}
(use-package malyon)
```

Hugo
----

### Header

Automatically insert `ox-hugo` header.

``` {.commonlisp org-language="emacs-lisp"}
(defun hugo-header ()
  (interactive)
  (insert "#+HUGO_SECTION: posts \n#+HUGO_BASE_DIR:~/documents/projects/knowledge")
  )
(define-key org-mode-map (kbd "C-c r e") 'hugo-header)
```

### Quick export

``` {.commonlisp org-language="emacs-lisp"}
(defun hugo-export ()
  (org-export-dispatch "H h")
  )
(define-key org-mode-map (kbd "C-c r d") 'hugo-export)
```

### ox

``` {.commonlisp org-language="emacs-lisp"}
(use-package ox-hugo
  :ensure t            ;Auto-install the package from Melpa (optional)
  :after ox)
```
