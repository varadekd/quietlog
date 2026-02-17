# Migration Guide — v1.x to v2.0.0

v2.0.0 is a breaking change. The configuration system has been completely replaced and several behaviors have changed. This guide covers everything you need to update.

---

## What Changed

### 1. Configuration: Environment Variables -> JSON File

**v1.x** used environment variables:

```bash
APP_NAME=myservice
LOG_COLOR=true
LOG_FILE_PATH=./logs
LOG_LINES=1000
LOG_FORMAT=AppName,Level,Timestamp,Message
LOG_TIME_FORMAT=02-Jan-2006 15:04:05
LOG_TIMEZONE=Asia/Kolkata
```

**v2.0.0** uses `quietlog_config.json`:

```json
{
  "app_name": "myservice",
  "log_color": true,
  "file_logging": true,
  "log_file_path": "./logs",
  "max_file_size_mb": 10,
  "log_format": ["AppName", "Level", "Timestamp", "Message"],
  "log_time_format": "02-Jan-2006 15:04:05 MST",
  "log_timezone": "UTC"
}
```

Drop this file in your project root. No other changes needed for basic configuration.

---

### 2. File Logging Is Now Opt-In

**v1.x** always created a log file automatically.

**v2.0.0** logs to stdout only by default. To enable file logging:

```json
{
  "file_logging": true,
  "log_file_path": "./logs"
}
```

If you relied on log files being created automatically, add these two fields.

---

### 3. Line-Based Limiting -> Size-Based Chunk Rotation

**v1.x** used `LOG_LINES=1000` to limit log file size by line count. Behavior on reaching the limit was unclear.

**v2.0.0** uses `max_file_size_mb` with clear, predictable behavior:

- When the current chunk reaches the size limit, it is closed
- A new chunk file is opened automatically
- No logs are lost during rotation
- Files are named `<app_name>_<timestamp>_chunk001.log`, `chunk002.log`, etc.

```json
{
  "max_file_size_mb": 50
}
```

---

### 4. Default Timezone Changed

**v1.x** defaulted to `Asia/Kolkata`.

**v2.0.0** defaults to `UTC`.

If you relied on the IST default, set it explicitly:

```json
{
  "log_timezone": "Asia/Kolkata"
}
```

---

### 5. `Warn` Level Added

**v1.x** had no `Warn` level.

**v2.0.0** adds `Warn` and `Warnf`:

```go
quietlog.Warn("disk usage above 80%")
quietlog.Warnf("memory at %d%%", usage)
```

No action required — just start using it where appropriate.

---

### 6. `Fatal` Panics Instead of Exiting

**v2.0.0** `Fatal` and `Fatalf` call `panic` after logging, not `os.Exit(1)`. This gives you a full stack trace, making it easier to diagnose what went wrong.

If you need guaranteed exit behavior without recovery:

```go
defer func() {
    if r := recover(); r != nil {
        os.Exit(1)
    }
}()
```

---

### 7. Custom Config Path via `quietlog.Config()`

**v1.x** had no way to specify a custom config path.

**v2.0.0** adds `quietlog.Config("path")`:

```go
func main() {
    quietlog.Config("./configs/quietlog_config.json")
    quietlog.Info("server started")
}
```

---

## Migration Checklist

```
[ ] Delete .env or environment variable config for quietlog
[ ] Create quietlog_config.json in project root
[ ] Add "file_logging": true if you need log files
[ ] Add "log_timezone": "Asia/Kolkata" if you relied on IST default
[ ] Replace LOG_LINES with max_file_size_mb (value is MB, not lines)
[ ] Run: go get github.com/varadekd/quietlog@v2.0.0
[ ] Run your tests — the public API is unchanged
```

---

## API Compatibility

The public API functions are unchanged between v1.x and v2.0.0. Your logging call sites require no changes.

```go
quietlog.Info(msg)
quietlog.Warn(msg)      // new in v2.0.0
quietlog.Error(msg)
quietlog.Fatal(msg)
quietlog.Debug(msg)
quietlog.Infof(format, args...)
quietlog.Warnf(format, args...)   // new in v2.0.0
quietlog.Errorf(format, args...)
quietlog.Fatalf(format, args...)
quietlog.Debugf(format, args...)
quietlog.Init()
quietlog.Config(path)   // new in v2.0.0
```
