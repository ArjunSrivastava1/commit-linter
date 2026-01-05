package git

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// Commit represents a Git commit
type Commit struct {
	Hash      string
	Author    string
	Date      string
	Message   string
	ShortHash string
}

// GetLastCommit returns the most recent commit
func (r *Repository) GetLastCommit() (*Commit, error) {
	commits, err := r.GetCommits(1)
	if err != nil {
		return nil, err
	}

	if len(commits) == 0 {
		return nil, fmt.Errorf("no commits found")
	}

	return commits[0], nil
}

// GetCommits returns a list of commits
func (r *Repository) GetCommits(limit int) ([]*Commit, error) {
	// Git log format: hash|author|date|message
	format := "%H|%an|%cd|%s"

	cmd := exec.Command("git", "log",
		"--pretty=format:"+format,
		"--date=short",
		"-n", strconv.Itoa(limit))
	cmd.Dir = r.Path

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commits: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var commits []*Commit

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 4)
		if len(parts) != 4 {
			continue
		}

		commit := &Commit{
			Hash:    parts[0],
			Author:  parts[1],
			Date:    parts[2],
			Message: parts[3],
		}

		// Get short hash
		shortHashCmd := exec.Command("git", "rev-parse", "--short", commit.Hash)
		shortHashCmd.Dir = r.Path
		if shortHash, err := shortHashCmd.Output(); err == nil {
			commit.ShortHash = strings.TrimSpace(string(shortHash))
		}

		commits = append(commits, commit)
	}

	return commits, nil
}

// GetCommitsInRange returns commits between two references
func (r *Repository) GetCommitsInRange(fromRef, toRef string) ([]*Commit, error) {
	format := "%H|%an|%cd|%s"

	cmd := exec.Command("git", "log",
		"--pretty=format:"+format,
		"--date=short",
		fromRef+".."+toRef)
	cmd.Dir = r.Path

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get commits in range: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var commits []*Commit

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "|", 4)
		if len(parts) != 4 {
			continue
		}

		commits = append(commits, &Commit{
			Hash:    parts[0],
			Author:  parts[1],
			Date:    parts[2],
			Message: parts[3],
		})
	}

	return commits, nil
}

// GetCurrentBranch returns the current branch name
func (r *Repository) GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = r.Path

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}

// GetCommitMessage gets the commit message for a specific hash
func (r *Repository) GetCommitMessage(hash string) (string, error) {
	cmd := exec.Command("git", "log", "--format=%B", "-n", "1", hash)
	cmd.Dir = r.Path

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get commit message: %v", err)
	}

	return strings.TrimSpace(string(output)), nil
}
