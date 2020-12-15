---
title: Saturn V
description: Saturn V rocket simulator.
published: 1
date: 2020-12-15T12:24:21.646Z
tags: 
editor: undefined
dateCreated: 2020-12-03T20:59:37.704Z
---

Rocket simulator for `Saturn V` and other Rockets. With support for more
Rockets and Planets. Written in *simple* `Golang` Code. A more detailed
documentation can be found in `documentation.org` or
`documentation.pdf`. Only tested on Linux. Most likely no Windows
support.

Run
===

1.  Build the program using `go build` in the project root.
2.  Add any rockets in the `rockets/` directory.
3.  Adjust settings and add planets in `settings.json`
4.  Execute the program with `./saturnv`

Rocket
======

Files must be located in the `rockets/` directory.

``` {.javascript}
{
    "name": "Saturn V",
    //Array of rocket stages, variable length. Minimum 1
    "stages": [
        {
            "name": "Lower stage",
            "full_weight": 2286000,
            "empty_weight": 135000,
            "exit_velocity": -2980
            "burn_time": 161,
        },
        {
            "name": "Center stage",
            "full_weight": 491000,
            "empty_weight": 39000,
            "exit_velocity": -4160,
            "burn_time": 387
        },
        {
            "name": "Upper stage",
            "full_weight": 120000,
            "empty_weight": 13300,
            "exit_velocity": -4180
            "burn_time": 319
        }
    ],
    "payload_weight": 41000,
    "steps": 50
    "c": 0.52,
    "area": 78.5
}
```

Settings
========

This file must be available as `settings.json` in project root.

``` {.javascript}
{
    "profile": {
        "gravity": true,
        "air": true
    }
    "planet": {
        "name": "Earth",
        "gravity": 9.81
    },
    {
        "name": "Mars",
        "gravity": 3.771
    },
    {
        "name": "Jupiter",
        "gravity": 24.79
    }
}
```

Dependencies
============

The entire program should run without any external dependencies except
for `Figlet`, which is used for the countdown.
