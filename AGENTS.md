# AGENTS.md

## Build, Lint, and Test Commands
- Format & tidy: `make tidy` (runs `go fmt ./...` and `go mod tidy -v`)
- Lint: `golangci-lint run` (or `make audit`)
- Test all: `make test` (runs `go test -tags assert -race -buildvcs ./...`)
- Test with coverage: `make test/cover`
- Run a single test: `go test -tags assert -race -buildvcs -run <TestName> ./...`
- Build: `make build`
- Run: `make run`
- Live reload: `make run/live` (uses `air`)

## Code Style Guidelines
- Use `go fmt` for formatting; run `make tidy` before commit.
- Group imports: stdlib, third-party, local.
- Use camelCase for variables, PascalCase for exported types/functions.
- Prefer explicit types; avoid `interface{}` unless necessary.
- Always handle errors; never ignore returned errors.
- Use `%w` or error wrapping for context.
- Keep functions short and focused.
- Use `context.Context` for request-scoped values.
- Avoid global state; prefer dependency injection.
- Write table-driven tests in `_test.go` files.
- Document exported functions/types with comments.
- Use `golangci-lint` for static analysis.

## Frameworks & References
- PicoCSS: https://picocss.com/ (lightweight, responsive CSS framework)
- Templ: https://templ.guide/ (Go HTML templating engine documentation)

No Cursor or Copilot rules found. Update this file if such rules are added.
