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
  gh ws edit          ... Edit workspace file in EDITOR
  gh ws remove        ... Remove workspace (interactive selection)
  gh ws -- <command>  ... Search via fzf and run <command> in the selected workspace directory
  gh ws <command>     ... Search via fzf and run <command> with selected workspace as argument
```

### Basic Usage

```bash
# 1. Create workspace from current repository and open VSCode
gh ws

# 2. Create workspace from multiple repositories
gh ws init --multi

# 3. Open existing workspace
gh ws code
```

## User Stories

**Single Repository Development**
```bash
cd ~/ghq/github.com/user/my-app
gh wt add feature/auth        # Create worktree for feature
gh ws                         # Create workspace with all worktrees & open VSCode
```

**Multi-Repository Project**
```bash
gh ws init frontend --multi   # Select related repos: frontend, backend, docs
                             # Creates workspace with all repos + their worktrees
```

**Daily Workflow**
```bash
gh ws list                   # See all workspaces
gh ws code                   # Quick open workspace in VSCode
gh ws edit                   # Customize workspace settings
```

## Requirements

- [GitHub CLI](https://cli.github.com/)
- [fzf](https://github.com/junegunn/fzf)
- [VSCode](https://code.visualstudio.com/) with `code` command in PATH

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