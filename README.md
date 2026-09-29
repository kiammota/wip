# wip

A tiny Unix-style CLI for keeping track of what you're currently working on.

`wip` stores the current work-in-progress status in a `.wip` file in the current directory.

## Usage

Show the current WIP:

```bash
wip
```

Set the current WIP:

```bash
wip "migrating the database to PostgreSQL"
```

This creates or updates `.wip` with:

```text
[2026/09/29 13:22] migrating the database to PostgreSQL
```

Because `.wip` lives in the current directory, each project can have its own WIP status.

## Installation

```bash
go install github.com/KiamMota/wip@latest
```


