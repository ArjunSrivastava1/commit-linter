package formatter

import (
	"commit-linter/internal/linter"
	"fmt"
	"strings"
)

// Color constants (ANSI escape codes)
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[90m"
	Bold   = "\033[1m"
)

// PrintValidationResult prints the validation result beautifully
func PrintValidationResult(message string, result *linter.ValidationResult) {
	commit := linter.ParseCommitMessage(message)

	fmt.Println()

	if result.IsValid {
		fmt.Println(Green + Bold + "✅ COMMIT MESSAGE VALIDATION" + Reset)
	} else {
		fmt.Println(Red + Bold + "❌ COMMIT MESSAGE VALIDATION FAILED" + Reset)
	}

	fmt.Println(strings.Repeat("─", 50))
	fmt.Println()

	// Message preview
	fmt.Println(Cyan + Bold + "📝 MESSAGE:" + Reset)
	fmt.Println("  " + commit.Raw)
	fmt.Println()

	// Parsed components
	fmt.Println(Blue + Bold + "🔍 PARSED COMPONENTS:" + Reset)
	if commit.Type != "" {
		fmt.Printf("  Type:        %s%s%s\n", Green, commit.Type, Reset)
	} else {
		fmt.Printf("  Type:        %sMissing%s\n", Red, Reset)
	}

	if commit.Scope != "" {
		fmt.Printf("  Scope:       %s%s%s\n", Gray, commit.Scope, Reset)
	} else {
		fmt.Printf("  Scope:       %s(optional)%s\n", Gray, Reset)
	}

	if commit.Description != "" {
		fmt.Printf("  Description: %s%s%s\n", Gray, commit.Description, Reset)
		fmt.Printf("  Length:      %d chars\n", len(commit.Description))
	}

	if commit.IsBreaking {
		fmt.Printf("  Breaking:    %s⚠️  BREAKING CHANGE%s\n", Yellow, Reset)
	}

	fmt.Println()

	// Validation summary
	fmt.Println(Purple + Bold + "📊 VALIDATION SUMMARY:" + Reset)
	fmt.Printf("  Status:      ")
	if result.IsValid {
		fmt.Printf("%sPASS%s\n", Green, Reset)
	} else {
		fmt.Printf("%sFAIL%s\n", Red, Reset)
	}

	fmt.Printf("  Score:       %s%d/100%s\n",
		getScoreColor(result.Score), result.Score, Reset)

	// Count violations by level
	errors := 0
	warnings := 0
	for _, v := range result.Violations {
		if v.Level == "error" {
			errors++
		} else {
			warnings++
		}
	}

	if errors > 0 {
		fmt.Printf("  Errors:      %s%d%s\n", Red, errors, Reset)
	}
	if warnings > 0 {
		fmt.Printf("  Warnings:    %s%d%s\n", Yellow, warnings, Reset)
	}

	fmt.Println()

	// Show violations
	if len(result.Violations) > 0 {
		fmt.Println(Yellow + Bold + "⚠️  ISSUES FOUND:" + Reset)
		for _, v := range result.Violations {
			icon := "•"
			color := Gray
			if v.Level == "error" {
				icon = "❌"
				color = Red
			} else {
				icon = "⚠️"
				color = Yellow
			}
			fmt.Printf("  %s %s%s%s\n", icon, color, v.Message, Reset)
		}
		fmt.Println()
	}

	// Show suggestions
	if len(result.Suggestions) > 0 {
		fmt.Println(Green + Bold + "💡 SUGGESTIONS:" + Reset)
		for _, s := range result.Suggestions {
			fmt.Printf("  • %s\n", s)
		}
		fmt.Println()
	}

	// Examples
	if !result.IsValid {
		fmt.Println(Blue + Bold + "📚 VALID EXAMPLES:" + Reset)
		examples := []string{
			"feat(auth): add login functionality",
			"fix(api): resolve null pointer in user endpoint",
			"docs(readme): update installation instructions",
			"style(css): format button padding",
			"refactor(auth): simplify token validation",
			"test(login): add unit tests for authentication",
			"chore(deps): update dependencies",
		}

		for _, ex := range examples {
			fmt.Printf("  • %s\n", ex)
		}
		fmt.Println()
	}

	// Final status
	fmt.Println(strings.Repeat("─", 50))
	if result.IsValid {
		fmt.Println(Green + Bold + "✅ Ready to commit!" + Reset)
	} else {
		fmt.Println(Red + Bold + "❌ Please fix the issues above before committing." + Reset)
	}
	fmt.Println()
}

func getScoreColor(score int) string {
	if score >= 80 {
		return Green
	} else if score >= 60 {
		return Yellow
	}
	return Red
}
