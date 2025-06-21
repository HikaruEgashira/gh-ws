# gh-ws

GitHub CLI extension for managing VSCode workspaces with git worktrees.

## Installation

```bash
gh extension install HikaruEgashira/gh-ws
```

## Usage

### Generate workspace from git worktrees
```bash
gh ws init
```

### List existing workspaces
```bash
gh ws list
```

### Clean workspace (remove deleted worktrees)
```bash
gh ws clean
```

### Open workspace
```bash
gh ws
```

## Features

- Automatically detects git worktrees in the current repository
- Creates VSCode workspace files in `~/ghq/workspaces/`
- Opens workspace directly in VSCode
- Cleans up removed worktrees from workspace configuration

## Integration with other gh extensions

Works well with:
- `gh-q`: Quick repository navigation
- `gh-wt`: Git worktree management

```bash
# Example workflow
gh q                    # Select repository
gh wt add feature/new   # Create new worktree
gh ws init              # Create/update workspace
```