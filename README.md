# README

## About
This is Mood tracker app built with [Wails Framewor](https://wails.io/). Purpose is to give your mood a rating from 1 to 10 once a day and to see statistics above.

## Screenshots
#### Initial empty layout
![Empty](/assets/empty-day.webp)

#### Filled up
![Filled up](/assets/rated-day-with-chart.webp)

## Stack
[Golang](https://go.dev/) on backend, [SQlite](https://www.sqlite.org/index.html) as Database solution and [Vue 3](https://vuejs.org/) on frontend side.

## Installation
1. Pull repo.
2. Build using `wails build` (take a look at their installation guide)
3. Executable binary will be at `./build/bin/mood-tracker`
   
## Extra details
It's developed and tested only within Linux Mint 21.2 x86_64. App's data stored inside of `~/.cache/mood-tracker/` directory.

Domain-driven-design was chosen as folder structure and architecture.

[Kanagawa](https://github.com/rebelot/kanagawa.nvim) used as colorscheme.

