# Send email

`mailer` is a small Go CLI for sending SMTP mail, generating a starter configuration, and checking for packaged updates. The binary is centered on one workflow: collect SMTP settings from flags, environment variables, or a TOML config file, then submit a single message with an optional attachment.

## Quick Start

Build the binary from the repository root:

```bash
go build ./cmd/mailer
```

If Go reports inconsistent vendoring, refresh the vendor tree first:

```bash
go mod vendor
```

Generate a starter configuration in the current directory:

```bash
./mailer buildconfig
```

Send a plain-text test message:

```bash
./mailer send \
  -smtp_server smtp.example.net \
  -smtp_port 587 \
  -smtp_username user@domain.net \
  -smtp_password 'redacted' \
  -from user@domain.net \
  -to user2@domain.net \
  -subject 'Smoke test' \
  -message 'hello from mailer'
```

## Configuration

The repository ships a sample file at `pkgr/config.toml`, and `buildconfig` writes a `config_example.toml` scaffold using the current flag values. Runtime logs go to `$HOME/colt3k/mailer.log` by default and can be redirected with `-log_dir`.

For a deeper breakdown of the supported config keys, command surface, and operational flow, see:

- `docs/developer.md`
- `docs/api.md`
- `docs/schema.md`
- `docs/dataflow.md`
- `docs/operator-runbook.md`

## Project Layout

- `cmd/mailer/`: CLI entry point plus version and change tracking files.
- `internal/update/`: self-update integration against the configured artifactory endpoint.
- `pkgr/`: packaging assets, sample config, autocomplete script, and deploy helpers.
- `magefile.go` and `build.toml`: release automation entry points.

## Built-in Help Reference

```text
USAGE:
    mailer [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
  -debug, -d
        flag set to debug (default false)

  -config, -c  string
    CONFIG_FILEPATH    (environment var)
        config file path

  -proxyhttp  string
    HTTP_PROXY    (environment var)
        Sets http_proxy for network connections

  -proxyhttps  string
    HTTPS_PROXY    (environment var)
        Sets https_proxy for network connections

  -noproxy  string
    NO_PROXY    (environment var)
        Sets no_proxy for network connections

  -skip_update, -skip
    SKIP_UPDATE    (environment var)
        set flag to skip update check (default false)

  -log_dir, -ld  string
    LOG_DIR    (environment var)
        set logging directory (default $HOME)

  -from  string
    FROM    (environment var)
        Who is the email from `user@domain.com`

  -to  string
    TO    (environment var)
        Whom to send the email to `user@domain.com`

  -cc  string
    CC    (environment var)
        Whom to cc on the email `user@domain.com`

  -ccname  string
    CCNAME    (environment var)
        Name to cc on the email

  -subject, -s  string
    SUBJECT    (environment var)
        Subject of the Email (default Test Message)

  -html
    HTML    (environment var)
        Message body default plain set for html (default false)

  -message, -m  string
    MESSAGE    (environment var)
        Message body to send in either quoted plaintext or html (default Hello, test message)

  -file, -f  string
    FILE    (environment var)
        Attach a file `FILE`

  -smtp_server  string
    SMTP_SERVER    (environment var)
        SMTP Server (default localhost)

  -smtp_port  int
    SMTP_PORT    (environment var)
        SMTP Port (default 587)

  -smtp_username  string
    SMTP_USERNAME    (environment var)
        SMTP Username

  -smtp_password  string
    SMTP_PASSWORD    (environment var)
        SMTP Password


COMMANDS:
  update:       (check for updates)
  buildconfig:  (build generic configuration you can fill in)
  send:         (send email)
```
