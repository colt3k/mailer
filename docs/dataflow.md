# Runtime Dataflow

This document describes how `mailer` moves from process start to SMTP delivery or update handling.

## High-Level Flow

```text
process start
  -> init()
     -> create process lock
     -> open default file + console logging
  -> main()
     -> setupFlags()
        -> register metadata, flags, and commands
        -> parse args / env / config file
        -> dispatch command action
           -> send -> run()
           -> buildconfig -> buildConfig()
           -> update -> update.CheckUpdate()
  -> unlock process lock
  -> exit
```

## Send Command Flow

1. Validate that `smtp_server`, `from`, and `to` are present.
2. Build a `mail.Message`.
3. Add `From`, `To`, optional `Cc`, and `Subject` headers.
4. Select `text/plain` or `text/html` body based on the `html` flag.
5. Attach the file named by `file` when present.
6. Create a `mail.Dialer` using the supplied SMTP host, port, username, and password.
7. Apply TLS policy:
   - Port `25`: no STARTTLS enforcement.
   - Any other port: `MandatoryStartTLS`.
8. Send the message and log success or a fatal error.

## Buildconfig Flow

`buildConfig()` builds a small TOML document in memory, then writes `config_example.toml` to the current working directory. It preserves current runtime values when they are set and falls back to placeholders for `from` and `to`.

## Update Flow

`update.CheckUpdate()` packages the current version metadata, probes the configured artifactory base URL, and only registers an update source when that endpoint is reachable. It then passes the resulting connection list to the shared updater library.

## Operational Side Effects

- Lock file: acquired during startup to avoid overlapping runs.
- Logs: written to `$HOME/colt3k/mailer.log` unless `-log_dir` changes the base directory.
- Network:
  - SMTP traffic for `send`
  - HTTP reachability and updater traffic for `update`
- Files:
  - Read: attachment path, optional config file
  - Write: log file, `config_example.toml`
