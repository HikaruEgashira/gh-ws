<h2 align="center">
    <p align="center">gh-ws</p>
</h2>

<h3 align="center">
🔹<a  href="https://github.com/HikaruEgashira/gh-ws/issues">Report Bug</a> &nbsp; &nbsp;
🔹<a  href="https://github.com/HikaruEgashira/gh-ws/issues">Request Feature</a>
</h3>

VSCode workspace management and Claude Code integration extension for GitHub CLI.

## Installation

```bash
gh extension install HikaruEgashira/gh-ws
```

## Usage

```bash
$ gh ws --help
Usage:
  gh ws               ... Create/open workspace from current repository worktrees in VSCode
  gh ws init          ... Generate workspace from git worktrees
  gh ws init <name>   ... Generate workspace with custom name
  gh ws init --multi  ... Generate workspace from multiple repositories
  gh ws init <name> --multi ... Generate multi-repository workspace with custom name
  gh ws list          ... List existing workspaces
  gh ws edit          ... Edit workspace file in EDITOR
  gh ws remove        ... Remove workspace (interactive selection)
  gh ws claude        ... Run Claude with current workspace folders or select workspace
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

# Edit workspace file in your preferred editor
gh ws edit

# Remove workspace (interactive selection)
gh ws remove

# Open workspace in VS Code (path as argument)
gh ws code

# Run commands in the selected workspace directory
gh ws -- ls
gh ws -- git status

# Run Claude Code with workspace folders
gh ws claude
```

## Command Execution Modes

There are two ways to execute commands with selected workspaces:

1. **`gh ws <command>`** - Passes the workspace path as an argument to the command
   - Example: `gh ws code` → executes `code /path/to/selected/workspace.code-workspace`
   - Useful for editors and tools that accept workspace files as arguments

2. **`gh ws -- <command>`** - Changes to the workspace directory and executes the command
   - Example: `gh ws -- ls` → changes to workspace directory then runs `ls`
   - Useful for commands that need to run within the workspace directory

3. **`gh ws claude`** - Runs Claude Code with workspace folders
   - Example: `gh ws claude` → executes `claude --add-dir "/path/folder1" --add-dir "/path/folder2"`
   - Automatically finds workspace file in current directory or prompts for selection
   - Extracts all folder paths from the workspace and passes them to Claude

## Requirements

- [GitHub CLI](https://cli.github.com/)
- [fzf](https://github.com/junegunn/fzf)
- [VSCode](https://code.visualstudio.com/) with `code` command in PATH (for VSCode integration)
- [Claude Code](https://docs.anthropic.com/en/docs/claude-code) (for `gh ws claude` command)
- Must be used within a git repository (for single-repo workspaces)

## How it works

`gh ws` creates VSCode workspace files in `~/ghq/workspaces/` directory. For single-repository workspaces, it detects git worktrees in the current repository. For multi-repository workspaces, it allows selection of multiple repositories from the entire `~/ghq` directory structure and automatically includes their worktrees.

## Integration with other gh extensions

Works well with:
- [`gh-q`](https://github.com/HikaruEgashira/gh-q): Quick repository navigation
- [`gh-wt`](https://github.com/HikaruEgashira/gh-wt): Git worktree management
- [Claude Code](https://docs.anthropic.com/en/docs/claude-code): AI-powered coding assistant

```bash
# Example workflow
gh q                    # Select repository
gh wt add feature/new   # Create new worktree
gh ws init              # Create/update workspace
gh ws claude            # Start Claude Code with all workspace folders

# Quick development workflow
gh q -- gh wt add feature/api    # Create worktree in selected repo
gh q -- gh ws                    # Create workspace and open VSCode
gh q -- gh ws claude             # Start Claude Code with workspace
```