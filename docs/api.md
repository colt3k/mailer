# CLI API Reference

`mailer` exposes a command-line API, not an HTTP or RPC API. The interface is defined by the flags and commands in `cmd/mailer/mailer.go`.

## Commands

| Command | Purpose | Side effects |
| --- | --- | --- |
| `send` | Sends one email through the configured SMTP server. | Opens an SMTP connection, writes logs, sends a message, optionally uploads an attachment with the message. |
| `buildconfig` | Writes a starter `config_example.toml` file in the current directory. | Creates or overwrites `config_example.toml`. |
| `update` | Checks the hard-coded artifactory source for a newer packaged release. | Makes an HTTP reachability check and may download/apply an update through the updater library. |

## Global and Message Flags

| Flag | Env var | Meaning |
| --- | --- | --- |
| `-config` | `CONFIG_FILEPATH` | Loads values from a config file handled by the CLI framework. |
| `-debug` | n/a | Enables debug logging. |
| `-proxyhttp` | `HTTP_PROXY` | Sets the HTTP proxy used for outbound network access. |
| `-proxyhttps` | `HTTPS_PROXY` | Sets the HTTPS proxy used for outbound network access. |
| `-noproxy` | `NO_PROXY` | Sets proxy exclusions. |
| `-skip_update` | `SKIP_UPDATE` | Skips the startup update check behavior provided by the CLI framework. |
| `-log_dir` | `LOG_DIR` | Changes the base directory for `colt3k/mailer.log`. |
| `-from` | `FROM` | Sender address. Required for `send`. |
| `-to` | `TO` | Recipient address. Required for `send`. |
| `-cc` | `CC` | Optional CC address. |
| `-ccname` | `CCNAME` | Optional display name for the CC address. |
| `-subject` | `SUBJECT` | Message subject. Defaults to `Test Message`. |
| `-html` | `HTML` | Sends the body as `text/html` instead of `text/plain`. |
| `-message` | `MESSAGE` | Message body. Defaults to `Hello, test message`. |
| `-file` | `FILE` | Attaches a single file path. |
| `-smtp_server` | `SMTP_SERVER` | SMTP host. Required for `send`. |
| `-smtp_port` | `SMTP_PORT` | SMTP port. Defaults to `587`. |
| `-smtp_username` | `SMTP_USERNAME` | SMTP username. |
| `-smtp_password` | `SMTP_PASSWORD` | SMTP password. |

## Behavioral Notes

- Port `25` disables STARTTLS enforcement in the current implementation.
- Any port other than `25` sets `MandatoryStartTLS`.
- Missing `smtp_server`, `from`, or `to` values cause a fatal log and stop execution.
- Attachments are optional and are added only when `-file` is non-empty after trimming whitespace.

## Example Calls

```bash
./mailer buildconfig
./mailer send -config pkgr/config.toml -subject "nightly report" -message "done"
./mailer send -config pkgr/config.toml -html -file ./report.html -message '<b>done</b>'
./mailer update
```
