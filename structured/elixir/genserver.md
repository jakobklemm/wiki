---
title: GenServer boilerplate
description: Required components for a GenServer implementation.
published: 1
date: 2020-12-01T12:03:34.155Z
tags: 
editor: markdown
dateCreated: 2020-12-01T11:05:41.379Z
---

## Komponenten
- Supervisor
- DynamicSupervisor
- Worker / GenServer
- Mix file

## Supervisor
``` elixir
defmodule App.Application do
  use Application
	# Hier wird nur der dynamische Supervisor gestartet, die eigentlichen Prozesse werden erst dort registriert.

  def start(_, _) do
    children = [
    	# Leeres Array als Start-Argument für den dynamischen Supervisor.
      {App.Dynamic, []},
    ]
  opts = [strategy: :one_for_one, name: {:global, __MODULE__}]
  Supervisor.start_link(children, opts)
  end
end
```

## mix.exs
``` elixir

defmodule Mas.MixProject do
  use Mix.Project

  def project do
    []
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger],
      mod: {App.Application, []}
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    []
  end
end
```

## DynamicSupervisor
``` elixir
defmodule App.Dynamic do
  use DynamicSupervisor

  def start_link(_) do
  	# Register globally (chane for registry or multiple supervisors)
    DynamicSupervisor.start_link(__MODULE__, :ok, name: {:global, __MODULE__})
  end

  def init(_) do
    DynamicSupervisor.init(strategy: :one_for_one)
  end

  def start_child(_) do
    DynamicSupervisor.start_child(
      __MODULE__,
      {App.Worker, []}
    )
  end
end
```

## Worker
``` elixir
defmodule App.Worker do
  use GenServer
  use DynamicSupervisor

  def start_link(uuid, params) do
    GenServer.start_link(__MODULE__, params, name: {:global, uuid})
  end

  def getCurrentUser(uuid) do
    GenServer.call({:global, uuid}, {:getState})
  end

  def updateUser(uuid, state) do
    GenServer.cast({:global, uuid}, {:new, state})
  end

  def getUserPID(uuid) do
    GenServer.whereis({:global, uuid})
  end

  #- - - - - - - - - - - - - - - - - - -

  def init(initial_state) do
    {:ok, initial_state}
  end

  def handle_call({:getState}, _from, current_state) do
    {:reply, current_state, current_state}
  end

  def handle_cast({:update_user, user}, state) do
    {:noreply, user}
  end

  def handle_info({:cache_timeout, counter}, state) do
     {:stop, :normal, state}
  end
end

```