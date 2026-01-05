package linter

import (
	"strings"
)

// Rule defines a validation rule
type Rule struct {
	Name    string
	Check   func(*CommitMessage) bool
	Message string
	Level   string // "error", "warning"
}

// Violation represents a rule violation
type Violation struct {
	Rule    string
	Message string
	Level   string
}

// ValidationResult contains the validation outcome
type ValidationResult struct {
	IsValid     bool
	Score       int
	Violations  []Violation
	Suggestions []string
}

// DefaultRules returns the standard validation rules
func DefaultRules() []Rule {
	return []Rule{
		{
			Name: "type-required",
			Check: func(msg *CommitMessage) bool {
				return msg.Type != ""
			},
			Message: "Commit type is required (feat, fix, docs, etc.)",
			Level:   "error",
		},
		{
			Name: "type-case",
			Check: func(msg *CommitMessage) bool {
				return msg.Type == strings.ToLower(msg.Type)
			},
			Message: "Type must be lowercase",
			Level:   "error",
		},
		{
			Name: "type-enum",
			Check: func(msg *CommitMessage) bool {
				validTypes := []string{
					"feat", "fix", "docs", "style",
					"refactor", "test", "chore", "perf",
				}
				for _, validType := range validTypes {
					if msg.Type == validType {
						return true
					}
				}
				return false
			},
			Message: "Type must be one of: feat, fix, docs, style, refactor, test, chore, perf",
			Level:   "error",
		},
		{
			Name: "description-required",
			Check: func(msg *CommitMessage) bool {
				return strings.TrimSpace(msg.Description) != ""
			},
			Message: "Description is required",
			Level:   "error",
		},
		{
			Name: "description-min-length",
			Check: func(msg *CommitMessage) bool {
				return len(msg.Description) >= 10
			},
			Message: "Description must be at least 10 characters",
			Level:   "warning",
		},
		{
			Name: "description-max-length",
			Check: func(msg *CommitMessage) bool {
				return len(msg.Description) <= 72
			},
			Message: "Description should not exceed 72 characters (GitHub truncates)",
			Level:   "warning",
		},
		{
			Name: "no-period",
			Check: func(msg *CommitMessage) bool {
				return !strings.HasSuffix(msg.Description, ".")
			},
			Message: "Description should not end with a period",
			Level:   "warning",
		},
		{
			Name: "imperative-mood",
			Check: func(msg *CommitMessage) bool {
				if msg.Description == "" {
					return true
				}
				// Simple check for imperative mood (starts with verb)
				firstWord := strings.Fields(msg.Description)[0]
				// Common non-imperative words
				nonImperative := map[string]bool{
					"added": true, "adds": true, "adding": true,
					"fixed": true, "fixes": true, "fixing": true,
					"updated": true, "updates": true, "updating": true,
					"changed": true, "changes": true, "changing": true,
				}
				return !nonImperative[strings.ToLower(firstWord)]
			},
			Message: "Use imperative mood (e.g., 'add' not 'added', 'fix' not 'fixed')",
			Level:   "warning",
		},
	}
}
