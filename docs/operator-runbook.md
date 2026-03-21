# Operator Runbook

## Scope

Use this runbook when operating the `mailer` CLI manually or from a scheduler. The tool sends one SMTP message per invocation and can also generate a starter config or check for packaged updates.

## Prerequisites

- A built or installed `mailer` binary.
- SMTP host, port, username, and password for the target mail server.
- A sender address and at least one recipient address.
- File-system access to the attachment path when using `-file`.

## Standard Procedures

### Generate a config template

```bash
./mailer buildconfig
```

This writes `config_example.toml` in the current directory. Move secrets into an operator-managed config outside the repository and pass it with `-config`.

### Send a smoke-test email

```bash
./mailer send -config /path/to/mailer.toml -subject "mailer smoke test" -message "ok"
```

### Send an HTML email with an attachment

```bash
./mailer send \
  -config /path/to/mailer.toml \
  -html \
  -file /path/to/report.html \
  -subject "daily report" \
  -message '<html><body><h1>done</h1></body></html>'
```

### Check for updates

```bash
./mailer update
```

The update path depends on the artifactory endpoint compiled into `internal/update/update.go`.

## Logs and Observability

- Default log file: `$HOME/colt3k/mailer.log`
- Custom log root: `-log_dir /path/to/logs`
- Enable verbose troubleshooting with `-debug`

Look for fatal log messages first; the CLI stops immediately when required SMTP inputs are missing or the SMTP send fails.

## Troubleshooting

| Symptom | Likely cause | Response |
| --- | --- | --- |
| `smtp server required` | Missing `smtp_server` value | Confirm flag, env var, or config key. |
| `from required` or `to required` | Missing address fields | Validate the config file and overrides. |
| TLS or authentication failure | Wrong port, credentials, or server policy | Recheck `smtp_port`, username/password, and whether the server expects STARTTLS or port `25`. |
| Attachment missing | Bad `-file` path | Verify the file exists and is readable by the invoking user. |
| No update found or update check appears idle | Endpoint unreachable | Validate network access to the compiled artifactory host. |

## Rollback

Rollback is manual. Reinstall the previous known-good binary or package artifact, then rerun the smoke-test command above to verify SMTP connectivity and logging.
