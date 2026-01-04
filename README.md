# quietlog

A tiny, opinionated, stdlib-first logging library for Go.

quietlog focuses on **clarity over cleverness**.
It avoids structured logging, reflection, and heavy abstractions — and instead provides readable, predictable logs suitable for both development and production environments.

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
- Built on Go’s `log` package
- Environment-driven configuration
- Safe to use from `init()` or `main()`
- Logs to stdout and/or file
- Simple line-based log limiting
- Customizable log format
- Predictable, production-safe defaults

---

## Installation

```bash
go get github.com/yourusername/quietlog
```

---

## Quick Start

```go
package main

import "github.com/yourusername/quietlog"

func main() {
	quietlog.Info("server started")
	quietlog.Debug("initializing cache")
	quietlog.Error("failed to connect to database")
}
```

You do **not** need to call `Init()` explicitly — quietlog initializes itself on first use.

---

## Log Output (Default)

```
[app] [INFO] [02-Jan-2026 10:24:12] : server started
[app] [DEBUG] [02-Jan-2026 10:24:13] : initializing cache
[app] [ERROR] [02-Jan-2026 10:24:14] : failed to connect to database
```

---

## Configuration (Environment Variables)

| Variable | Default | Description |
|--------|--------|------------|
| `APP_NAME` | `app` | Application name used in logs |
| `LOG_COLOR` | `true` | Enable or disable colored output |
| `LOG_FILE_PATH` | `./<app>_<timestamp>.log` | Log file path |
| `LOG_LINES` | `1000` | Max log lines before reset |
| `LOG_FORMAT` | `AppName,Level,Timestamp,Message` | Log format order |
| `LOG_TIME_FORMAT` | `02-Jan-2006 15:04:05` | Timestamp format |
| `LOG_TIMEZONE` | `Asia/Kolkata` | Timezone format |

> we have created and .env.template for you, you can simply copy and paste the items as per your choice.
---

## Custom Log Format

```bash
export LOG_FORMAT=Timestamp,Level,Message
```

Supported fields:
- AppName
- Level
- Timestamp
- Message

---

## File Logging

If LOG_FILE_PATH is not provided, quietlog automatically creates a file:

```bash
./<app_name>_<timestamp>.log
```

Logs are written to both stdout and file by default.

---

## Production Use

quietlog is safe for production use when:

- Logs are primarily consumed by humans
- Simplicity is preferred over features
- External log rotation tools are used (recommended)

quietlog intentionally does not implement:

- Log rotation
- Compression
- Structured logging
- Context propagation

These concerns are best handled outside the application.

---

## Design Philosophy

- One job, done well
- Predictable behavior
- No hidden magic
- Minimal global state
- Easy to reason about

quietlog is not trying to replace Zap, Zerolog, or Logrus.
It exists for a different class of problems.

---

## Final note (important)

This project aligns strongly with:
- Your engineering style
- Your preference for stable, long-term systems
- Go’s philosophy

## License

MIT
