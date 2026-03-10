# jot

A lightweight CLI for managing notes — locally per project, or globally.

## Usage

### Add a note

```bash
jot add --title "Write Tests"
jot add --title "Deploy steps" --description "Run migrations first" --global
```

### List notes

```bash
jot list            # notes in current directory
jot list --global   # global notes
jot list --all      # all notes across all directories
```

```
ID   TITLE          CREATED
--   -----          -------
1    Add Test       2026-03-09
2    Deploy steps   2026-03-09
```

### Read a note

```bash
jot read 1
```

### Delete a note

```bash
jot rm 1
```

