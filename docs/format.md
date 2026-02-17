# Log Format

quietlog produces simple, line-based logs. Every line contains the same set of fields in a configurable order.

---

## Default Format

```
[APP] [INFO] : server started [17-Feb-2026 10:24:12 UTC]
```

Default field order: `AppName -> Level -> Message -> Timestamp`

---

## Changing Field Order

In `quietlog_config.json`:

```json
{
  "log_format": ["Timestamp", "Level", "AppName", "Message"]
}
```

Output:
```
[17-Feb-2026 10:24:12 UTC] [INFO] [myservice] : server started
```

---

## Available Fields

| Field name | Output example | Notes |
|------------|----------------|-------|
| `AppName` | `[myservice]` | Set via `app_name` config |
| `Level` | `[INFO]` | Always uppercase |
| `Timestamp` | `[17-Feb-2026 10:24:12 UTC]` | Format via `log_time_format` |
| `Message` | `: server started` | Always prefixed with `: ` |

Fields not listed in `log_format` are omitted entirely. Empty fields (e.g. empty `app_name`) are also omitted — no blank brackets ever appear.

---

## Timestamp Format

Uses Go's reference time. The default is:

```
02-Jan-2006 15:04:05 MST
```

Other examples:

```json
{ "log_time_format": "2006-01-02T15:04:05Z07:00" }
```
Output: `2026-02-17T10:24:12Z`

```json
{ "log_time_format": "15:04:05" }
```
Output: `10:24:12`

See [Go time format docs](https://pkg.go.dev/time#Layout) for all format tokens.

---

## Timezone

```json
{ "log_timezone": "America/New_York" }
```

Accepts any IANA timezone name. Defaults to `UTC`. Falls back to UTC with a warning if an invalid timezone is provided.

Common values: `"UTC"`, `"Local"`, `"America/New_York"`, `"Europe/London"`, `"Asia/Kolkata"`, `"Asia/Tokyo"`.

---

## Color

```json
{ "log_color": true }
```

When enabled, the `Level` field is colorized in stdout output:

| Level | Color |
|-------|-------|
| INFO | Green |
| DEBUG | Cyan |
| WARN | Yellow |
| ERROR | Red |
| FATAL | Magenta |

Color is applied to stdout only. File output is always plain text. Disable color when stdout is being piped or redirected to avoid raw ANSI codes in your output.

---

## Minimal Format Example

For the most compact output:

```json
{
  "log_format": ["Level", "Message"]
}
```

Output:
```
[INFO] : server started
[ERROR] : connection refused
```
