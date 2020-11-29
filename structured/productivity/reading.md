---
title: Reading system
description: Description of a system for collecting notes and ideas from books.
published: 1
date: 2020-11-29T21:27:46.773Z
tags: struc
editor: markdown
dateCreated: 2020-11-29T10:28:01.136Z
---

# Problem
Beim lesen ist das Verstehen der Informationen weniger das Problem als das Verbinden und Wiederaufbringen. Daher benötigt man ein System, mit welchem Informationen direkt kategorisiert werden können, um spätere Verarbeitung zu erleichtern.

Aber das einfache Anstreichen oder Unterstreichen bringt meistens nichts, da man beim späteren überfliegen der Informationen trotzdem den aktuellen Kontext verstehen muss, nur um die angeblich wichtige Textstelle zu interpretieren.

Das ausführliche Ausarbeiten und Kritisieren der Ideen am Rand des Buchs oder zwischen den Zeilen ist allerdings meist zu viel Arbeit und bricht den Fluss des lesens. Dazu bleiben am Ende hunderte von Notizen in Büchern verstreut, ohne eine Möglichkeit diese Effizient wieder ein zu sehen.

# System
Im Laufe der Zeit wurde es nötig, ein zentrales System zu erarbeiten. Durch viele Iterationen und Experimente wurden die folgenden vier Komponenten erarbeitet:
## Supervisor
Mit dem Supervisor enthält das System eine modifizierte, angepasste Version von **Getting Things Done**, welches mit Komponenten von **Building a second brain** erweitert wurden. 
Der Supervisor enthält folgende Dateien / Komponenten:
- GSD
(Getting *stuff* done)
Diese Datei enthält meist nur wenige Aufgaben, welchen jeden Abend neu hinzugefügt werden. Diese sollen dann im Laufe des Tages erledigt werden. Das System hierfür erinnert an die **Ivy-Lee-Methode**.
- Inbox
Jegliche Aufgaben, Notizen oder Events landen als erstes in der Inbox. Diese enthält noch keien Struktur und die Aufgaben werden nicht von hier aus direkt erledigt.
- Projects
Jeden Abend werden neue Einträge der Inbox den verschiedenen Projekten zugewiesen. Dabei ist der Begriff Projekt eher grob gefasst und für alle Gruppierungen von Aufgaben gibt es ein eigenes Projekt. Hier werden die Aufgaben auch mit org-headings sortiert.
Zusätzlich sollte es für jedes Projekt ein "NEXT"-TODO geben. Dieses sollte ohne grosse Vorbereitung oder direkt im Anschluss an andere Aufgaben zu erledigen sein und einen guten Einsteigspunkt in ein Projekt geben.
- Events
Die meisten der oben beschriebenen Aufgaben haben keine klare Deadline oder einen Zeitpunkt. Ausser den **GSD** Aufgaben gibt es keinen definieren Zeitpunkt für die Aufgaben. Sollte es aber Terimne oder ähnliches geben,  werden diese hier gespeichert und entsprechend angezeigt.
- Last
Alle Aufgaben, Notizen und Termine werden in diese Datei archiviert. Von hier aus wird ein kurzer Wochenrückblick geschrieben und der gesamte Tree wird dann ins Archiv verschoben.
## Projects
Hier befinden sich die tatsächlichen Dateien für Projekte. Dabei geht es nicht um Aufgaben, Hilfsmaterial oder Notizen, sondern um Code und Libraries. Von hier aus wird dann direkt an den einzelnen Projekten gearbeitet. Dabei handelt es sich meistens um git-repositories.
## Vaults
Für extrahierte Notizen, Konzepte und Gedanken wird ist in der Datenbank Platz vorhanden. Dort werden sie mit org-roam verbunden und mit diesem Wiki dargestellt. Dazu gibt es ein zweites Vault, welches für schulische Dateien verantwortlich ist. Diese haben keinen praktischen, langfristigen Nutzen, sind aber trotzdem für den Moment von Nöten.
Hier wird auch Project-Hoth benutzt, durch welches nur org-mode Dateien in den Vaults zu finden sind.
## Archive
Dieser Teil (Ordner) wird mehrheitlich von automatisierten Systemen wie org-archive verwendet. Dort werden auch Kontext-Lose Rocketbook Notizen gelagert.