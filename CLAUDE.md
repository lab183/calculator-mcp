# calculator-mcp

## Build & run

```bash
make build   # compile → ./calculator-mcp
make run     # build + run (stdio, default transport)
make         # list all available targets
```

Transport selection via flags:

```bash
./calculator-mcp                          # stdio (default)
./calculator-mcp --transport stdio        # stdio (explicit)
./calculator-mcp --transport http         # HTTP on :8080
./calculator-mcp --transport http --addr :9090  # HTTP on custom port
```

## Project structure

- `main.go` — entire server: tool definition + handler + `server.ServeStdio`
- `Makefile` — build, run, fmt, vet, tidy, help targets
- `go.mod` / `go.sum` — dependencies (primary: `github.com/mark3labs/mcp-go`)

## Key conventions

- All tool logic lives in the single handler closure in `main.go`.
- Use `request.RequireString` / `request.RequireFloat` / `request.GetFloat` to read arguments — do **not** index `request.Params.Arguments` directly (it is typed `any` in mcp-go v0.54+).
- `y` is intentionally optional to support `sqrt`; validate its presence inside the `default` branch for ops that need it.
