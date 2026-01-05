package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"commit-linter/internal/formatter"
	"commit-linter/internal/git"
	"commit-linter/internal/linter"
)

func main() {
	var (
		filePath      string
		showHelp      bool
		installHook   bool
		uninstallHook bool
		checkHook     bool
		lastCommit    bool
		commitCount   int
		commitRange   string
		force         bool
	)

	flag.StringVar(&filePath, "file", "", "Validate commit message from file")
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&installHook, "install", false, "Install Git commit-msg hook")
	flag.BoolVar(&uninstallHook, "uninstall", false, "Uninstall Git commit-msg hook")
	flag.BoolVar(&checkHook, "check", false, "Check if hook is installed")
	flag.BoolVar(&lastCommit, "last", false, "Validate last commit")
	flag.IntVar(&commitCount, "count", 1, "Number of commits to validate (with --last)")
	flag.StringVar(&commitRange, "range", "", "Validate commits in range (e.g., HEAD~3..HEAD)")
	flag.BoolVar(&force, "force", false, "Force overwrite existing hook")

	flag.Parse()

	if showHelp {
		printHelp()
		return
	}

	// Initialize Git repository
	repo, err := git.NewRepository("")
	if err != nil && (installHook || uninstallHook || checkHook || lastCommit || commitRange != "") {
		fmt.Printf("❌ Not a Git repository: %v\n", err)
		os.Exit(1)
	}

	// Handle Git operations
	switch {
	case installHook:
		handleInstallHook(repo, force)
		return
	case uninstallHook:
		handleUninstallHook(repo)
		return
	case checkHook:
		handleCheckHook(repo)
		return
	case lastCommit:
		handleLastCommits(repo, commitCount)
		return
	case commitRange != "":
		handleCommitRange(repo, commitRange)
		return
	}

	// Original validation logic
	var message string

	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
		message = string(data)
	} else if len(flag.Args()) > 0 {
		message = flag.Arg(0)
	} else {
		// Try to read from Git commit message file
		if repo != nil {
			commitMsgFile, err := repo.GetCommitMessageFilePath()
			if err == nil {
				data, err := os.ReadFile(commitMsgFile)
				if err == nil {
					message = string(data)
				}
			}
		}

		if message == "" {
			fmt.Println("Error: No commit message provided")
			fmt.Println("\nUsage:")
			fmt.Println("  commit-lint \"feat(auth): add login\"")
			fmt.Println("  commit-lint --file .git/COMMIT_EDITMSG")
			fmt.Println("  commit-lint --install  # Install Git hook")
			fmt.Println("  commit-lint --last     # Validate last commit")
			os.Exit(1)
		}
	}

	// Validate the message
	result := linter.Validate(message)

	// Print results
	formatter.PrintValidationResult(message, result)

	// Exit with appropriate code
	if !result.IsValid {
		os.Exit(1)
	}
}

func handleInstallHook(repo *git.Repository, force bool) {
	fmt.Println("🔧 Installing Git commit-msg hook...")
	fmt.Println()

	if err := repo.InstallCommitMsgHook(force); err != nil {
		fmt.Printf("❌ Failed to install hook: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("✅ Hook installed successfully!")
	fmt.Println()
	fmt.Println("From now on, all commit messages will be automatically validated.")
	fmt.Println()
	fmt.Println("To test the hook:")
	fmt.Println("  git commit -m \"test message\"  # Should fail validation")
	fmt.Println("  git commit -m \"feat: add feature\"  # Should pass")
}

func handleUninstallHook(repo *git.Repository) {
	fmt.Println("🔧 Uninstalling Git commit-msg hook...")

	if err := repo.UninstallCommitMsgHook(); err != nil {
		fmt.Printf("❌ Failed to uninstall hook: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Hook uninstalled successfully!")
}

func handleCheckHook(repo *git.Repository) {
	installed, err := repo.IsHookInstalled()
	if err != nil {
		fmt.Printf("❌ Error checking hook: %v\n", err)
		os.Exit(1)
	}

	if installed {
		fmt.Println("✅ commit-lint hook is installed")
	} else {
		fmt.Println("❌ commit-lint hook is NOT installed")
		fmt.Println("   Install with: commit-lint --install")
	}
}

func handleLastCommits(repo *git.Repository, count int) {
	fmt.Printf("📊 Validating last %d commit(s)...\n\n", count)

	commits, err := repo.GetCommits(count)
	if err != nil {
		fmt.Printf("❌ Failed to get commits: %v\n", err)
		os.Exit(1)
	}

	allValid := true
	totalScore := 0

	for i, commit := range commits {
		fmt.Printf("[%d/%d] Commit: %s\n", i+1, len(commits), commit.ShortHash)
		fmt.Printf("     Message: %s\n", commit.Message)
		fmt.Printf("     Author:  %s\n", commit.Author)
		fmt.Printf("     Date:    %s\n", commit.Date)

		result := linter.Validate(commit.Message)

		if result.IsValid {
			fmt.Printf("     Status:  ✅ Valid (%d/100)\n", result.Score)
		} else {
			fmt.Printf("     Status:  ❌ Invalid (%d/100)\n", result.Score)
			allValid = false
		}

		totalScore += result.Score

		if i < len(commits)-1 {
			fmt.Println()
		}
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 50))

	if allValid {
		avgScore := totalScore / len(commits)
		fmt.Printf("✅ All %d commits are valid! (Average score: %d/100)\n",
			len(commits), avgScore)
	} else {
		fmt.Printf("❌ Some commits failed validation\n")
		os.Exit(1)
	}
}

func handleCommitRange(repo *git.Repository, rangeStr string) {
	fmt.Printf("📊 Validating commits in range: %s...\n\n", rangeStr)

	// Parse range (e.g., "HEAD~3..HEAD")
	parts := strings.Split(rangeStr, "..")
	if len(parts) != 2 {
		fmt.Println("❌ Invalid range format. Use: from..to")
		fmt.Println("   Example: HEAD~3..HEAD")
		os.Exit(1)
	}

	commits, err := repo.GetCommitsInRange(parts[0], parts[1])
	if err != nil {
		fmt.Printf("❌ Failed to get commits: %v\n", err)
		os.Exit(1)
	}

	if len(commits) == 0 {
		fmt.Println("ℹ️  No commits found in the specified range")
		return
	}

	allValid := true
	validCount := 0
	totalScore := 0

	for _, commit := range commits {
		result := linter.Validate(commit.Message)

		status := "✅"
		if !result.IsValid {
			status = "❌"
			allValid = false
		} else {
			validCount++
		}

		totalScore += result.Score

		fmt.Printf("%s [%s] %s - %s\n",
			status,
			commit.ShortHash,
			commit.Date,
			commit.Message)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 50))

	avgScore := 0
	if len(commits) > 0 {
		avgScore = totalScore / len(commits)
	}

	fmt.Printf("📈 SUMMARY (%d commits):\n", len(commits))
	fmt.Printf("   Valid: %d/%d (%.0f%%)\n",
		validCount, len(commits),
		float64(validCount)/float64(len(commits))*100)
	fmt.Printf("   Average score: %d/100\n", avgScore)

	if !allValid {
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`
🌳 COMMIT MESSAGE LINTER v2.0
─────────────────────────────

Validate commit messages against Conventional Commits standard.

BASIC USAGE:
  commit-lint "feat(auth): add login functionality"
  commit-lint --file .git/COMMIT_EDITMSG

GIT INTEGRATION:
  commit-lint --install              Install Git commit-msg hook
  commit-lint --install --force      Force install (overwrite)
  commit-lint --uninstall            Uninstall hook
  commit-lint --check                Check if hook is installed

VALIDATE HISTORY:
  commit-lint --last                 Validate last commit
  commit-lint --last --count 5       Validate last 5 commits
  commit-lint --range HEAD~3..HEAD   Validate commits in range

EXAMPLES:
  • feat: add new feature
  • fix: resolve bug
  • docs: update documentation
  • style: code formatting
  • refactor: code restructuring
  • test: add tests
  • chore: maintenance tasks

GIT HOOK FEATURES:
  • Automatically validates every commit
  • Prevents invalid commits
  • Shows helpful suggestions
  • Can be bypassed with --no-verify

Learn more: https://www.conventionalcommits.org/`)
}
