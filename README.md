<h1>
  <br>
  <img src="https://raw.githubusercontent.com/ArjunSrivastava1/commit-linter/main/assets/icon.svg" alt="commit-linter" width="100">
  <br>
</h1>

<h4>A 'git commit' messages tool for ease of coding</h4>

<p>
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-GPL%20v2-blue.svg" alt="License"></a>
  <a href="https://goreportcard.com/report/github.com/ArjunSrivastava1/commit-linter"><img src="https://goreportcard.com/badge/github.com/ArjunSrivastava1/commit-linter" alt="Go Report Card"></a>
</p>

<p>
  <a href="#-about">About</a> •
  <a href="#-features">Features</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-team-configuration">Configuration</a> •
  <a href="#-contributing">Contributing</a>
</p>

<p>
  <img src="https://raw.githubusercontent.com/ArjunSrivastava1/commit-linter/main/assets/demo.gif" alt="Demo" width="600">
</p>

## About

commit-linter is a command line interface (CLI) tool written in GO that enforces standards for git commits with elegance and precision. It validates commit messages in real-time, provides actionable feedback, and seamlessly integrates with Git to maintain consistency across your project's history. 

## ✨ Features

| Category | Features |
|----------|----------|
| **🎯 Validation** | Conventional Commits • Type/Scope checking • Length rules • Imperative mood |
| **🔧 Git Integration** | Auto-hook install • Historical validation • Pre-commit blocking • Team consistency |
| **🎨 Beautiful Output** | Color-coded feedback • Actionable suggestions • Validation scoring • Clean tables |

## 🚀 Quick Start

### 📦 Installation
```bash
# One-liner install
go install github.com/ArjunSrivastava1/commit-linter
```

## 🎯 Basic Usage
```bash
# Validate a message
commit-lint "feat(auth): add login"

# Install Git hook (auto-validates all commits)
commit-lint --install

# Check last 3 commits
commit-lint --last --count 3
```

### 🛠️ Usage

### Git Hook Automation
```bash
# Install once, validate forever
commit-lint --install

# Now try committing:
git commit -m "bad message"   # ❌ Blocked
git commit -m "feat: add feature"  # ✅ Allowed
```

### CI/CD Integration
```yaml
# GitHub Actions
- name: Validate Commits
  run: commit-lint --last-commit
```

### Validate Git History
```bash
# Check last commit
commit-lint --last-commit

# Check range of commits
commit-lint --range HEAD~5..HEAD

# Check all commits in branch
commit-lint --branch main
```

## ⚙️ Team Configuration
```yaml
# .commitlint.yml
rules:
  type-enum: [feat, fix, docs, style, refactor, test, chore]
  subject-min-length: 10
  subject-max-length: 72
```

## 🤝 Contributing

1. Fork & clone
2. Create feature branch
3. Commit with Conventional Commits
4. Push & open PR

See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## 📄 License

GPL v2.0 - See [LICENSE](LICENSE)

---

<p align="center">
  Built with ❤️ by <a href="https://github.com/ArjunSrivastava1">Arjun Srivastava</a>
</p>

<p align="center">
  <a href="https://github.com/ArjunSrivastava1/commit-linter/issues">Report Bug</a> • 
  <a href="https://github.com/ArjunSrivastava1/commit-linter/issues">Request Feature</a> •
  <a href="https://github.com/ArjunSrivastava1/enva">enva</a> •
  <a href="https://github.com/ArjunSrivastava1/port-scanner">port-scanner</a>
</p>

<p align="center">
