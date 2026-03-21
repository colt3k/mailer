# Developer Guide

## Purpose

This repository builds a single CLI binary, `mailer`, from `cmd/mailer/`. The main command wiring lives in `cmd/mailer/mailer.go`, while `internal/update/update.go` contains the self-update path. Packaging metadata and release helpers sit in `build.toml`, `magefile.go`, and `pkgr/`.

## Local Tooling

- Use the Go toolchain declared in `go.mod` as the source of truth.
- Install Mage if you need packaging or release workflows.
- Keep `vendor/` aligned with `go.mod` when dependencies change.

## Common Commands

```bash
go build ./cmd/mailer
go test ./...
gofmt -w cmd/mailer/*.go internal/update/*.go
go mod vendor
mage install
mage release
```

`go build` and `go test` are the quickest sanity checks. `go mod vendor` is required whenever dependency bumps leave `vendor/modules.txt` behind. `mage install` and `mage release` delegate to the shared `github.com/colt3k/utils/mymg` targets referenced by `magefile.go`.

## Development Notes

- Keep CLI flag names, environment variables, and config keys aligned. The current code uses names like `smtp_server`, `skip_update`, and `log_dir`.
- Prefer small command handlers. `setupFlags()` should remain the place where flags and commands are declared, while message delivery logic stays in `run()`.
- Treat `pkgr/config.toml` as a sample, not a secret store. Use local overrides outside version control for real credentials.

## Known Repository Caveats

- The repository includes a vendored dependency tree. If compilation starts failing with "inconsistent vendoring", rerun `go mod vendor`.
- Dependency updates around `github.com/colt3k/utils/updater` should be validated carefully. The updater API has changed across versions and can break `internal/update/update.go`.
- `.github/workflows/go.yml` is a legacy workflow. Check `go.mod` before assuming the workflow's toolchain version is still correct.
