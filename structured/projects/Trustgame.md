---
title: Trustgame
description: Simple (and bad) implementation of a basic game theory game.
published: 1
date: 2020-12-03T21:06:14.523Z
tags: 
editor: markdown
dateCreated: 2020-12-03T21:06:14.523Z
---

# Trustgame

## by Jeykey @ jeykey.net

Simple trustgame implemented in Elixir Phoenix, hosted on Gitlab

### How to play:
The goal of the game is to earn the most ammount of coins possible. This is achieved by playing against other players and tricking them into giving them more money then you give them. In a perfect world, both players would give each other the maximum ammount possible each round. But the nobody wins. So you have to gain the trust of your opponent until you can trick them and make a profit. After that, your Score (ELO) will be updated with a custom ELO algorithm (still quite broken, patch coming soooooon.....    (trust.me)).
It is important to note, that no currency is created, unless new accounts are created. So all money in the game gets moved arround from player to player. 

### Notes before playing:
This game was build as an learing experience for Elixir Phoenix. It is incomplete, has loads of bugs and delivers very inconsistant results. 
This is more a proof of concept and a learning experience than an actual game or product.

## Todo's
- Add archives
- Update Elo via custom algorithm (partially working), but with bugs
- Automaticly log users in, after account creation
- Update /play screen on creation of a new game
- Only show `open` games in the /play list
- Make `leaderboard` actual leaderboard => Only show top 10 players
- Fix UI for diffrent screen sizes
- Fix timer bug => Store `timer_started` in the DB

## Bugs:
List of bugs that have already been found, hopefully will be fixed as soon as possible.
- After end of game, User Data will not be updated. Cause: Unknown
- Opponent inputs may be processed as player inputs. Cause: Overlapping socket logic.
- Rejoining an ongoing Game can pause the css parsing or not work at all. Cause: Unknown.
- Joining a deleted match will result in an unhandled error. Cause: No fitting edgecase handler.
- Editing the js or causing any error in it may result in strange behaviour on the entire site => Delete links and Logout wont work. Cause: Phoenix Design flaw or Bug in Phoenix Code.

## To start your Phoenix server:

  * Install dependencies with `mix deps.get`
  * Create and migrate your database with `mix ecto.create && mix ecto.migrate`
  * Install Node.js dependencies with `cd assets && npm install`
  * Start Phoenix endpoint with `mix phx.server`

Now you can visit [`localhost:8080`](http://localhost:4000) from your browser.

# License:
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.