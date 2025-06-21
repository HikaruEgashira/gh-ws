package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type WorkspaceFolder struct {
	Path string `json:"path"`
}

type Workspace struct {
	Folders []WorkspaceFolder `json:"folders"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gh ws <command>")
		fmt.Println("Commands:")
		fmt.Println("  init    Generate workspace from git worktrees")
		fmt.Println("  list    List existing workspaces")
		fmt.Println("  clean   Remove deleted worktrees from workspace")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		initWorkspace()
	case "list":
		listWorkspaces()
	case "clean":
		cleanWorkspace()
	default:
		openWorkspace()
	}
}

func initWorkspace() {
	worktrees, err := getWorktrees()
	if err != nil {
		fmt.Printf("Error getting worktrees: %v\n", err)
		os.Exit(1)
	}

	if len(worktrees) == 0 {
		fmt.Println("No worktrees found")
		return
	}

	workspace := createWorkspace(worktrees)
	workspacePath := getWorkspacePath()

	if err := saveWorkspace(workspace, workspacePath); err != nil {
		fmt.Printf("Error saving workspace: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Workspace created: %s\n", workspacePath)
	
	// Open VSCode with the workspace
	cmd := exec.Command("code", workspacePath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error opening VSCode: %v\n", err)
	}
}

func getWorktrees() ([]string, error) {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var worktrees []string
	lines := strings.Split(string(output), "\n")
	
	for _, line := range lines {
		if strings.HasPrefix(line, "worktree ") {
			path := strings.TrimPrefix(line, "worktree ")
			worktrees = append(worktrees, path)
		}
	}

	return worktrees, nil
}

func createWorkspace(worktrees []string) Workspace {
	var folders []WorkspaceFolder
	
	for _, worktree := range worktrees {
		folders = append(folders, WorkspaceFolder{Path: worktree})
	}

	return Workspace{Folders: folders}
}

func getWorkspacePath() string {
	homeDir, _ := os.UserHomeDir()
	workspaceDir := filepath.Join(homeDir, "ghq", "workspaces")
	
	// Create directory if it doesn't exist
	os.MkdirAll(workspaceDir, 0755)
	
	// Get repository name for workspace filename
	repoName := getRepoName()
	return filepath.Join(workspaceDir, repoName+".code-workspace")
}

func getRepoName() string {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "workspace"
	}

	url := strings.TrimSpace(string(output))
	// Extract repo name from git URL
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		name := parts[len(parts)-1]
		name = strings.TrimSuffix(name, ".git")
		return name
	}

	return "workspace"
}

func saveWorkspace(workspace Workspace, path string) error {
	data, err := json.MarshalIndent(workspace, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func listWorkspaces() {
	homeDir, _ := os.UserHomeDir()
	workspaceDir := filepath.Join(homeDir, "ghq", "workspaces")
	
	files, err := os.ReadDir(workspaceDir)
	if err != nil {
		fmt.Printf("Error reading workspace directory: %v\n", err)
		return
	}

	fmt.Println("Available workspaces:")
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".code-workspace") {
			name := strings.TrimSuffix(file.Name(), ".code-workspace")
			fmt.Printf("  %s\n", name)
		}
	}
}

func cleanWorkspace() {
	worktrees, err := getWorktrees()
	if err != nil {
		fmt.Printf("Error getting worktrees: %v\n", err)
		os.Exit(1)
	}

	workspacePath := getWorkspacePath()
	workspace := createWorkspace(worktrees)

	if err := saveWorkspace(workspace, workspacePath); err != nil {
		fmt.Printf("Error updating workspace: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Workspace cleaned: %s\n", workspacePath)
}

func openWorkspace() {
	workspacePath := getWorkspacePath()
	
	if _, err := os.Stat(workspacePath); os.IsNotExist(err) {
		fmt.Printf("Workspace not found: %s\n", workspacePath)
		fmt.Println("Run 'gh ws init' to create a workspace")
		return
	}

	cmd := exec.Command("code", workspacePath)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error opening VSCode: %v\n", err)
	}
}