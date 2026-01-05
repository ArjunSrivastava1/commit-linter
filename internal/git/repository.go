package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Repository represents a Git repository
type Repository struct {
	Path string
}

// NewRepository creates a new Repository instance
func NewRepository(path string) (*Repository, error) {
	// If no path provided, use current directory
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get current directory: %v", err)
		}
		path = dir
	}

	// Check if it's a Git repository
	if !isGitRepository(path) {
		return nil, fmt.Errorf("not a Git repository: %s", path)
	}

	return &Repository{Path: path}, nil
}

// isGitRepository checks if a directory is a Git repository
func isGitRepository(path string) bool {
	gitDir := filepath.Join(path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return true
	}

	// Check if we're inside a Git worktree
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = path
	if err := cmd.Run(); err == nil {
		return true
	}

	return false
}

// GetGitDir returns the .git directory path
func (r *Repository) GetGitDir() (string, error) {
	// First try .git directory
	gitDir := filepath.Join(r.Path, ".git")
	if _, err := os.Stat(gitDir); err == nil {
		return gitDir, nil
	}

	// Fallback to git command
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	cmd.Dir = r.Path
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get Git directory: %v", err)
	}

	dir := strings.TrimSpace(string(output))
	if !filepath.IsAbs(dir) {
		dir = filepath.Join(r.Path, dir)
	}

	return dir, nil
}

// GetHooksDir returns the hooks directory path
func (r *Repository) GetHooksDir() (string, error) {
	gitDir, err := r.GetGitDir()
	if err != nil {
		return "", err
	}

	hooksDir := filepath.Join(gitDir, "hooks")
	return hooksDir, nil
}

// GetCommitMessageFilePath returns the path to commit message file
func (r *Repository) GetCommitMessageFilePath() (string, error) {
	// During commit, Git stores message in .git/COMMIT_EDITMSG
	gitDir, err := r.GetGitDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(gitDir, "COMMIT_EDITMSG"), nil
}
