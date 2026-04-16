# sunnyside

TUI-based recipe repository & meal planner (Go).

## Project layout

- `main.go` → `cmd.Execute()` — CLI entry
- `cmd/root.go` — cobra root command (currently a no-op)
- `cmd/setup.go` — `sunnyside setup` / `sunnyside s` — runs TUI init prompt
- `internal/cfg/config.go` — TOML config read/write
- `internal/tui/model.go` — bubbletea TUI model for setup prompt
- No tests exist yet

## Known issues

- `internal/cfg/config.go` — `ToFile` (line 28) is incomplete; `toml.Marshaler` call is wrong (use `toml.Marshal`), and the function has a missing return and unused imports (`errors`, `path`).
- `cmd/setup.go:38` — `cfgPath` is undefined; the setup command references this variable but it's never declared in scope.
- `internal/tui/model.go:82` — `strings.Builder` type conversion error (LSP).
