# Contributing to DILEX

DILEX is in early development - Before writing anything, open a discussion issue first.

---

## Scope

V1 does one thing: guaranteed, priority-ordered data delivery between two nodes across an unreliable connection. Contributions that pull toward V2+ scope will be deferred.

---

## Reporting Issues

Include: description of the behavior, steps to reproduce, `dilex version` or commit SHA, OS/arch/Go version, and relevant logs.

**Security vulnerabilities:** Email [security@relumic.com](mailto:security@relumic.com) - do not file as public issues.

---

## Development Setup

**Requirements:** Go 1.26+, `protoc` + `protoc-gen-go`, SQLite (system library).

```bash
git clone https://github.com/relumic/dilex.git
cd dilex
go build ./...
go test -race ./...                         # race detector is mandatory
go test -race -tags integration ./...       # integration tests
gofmt -l .
go vet ./...
golangci-lint run
```

---

## Code Standards

- Clarity over cleverness. Export only what must be exported.
- All exported symbols need doc comments. Errors are returned, not swallowed.
- No global state, no `panic` outside initialization, no commented-out code in PRs.
- Race detector must pass. `gofmt`-clean. Flaky tests will be rejected.

**Commits:**
```
<type>: <short summary in imperative mood>
```
Types: `feat`, `fix`, `docs`, `test`, `refactor`, `chore`. Subject line under 72 characters.

---

## Pull Requests

Fork → branch from `main` → make changes → `go test -race ./...` + `gofmt -l .` → open PR against `main`. One logical change per PR. We aim to review within one week.

---

## License

Contributions are licensed under [Apache 2.0](LICENSE).
