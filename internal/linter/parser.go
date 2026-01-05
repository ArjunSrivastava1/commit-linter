package linter

import (
	"regexp"
	"strings"
)

// CommitMessage represents a parsed commit message
type CommitMessage struct {
	Raw         string
	Type        string
	Scope       string
	Description string
	Body        string
	IsBreaking  bool
}

// ParseCommitMessage parses a commit message using Conventional Commits format
func ParseCommitMessage(message string) *CommitMessage {
	msg := &CommitMessage{Raw: message}

	// Remove leading/trailing whitespace
	message = strings.TrimSpace(message)

	// Check for breaking change indicator
	if strings.Contains(message, "!:") {
		msg.IsBreaking = true
		message = strings.Replace(message, "!:", ":", 1)
	}

	// Regex for conventional commits: type(scope): description
	// Example: feat(auth): add login functionality
	re := regexp.MustCompile(`^(\w+)(?:\(([^)]+)\))?!?: (.+)$`)

	matches := re.FindStringSubmatch(message)
	if matches != nil {
		msg.Type = matches[1]
		msg.Scope = matches[2]
		msg.Description = matches[3]

		// Check for body (separated by two newlines)
		parts := strings.SplitN(message, "\n\n", 2)
		if len(parts) > 1 {
			msg.Body = parts[1]
		}
	} else {
		// Fallback: try to parse just type: description
		reSimple := regexp.MustCompile(`^(\w+): (.+)$`)
		simpleMatches := reSimple.FindStringSubmatch(message)
		if simpleMatches != nil {
			msg.Type = simpleMatches[1]
			msg.Description = simpleMatches[2]
		}
	}

	return msg
}
