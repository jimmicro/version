# Contributing to version

Thank you for your interest in contributing.

## How to Contribute

### Reporting Bugs

Open a bug report with the GitHub issue template.

### Suggesting Features

Open a feature request with the GitHub issue template.

### Submitting Changes

1. Fork the repository
2. Create a branch: `git checkout -b feat/my-change`
3. Keep commits small and explicit
4. Run checks before opening a PR
5. Submit a Pull Request

## Commit Convention

This project uses small, descriptive commit messages in Conventional Commits style:

```text
<type>(<scope>): <subject>
```

Examples:

```text
feat(version): prefer build info version
fix(readme): clarify build metadata fallback
docs: add contributing guide
```

## Development

You can use plain Go commands:

```bash
go test ./...
go vet ./...
```

Or use `Taskfile` helpers:

```bash
task deps
task lint
task test
task build
```

## Code Style

- Format with `gofmt`
- Keep imports organized
- Pass `go vet` and `staticcheck`
- Add tests for behavior changes
