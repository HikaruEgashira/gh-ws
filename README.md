<h2 align="center">
    <p align="center">gh-ws</p>
</h2>

<h3 align="center">
🔹<a  href="https://github.com/HikaruEgashira/gh-ws/issues">Report Bug</a> &nbsp; &nbsp;
🔹<a  href="https://github.com/HikaruEgashira/gh-ws/issues">Request Feature</a>
</h3>

VSCode workspace management extension for GitHub CLI.

## Installation

```bash
gh extension install HikaruEgashira/gh-ws
```

## Usage

```bash
$ gh ws --help
Usage:
  gh ws               ... Generate workspace and open in VSCode
  gh ws init          ... Generate workspace from git worktrees
  gh ws init <name>   ... Generate workspace with custom name
  gh ws init --multi  ... Generate workspace from multiple repositories
  gh ws init <name> --multi ... Generate multi-repository workspace with custom name
  gh ws list          ... List existing workspaces
  gh ws remove        ... Remove workspace (interactive selection)
  gh ws -- <command>  ... Search via fzf and run <command> in the selected workspace directory
  gh ws <command>     ... Search via fzf and run <command> with selected workspace as argument
```

### Examples

```bash
# Generate workspace from current repository worktrees and open in VSCode
gh ws

# Create workspace file only (no VSCode launch)
gh ws init

# Create workspace with custom name
gh ws init my-workspace

# Create multi-repository workspace (interactive selection)
gh ws init --multi

# Create named multi-repository workspace
gh ws init frontend --multi

# List all existing workspaces
gh ws list

# Remove workspace (interactive selection)
gh ws remove

# Open workspace in VS Code (path as argument)
gh ws code

# Run commands in the selected workspace directory
gh ws -- ls
gh ws -- git status
```

## Command Execution Modes

There are two ways to execute commands with selected workspaces:

1. **`gh ws <command>`** - Passes the workspace path as an argument to the command
   - Example: `gh ws code` → executes `code /path/to/selected/workspace.code-workspace`
   - Useful for editors and tools that accept workspace files as arguments

2. **`gh ws -- <command>`** - Changes to the workspace directory and executes the command
   - Example: `gh ws -- ls` → changes to workspace directory then runs `ls`
   - Useful for commands that need to run within the workspace directory

## Requirements

- [GitHub CLI](https://cli.github.com/)
- [fzf](https://github.com/junegunn/fzf)
- [VSCode](https://code.visualstudio.com/) with `code` command in PATH
- Must be used within a git repository (for single-repo workspaces)

## How it works

`gh ws` creates VSCode workspace files in `~/ghq/workspaces/` directory. For single-repository workspaces, it detects git worktrees in the current repository. For multi-repository workspaces, it allows selection of multiple repositories from the entire `~/ghq` directory structure.

## Integration with other gh extensions

Works well with:
- [`gh-q`](https://github.com/HikaruEgashira/gh-q): Quick repository navigation
- [`gh-wt`](https://github.com/HikaruEgashira/gh-wt): Git worktree management

```bash
# Example workflow
gh q                    # Select repository
gh wt add feature/new   # Create new worktree
gh ws init              # Create/update workspace
```