# calculator-mcp

## Build & run

```bash
go build ./...          # compile
go build -o calculator-mcp . && ./calculator-mcp  # run (stdio transport)
```

## Project structure

- `main.go` — entire server: tool definition + handler + `server.ServeStdio`
- `go.mod` / `go.sum` — dependencies (primary: `github.com/mark3labs/mcp-go`)

## Key conventions

- All tool logic lives in the single handler closure in `main.go`.
- Use `request.RequireString` / `request.RequireFloat` / `request.GetFloat` to read arguments — do **not** index `request.Params.Arguments` directly (it is typed `any` in mcp-go v0.54+).
- `y` is intentionally optional to support `sqrt`; validate its presence inside the `default` branch for ops that need it.
