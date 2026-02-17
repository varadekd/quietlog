# quietlog

A tiny, opinionated, stdlib-first logging library for Go.

quietlog focuses on **clarity over cleverness**. It avoids structured logging, reflection, and heavy abstractions — and instead provides readable, predictable logs suitable for both development and production environments.

---

## Why quietlog?

Most logging libraries in Go are built for scale, pipelines, or observability stacks.

quietlog is built for:

- Human-readable logs
- Small to medium services
- Internal tools & CLIs
- Long-running production apps that value stability
- Developers who prefer the Go standard library

If you need structured JSON logs, distributed tracing, or high-volume ingestion pipelines — quietlog is **not** the right tool.

---

## Features

- Zero framework dependencies
- Built on Go's `log` package
- Auto-initializes on first use — no setup required
- Always logs to stdout
- Optional file logging with automatic chunk rotation
- Configurable via `quietlog_config.json` — fully optional
- Custom log format, timezone, and timestamp
- Colored output support
- Debug level toggle — off by default
- Production-safe defaults
- Concurrent-safe

---

## Installation

```bash
go get github.com/varadekd/quietlog
```

---

## Quick Start

```go
package main

import "github.com/varadekd/quietlog"

func main() {
    quietlog.Info("server started")
    quietlog.Warn("disk usage above 80%")
    quietlog.Error("failed to connect to database")
}
```

No configuration needed. quietlog initializes itself on first use with safe defaults.

**Output:**
```
[APP] [INFO] : server started [17-Feb-2026 10:24:12 UTC]
[APP] [WARN] : disk usage above 80% [17-Feb-2026 10:24:12 UTC]
[APP] [ERROR] : failed to connect to database [17-Feb-2026 10:24:12 UTC]
```

---

## Log Levels

| Function | Description |
|----------|-------------|
| `quietlog.Info(msg)` | General information |
| `quietlog.Warn(msg)` | Something unexpected, not fatal |
| `quietlog.Error(msg)` | Something failed |
| `quietlog.Fatal(msg)` | Logs then panics with full stack trace |
| `quietlog.Debug(msg)` | Verbose output — off by default |

Formatted variants are available for all levels:

```go
quietlog.Infof("user %s logged in from %s", username, ip)
quietlog.Warnf("disk usage at %d%%", usage)
quietlog.Errorf("failed to connect to %s:%d", host, port)
quietlog.Fatalf("could not bind to port %d", port)
quietlog.Debugf("cache hit rate: %.2f%%", rate)
```

> **Note:** `Fatal` and `Fatalf` log the message and then call `panic`. This is intentional — if something is fatal, you want a full stack trace, not a silent exit.

---

## Configuration

quietlog works with zero configuration. When you want to customize behavior, drop a `quietlog_config.json` file in your project root — quietlog will find it automatically on startup.

### Auto-discovery

quietlog looks for `quietlog_config.json` in the current directory on startup. If found, it loads it silently. If not found, it uses defaults silently. Your app never breaks either way.

### Custom Config Path

If your config file is not at the project root:

```go
func main() {
    quietlog.Config("./configs/quietlog_config.json")
    quietlog.Info("server started")
}
```

Call `Config` once before your first log statement.

### Config Schema

```json
{
  "app_name": "myservice",
  "debug_level": false,
  "log_color": false,
  "file_logging": false,
  "log_file_path": "./logs",
  "max_file_size_mb": 10,
  "clean_logs": false,
  "log_format": ["AppName", "Level", "Message", "Timestamp"],
  "log_time_format": "02-Jan-2006 15:04:05 MST",
  "log_timezone": "UTC"
}
```

All fields are optional. Any field you omit uses the default value shown above.

### Config Reference

| Field | Default | Description |
|-------|---------|-------------|
| `app_name` | `APP` | Name shown in every log line |
| `debug_level` | `false` | Enable debug logs |
| `log_color` | `false` | Colorize level output in terminal |
| `file_logging` | `false` | Write logs to file in addition to stdout |
| `log_file_path` | `./` | Directory to write log files |
| `max_file_size_mb` | `10` | Max size in MB before rotating to a new chunk |
| `clean_logs` | `false` | Delete old log files on startup |
| `log_format` | `["AppName","Level","Message","Timestamp"]` | Order of fields in each log line |
| `log_time_format` | `02-Jan-2006 15:04:05 MST` | Go time format string |
| `log_timezone` | `UTC` | Timezone for timestamps |

---

## File Logging

File logging is **off by default**. Stdout is always on.

To enable:

```json
{
  "file_logging": true,
  "log_file_path": "./logs"
}
```

Log files are created in chunks:

```
myservice_20260217_102412_chunk001.log
myservice_20260217_102412_chunk002.log
```

When a chunk reaches `max_file_size_mb`, quietlog closes it and opens the next one automatically. No logs are lost during rotation.

### Clean on Startup

```json
{
  "file_logging": true,
  "log_file_path": "./logs",
  "clean_logs": true
}
```

> **Warning:** `clean_logs: true` deletes all previous log files matching your app name on startup. Use with caution in production.

---

## Log Format

Reorder the fields in each log line:

```json
{
  "log_format": ["Timestamp", "Level", "AppName", "Message"]
}
```

**Supported fields:**

| Field | Example output |
|-------|----------------|
| `AppName` | `[myservice]` |
| `Level` | `[INFO]` |
| `Timestamp` | `[17-Feb-2026 10:24:12 UTC]` |
| `Message` | `: server started` |

---

## What quietlog Does NOT Do

quietlog intentionally does not implement:

- Structured / JSON logging
- Log rotation by time or date
- Log compression
- Context propagation
- Hook systems

These concerns are best handled outside the application.

---

## Compared To

| | quietlog | stdlib `log` | `slog` | `zap` / `zerolog` |
|--|---------|-------------|--------|-------------------|
| Human readable | ✅ | ✅ | ⚠️ | ❌ (JSON first) |
| Zero dependencies | ✅ | ✅ | ✅ | ❌ |
| Log levels | ✅ | ❌ | ✅ | ✅ |
| File logging | ✅ | ❌ | ❌ | ✅ |
| Auto-init | ✅ | ✅ | ❌ | ❌ |
| Config file | ✅ | ❌ | ❌ | ❌ |
| Structured output | ❌ | ❌ | ✅ | ✅ |
| High throughput | ⚠️ | ✅ | ✅ | ✅ |

quietlog is not trying to replace Zap, Zerolog, or slog. It exists for a different class of problems.

---

## Design Philosophy

- One job, done well
- Predictable behavior, no hidden magic
- Stdout first — files are opt-in
- Never crash on bad config — log it, fall back to defaults
- Concurrent-safe by default
- Easy to reason about

---

## Migration from v1.x

v2.0.0 is a breaking change.

| v1.x | v2.0.0 |
|------|--------|
| Config via environment variables | Config via `quietlog_config.json` |
| `LOG_LINES` line-based limiting | `max_file_size_mb` size-based chunk rotation |
| File logging always on | File logging opt-in via `file_logging: true` |
| `Asia/Kolkata` default timezone | `UTC` default timezone |
| No `Warn` level | `Warn` and `Warnf` added |

See [docs/migration.md](docs/migration.md) for the full migration guide.

---

## Documentation

- [Configuration Reference](docs/configuration.md)
- [File Logging & Rotation](docs/file-logging.md)
- [Log Format](docs/format.md)
- [Migration from v1.x](docs/migration.md)

---

## License

MIT
