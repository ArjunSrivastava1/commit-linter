<h1 align="center">
  <br>
  <img src="https://raw.githubusercontent.com/ArjunSrivastava1/commit-linter/main/assets/icon.svg" alt="commit-linter" width="100">
  <br>
  🌳 commit-linter
  <br>
</h1>

<h4 align="center">Automate Conventional Commits with beautiful feedback & Git hooks.</h4>

<p align="center">
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go&logoColor=white" alt="Go Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/License-GPL%20v2-blue.svg" alt="License"></a>
  <a href="https://github.com/ArjunSrivastava1/commit-linter/releases"><img src="https://img.shields.io/github/v/release/ArjunSrivastava1/commit-linter" alt="Release"></a>
  <a href="https://goreportcard.com/report/github.com/ArjunSrivastava1/commit-linter"><img src="https://goreportcard.com/badge/github.com/ArjunSrivastava1/commit-linter" alt="Go Report Card"></a>
</p>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-demo">Demo</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-project-structure">Structure</a>
</p>

<p align="center">
  <img src="https://raw.githubusercontent.com/ArjunSrivastava1/commit-linter/main/assets/demo.gif" alt="Demo" width="600">
</p>

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
go install github.com/ArjunSrivastava1/commit-linter/cmd/commit-lint@latest
```

### 🎯 Basic Usage
```bash
# Validate a message
commit-lint "feat(auth): add login"

# Install Git hook (auto-validates all commits)
commit-lint --install

# Check last 3 commits
commit-lint --last --count 3
```

## 📊 Demo

### ✅ Success Case
```bash
$ commit-lint "feat(auth): add JWT authentication"

┌─────────────────────────────────────────────┐
│ 🌳 COMMIT MESSAGE VALIDATION - 95/100       |
├─────────────────────────────────────────────┤
│ Type:    feat (New feature)        ✅      │
│ Scope:   auth                      ✅      │
│ Length:  24 chars                  ✅      │
└─────────────────────────────────────────────┘
```

### ❌ Failure Case
```bash
$ commit-lint "fix bug"

┌─────────────────────────────────────────────┐
│ ❌ COMMIT VALIDATION FAILED - 0/100         │
├─────────────────────────────────────────────┤
│ Type:    Missing                   ❌      │
│ Length:  7 chars (min 10)          ❌      │
├─────────────────────────────────────────────┤
│ 💡 SUGGESTIONS:                            │
│ • Use: fix: resolve authentication bug     │
│ • Or:  fix(api): handle null response      │
└─────────────────────────────────────────────┘
```

## 🛠️ Usage

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

### Team Configuration
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
  Made with ❤️ by <a href="https://github.com/ArjunSrivastava1">Arjun Srivastava</a>
</p>
