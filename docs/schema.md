# Configuration Schema

The CLI framework can hydrate runtime values from flags, environment variables, and an optional TOML file passed through `-config`. When using a config file, keep keys aligned with the flag names shown below.

## Core SMTP Fields

| Key | Type | Required for `send` | Default | Notes |
| --- | --- | --- | --- | --- |
| `smtp_server` | string | Yes | `localhost` | SMTP host name or IP. |
| `smtp_port` | integer | No | `587` | Uses STARTTLS unless the value is `25`. |
| `smtp_username` | string | Usually | none | Needed when the server requires authentication. |
| `smtp_password` | string | Usually | none | Store outside version control. |
| `from` | string | Yes | none | Sender email address. |
| `to` | string | Yes | none | Primary recipient address. |

## Optional Message Fields

| Key | Type | Default | Notes |
| --- | --- | --- | --- |
| `cc` | string | none | Optional CC address. |
| `ccname` | string | none | Display name used with `cc`. |
| `subject` | string | `Test Message` | Subject line. |
| `html` | boolean | `false` | Switches the body MIME type to HTML. |
| `message` | string | `Hello, test message` | Message body content. |
| `file` | string | none | Path to one attachment. |
| `log_dir` | string | `$HOME` | Base directory for `colt3k/mailer.log`. |
| `skip_update` | boolean | `false` | Skip update checks handled by the CLI framework. |

## Example TOML

```toml
smtp_server = "smtp.example.net"
smtp_port = 587
smtp_username = "user@domain.net"
smtp_password = "use-a-secret-manager"
from = "user@domain.net"
to = "user2@domain.net"
subject = "Smoke test"
message = "hello from mailer"
html = false
```

## Generated and Sample Files

- `pkgr/config.toml` is the repository sample.
- `mailer buildconfig` writes `config_example.toml` in the current working directory.
- The generated file includes the required SMTP keys plus `from` and `to`, using live flag values when available and placeholder addresses otherwise.
