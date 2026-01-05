package linter

import "strings"

// Validate validates a commit message against rules
func Validate(message string) *ValidationResult {
	result := &ValidationResult{
		IsValid:    true,
		Score:      100,
		Violations: []Violation{},
	}

	// Parse the commit message
	commit := ParseCommitMessage(message)

	// If we can't parse it at all, it's invalid
	if commit.Type == "" && commit.Description == "" {
		result.Violations = append(result.Violations, Violation{
			Rule:    "parse-failed",
			Message: "Commit message doesn't follow Conventional Commits format",
			Level:   "error",
		})
		result.IsValid = false
		result.Score = 0
		return result
	}

	// Apply all rules
	rules := DefaultRules()
	errorCount := 0
	warningCount := 0

	for _, rule := range rules {
		if !rule.Check(commit) {
			violation := Violation{
				Rule:    rule.Name,
				Message: rule.Message,
				Level:   rule.Level,
			}
			result.Violations = append(result.Violations, violation)

			if rule.Level == "error" {
				result.IsValid = false
				errorCount++
			} else {
				warningCount++
			}
		}
	}

	// Calculate score (simple formula)
	if errorCount > 0 {
		result.Score = 0
	} else {
		// Start at 100, deduct for warnings
		result.Score = 100 - (warningCount * 10)
		if result.Score < 0 {
			result.Score = 0
		}
	}

	// Generate suggestions
	result.Suggestions = generateSuggestions(commit, result.Violations)

	return result
}

func generateSuggestions(commit *CommitMessage, violations []Violation) []string {
	suggestions := []string{}

	hasTypeIssue := false
	hasFormatIssue := false

	for _, v := range violations {
		switch v.Rule {
		case "type-required", "type-enum":
			hasTypeIssue = true
		case "parse-failed":
			hasFormatIssue = true
		case "description-min-length":
			suggestions = append(suggestions, "Make the description more descriptive")
		case "imperative-mood":
			suggestions = append(suggestions, "Start with a verb like 'add', 'fix', 'update', 'remove'")
		}
	}

	if hasTypeIssue {
		suggestions = append(suggestions,
			"Start with a valid type: feat:, fix:, docs:, style:, refactor:, test:, chore:")
	}

	if hasFormatIssue {
		suggestions = append(suggestions,
			"Use format: type(scope): description\ne.g., feat(auth): add login functionality")
	}

	// Example suggestion
	if commit.Type != "" && commit.Description != "" {
		example := "Example: "
		if commit.Scope != "" {
			example += commit.Type + "(" + commit.Scope + "): "
		} else {
			example += commit.Type + ": "
		}

		// Improve the description
		words := strings.Fields(commit.Description)
		if len(words) > 0 {
			// Suggest imperative mood
			firstWord := strings.ToLower(words[0])
			if strings.HasSuffix(firstWord, "ed") {
				newFirst := strings.TrimSuffix(firstWord, "ed")
				example += newFirst + " " + strings.Join(words[1:], " ")
			} else {
				example += commit.Description
			}
		}
		suggestions = append(suggestions, example)
	}

	return suggestions
}
