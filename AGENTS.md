# Repository Guidelines

## Project Structure & Module Organization
- `main.go` wires the CLI root command and registers subcommands.
- `encrypt/` and `keys/` contain Cobra command implementations for encryption and key management.
- `common/` holds shared helpers (e.g., encryption utilities).
- `go.mod`/`go.sum` define dependencies; `Makefile` provides a simple build target.

## Build, Test, and Development Commands
- `make build`: Builds the `tools` binary into the repo root (removes any existing binary first).
- `go build -o ./tools ./main.go`: Equivalent manual build.
- `go test ./...`: Runs Go tests for all packages (currently there are no test files, so it should be quick).

## Coding Style & Naming Conventions
- Language: Go 1.25.6.
- Indentation: tabs, as per `gofmt` defaults.
- Naming: Go idioms (CamelCase for exported identifiers, lowerCamelCase for unexported).
- Formatting: run `gofmt -w .` on modified `.go` files before commits.

## Testing Guidelines
- Framework: Go’s standard `testing` package (no external framework configured).
- Naming: `*_test.go` files with `TestXxx` functions.
- Coverage: no explicit coverage targets; add tests for new command behaviors and crypto helpers when you change them.

## Commit & Pull Request Guidelines
- Commit history uses short, imperative summaries (often in Chinese) without scopes or prefixes.
  - Example: “调整目录结构” or “支持加密自定义输入”.
- Keep commits small and focused; one logical change per commit.
- PRs should include:
  - A brief description of the change and affected commands.
  - Test results (e.g., `go test ./...`).
  - Any CLI usage updates or example commands if behavior changes.

## Security & Configuration Tips
- The `keys/` package handles key material. Avoid logging secrets and be cautious with sample data in tests or docs.
- If you add config files, keep defaults out of the repo and document expected paths (e.g., `$HOME/.tools.yaml`).
