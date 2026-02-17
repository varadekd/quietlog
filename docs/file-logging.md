# File Logging & Rotation

quietlog always writes to stdout. File logging is opt-in and designed to complement stdout, not replace it.

---

## Enabling File Logging

Add to your `quietlog_config.json`:

```json
{
  "file_logging": true,
  "log_file_path": "./logs"
}
```

That's it. quietlog creates the directory if it doesn't exist and starts writing log files immediately.

---

## How Chunk Rotation Works

Rather than a single ever-growing log file, quietlog writes to **chunks** — a series of files with a shared timestamp prefix and an incrementing number:

```
logs/
  myservice_20260217_102412_chunk001.log
  myservice_20260217_102412_chunk002.log
  myservice_20260217_102412_chunk003.log
```

**The timestamp** in the filename is set once when the app starts and shared by all chunks from that run.

**A new chunk is opened** when the current chunk exceeds `max_file_size_mb`. The default is 10MB.

**No logs are lost during rotation** — the write and size check happen atomically under a lock. Every line is guaranteed to land in exactly one chunk.

---

## Configuring Chunk Size

```json
{
  "file_logging": true,
  "log_file_path": "./logs",
  "max_file_size_mb": 50
}
```

Accepts decimal values:

```json
{ "max_file_size_mb": 0.5 }
```

This sets the limit to 500KB, useful in memory-constrained environments or testing.

---

## Cleaning Old Logs on Startup

```json
{
  "file_logging": true,
  "log_file_path": "./logs",
  "clean_logs": true
}
```

When `clean_logs` is `true`, quietlog deletes all files matching `<app_name>_*_chunk*.log` in `log_file_path` before creating the first chunk of the new run.

> **Warning:** This permanently deletes old log files on every startup. Only use this in development environments or short-lived processes where logs are already shipped elsewhere before shutdown.

---

## What Happens When the File Path Is Invalid

If quietlog cannot open the log file (permission denied, path cannot be created), it logs an error to stderr and continues writing to stdout only. Your application never crashes.

```
quietlog: failed to open log file /var/log/myservice/...: permission denied
```

---

## Recommended Production Setup

```json
{
  "app_name": "myservice",
  "file_logging": true,
  "log_file_path": "./logs",
  "max_file_size_mb": 100,
  "clean_logs": false
}
```

Manage log retention externally using:

- **logrotate** on Linux
- **Fluentd / Filebeat** to ship logs to a central store
- **systemd journal** if running as a systemd service

quietlog intentionally does not implement log compression, time-based rotation, or remote shipping. These are solved problems with better dedicated tools.

---

## Stdout vs File

| | Stdout | File |
|--|--------|------|
| Always on | ✅ | ❌ opt-in |
| Color support | ✅ | ❌ plain text only |
| Chunk rotation | N/A | ✅ |
| Requires config | ❌ | ✅ |

Both outputs receive the same log lines. Color ANSI codes are written to stdout only — file output is always plain text regardless of `log_color` setting.
