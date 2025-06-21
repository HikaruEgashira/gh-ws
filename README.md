# gh-ws

GitHub CLI extension for managing VSCode workspaces with git worktrees.

## Installation

```bash
gh extension install HikaruEgashira/gh-ws
```

## Requirements

- [GitHub CLI](https://cli.github.com/)
- [fzf](https://github.com/junegunn/fzf) (for interactive selection in related tools)
- [VSCode](https://code.visualstudio.com/) with `code` command in PATH
- Must be used within a git repository

## Usage

### Generate workspace and open in VSCode
```bash
gh ws               # Create workspace and open in VSCode
```

### Generate workspace from git worktrees
```bash
gh ws init          # Create workspace file only
```

### List existing workspaces
```bash
gh ws list
```

### Clean workspace (remove deleted worktrees)
```bash
gh ws clean
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