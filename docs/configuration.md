# Configuration Reference

quietlog works out of the box with zero configuration. This page covers all available options when you want to customize behavior.

---

## How Configuration Works

quietlog uses a simple priority order:

1. **`quietlog_config.json` at project root** — auto-discovered on startup
2. **Custom path via `quietlog.Config("path")`** — use when your config is elsewhere
3. **Built-in defaults** — used when no config file is found

The system never crashes on bad config. If the file is missing, unreadable, or contains invalid JSON, quietlog logs what went wrong to stdout and continues with defaults.

---

## Config File

Create a file named `quietlog_config.json` in your project root:

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

All fields are optional. Omitted fields use their default values.

---

## Field Reference

### `app_name`
**Type:** string
**Default:** `"APP"`

The application name shown in brackets on every log line.

```json
{ "app_name": "api-gateway" }
```
```
[api-gateway] [INFO] : server started [...]
```

---

### `debug_level`
**Type:** boolean
**Default:** `false`

Controls whether `Debug` and `Debugf` calls produce output. When `false`, all debug calls are silently ignored. Info, Warn, Error, and Fatal are always on.

```json
{ "debug_level": true }
```

---

### `log_color`
**Type:** boolean
**Default:** `false`

Colorizes the level label in stdout output. Useful for local development.

```json
{ "log_color": true }
```

Colors by level: INFO -> green, DEBUG -> cyan, WARN -> yellow, ERROR -> red, FATAL -> magenta.

> Disable this in production or when stdout is piped or redirected — ANSI codes will appear as raw characters in log aggregators.

---

### `file_logging`
**Type:** boolean
**Default:** `false`

Enables writing logs to files in addition to stdout. Stdout is always on regardless of this setting.

```json
{ "file_logging": true }
```

When enabled, also set `log_file_path` to control where files are written.

---

### `log_file_path`
**Type:** string
**Default:** `"./"`

The directory where log files are written. quietlog creates the directory if it does not exist.

```json
{ "log_file_path": "./logs" }
```

> This field is ignored when `file_logging` is `false`. If set without `file_logging: true`, quietlog will warn you.

---

### `max_file_size_mb`
**Type:** number
**Default:** `10`

Maximum size in megabytes before quietlog closes the current log file and opens a new chunk. Accepts decimal values.

```json
{ "max_file_size_mb": 50 }
```

> This field is ignored when `file_logging` is `false`.

---

### `clean_logs`
**Type:** boolean
**Default:** `false`

When `true`, quietlog deletes all existing log files matching the app name pattern on startup before creating the first chunk.

```json
{ "clean_logs": true }
```

> **Warning:** This is destructive. All previous log files for this app will be permanently deleted on every startup. Do not use in production unless you are certain this is what you want.

---

### `log_format`
**Type:** array of strings
**Default:** `["AppName", "Level", "Message", "Timestamp"]`

Controls the order of fields in each log line. Supported values:

| Value | Output |
|-------|--------|
| `"AppName"` | `[myservice]` |
| `"Level"` | `[INFO]` |
| `"Timestamp"` | `[17-Feb-2026 10:24:12 UTC]` |
| `"Message"` | `: your message here` |

```json
{ "log_format": ["Timestamp", "Level", "AppName", "Message"] }
```

Output:
```
[17-Feb-2026 10:24:12 UTC] [INFO] [myservice] : server started
```

Fields not included in the list are omitted entirely. Empty fields (e.g. an empty `app_name`) are also omitted — no blank brackets.

---

### `log_time_format`
**Type:** string
**Default:** `"02-Jan-2006 15:04:05 MST"`

The timestamp format using Go's reference time (`Mon Jan 2 15:04:05 MST 2006`).

```json
{ "log_time_format": "2006-01-02 15:04:05" }
```

See [Go time format documentation](https://pkg.go.dev/time#Layout) for all available tokens.

---

### `log_timezone`
**Type:** string
**Default:** `"UTC"`

The timezone for timestamps. Accepts any IANA timezone name.

```json
{ "log_timezone": "America/New_York" }
```

If an invalid timezone is provided, quietlog falls back to UTC and logs a warning.

Common values: `"UTC"`, `"Local"`, `"America/New_York"`, `"Europe/London"`, `"Asia/Kolkata"`, `"Asia/Tokyo"`.

---

## Custom Config Path

If your config file is not in the project root:

```go
func main() {
    quietlog.Config("./config/quietlog_config.json")
    quietlog.Info("server started")
}
```

`Config` must be called before the first log statement. If you log first, quietlog auto-initializes with defaults and `Config` will have no effect.

---

## What Happens When Config Goes Wrong

| Situation | Behavior |
|-----------|----------|
| File not found at root | Silent — uses defaults |
| File not found at custom path | Logs warning to stdout, uses defaults |
| Invalid JSON | Logs warning to stdout, uses defaults |
| Invalid timezone | Logs warning to stdout, falls back to UTC |
| Invalid file path for logs | Logs error to stderr, falls back to stdout only |
| `file_logging: false` but path provided | Logs warning, no file created |
