# Repository Guidelines

## Project Structure & Module Organization
`cmd/mailer/` contains the CLI entry point and release metadata files such as `VERSION.txt` and `CHANGES.txt`. `internal/update/` holds the self-update logic used by the `update` command. Packaging and deployment assets live in `pkgr/`, including sample TOML config and platform deploy scripts. Root-level build wiring is in `magefile.go`, `build.toml`, and `.github/workflows/go.yml`.

## Build, Test, and Development Commands
Use `go build ./cmd/mailer` to compile the CLI. Use `go test ./...` to run package tests; because this repo currently has no `_test.go` files, that command mainly acts as a compile check. If Go reports inconsistent vendoring, sync dependencies with `go mod vendor` before building or testing. Format touched files with `gofmt -w cmd/mailer/*.go internal/update/*.go`.

For release packaging, Mage is wired through `magefile.go` and `build.toml`. Common flows are `mage install` for a local install and `mage release` for packaging across the targets listed in `build.toml`.

## Coding Style & Naming Conventions
Follow standard Go style: tabs via `gofmt`, short package names, exported identifiers in `PascalCase`, internal helpers in `camelCase`. Keep CLI flag and config names lowercase with underscores, matching existing keys such as `smtp_server`, `smtp_port`, and `skip_update`. Prefer small functions and keep command wiring in `cmd/mailer/`, not in `internal/`.

## Testing Guidelines
Add tests as package-local `*_test.go` files next to the code they cover. Favor table-driven tests for flag parsing, config loading, and SMTP/update edge cases. Run `go test ./...` before opening a PR, and note any build or vendor issues if they block a clean run.

## Commit & Pull Request Guidelines
Recent history uses short, imperative, lowercase subjects like `update deps` and `update crypt, hash, mymg`. Keep commit titles brief and scoped to one change. PRs should include a concise summary, the commands you ran (`go build`, `go test`, `mage release`, as applicable), and any release-file updates in `cmd/mailer/`. If a PR changes CLI behavior, include an example invocation or help text snippet instead of screenshots.

## Security & Configuration Tips
Do not commit real SMTP credentials, proxy values, or artifactory secrets. Treat `pkgr/config.toml` as an example only, and keep any local credential files or override configs out of version control.
