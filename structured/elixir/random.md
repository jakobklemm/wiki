---
title: Random
description: Elixir random and crypto help.
published: 1
date: 2021-01-27T20:27:08.865Z
tags: 
editor: markdown
dateCreated: 2021-01-27T20:27:08.865Z
---

# 160 Bit Node-ID
``` elixir
iex> Stream.repeatedly(fn -> :rand.uniform 255 end) |> Enum.take(20) |> :binary.list_to_bin |> :binary.decode_unsigned
418761851173619023669700471877685253548864849034
```

# Seed & Entropy
``` elixir
iex> :rand.seed(:exs64, :os.timestamp)
{%{
   max: 18446744073709551615,
   next: #Function<10.10897371/1 in :rand.mk_alg/1>,
   type: :exs64
}, 15599426412838157331}
```
