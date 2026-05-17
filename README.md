# calculator-mcp

A simple [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server written in Go that exposes a calculator tool. Use it as a reference for building MCP servers with [mcp-go](https://github.com/mark3labs/mcp-go).

## Tools

### `calculate`

Performs basic arithmetic operations.

| Parameter   | Type   | Required | Description                                      |
|-------------|--------|----------|--------------------------------------------------|
| `operation` | string | yes      | One of: `add`, `subtract`, `multiply`, `divide`, `power`, `sqrt` |
| `x`         | number | yes      | First operand                                    |
| `y`         | number | no*      | Second operand (*required for all ops except `sqrt`) |

**Examples**

| operation  | x  | y  | result |
|------------|----|----|--------|
| add        | 3  | 4  | 7.00   |
| subtract   | 10 | 3  | 7.00   |
| multiply   | 6  | 7  | 42.00  |
| divide     | 22 | 7  | 3.14   |
| power      | 2  | 8  | 256.00 |
| sqrt       | 9  |    | 3.00   |

## Requirements

- Go 1.25+
- Docker (optional, for container builds)

## Usage

### Stdio (default)

Used by Claude Desktop and other MCP CLI hosts:

```bash
make run
# or
./calculator-mcp --transport stdio
```

### HTTP (Streamable HTTP transport)

Used by browser-based or networked MCP clients:

```bash
./calculator-mcp --transport http --addr :8080
```

Verify with curl:

```bash
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list","params":{}}'
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--transport` | `stdio` | `stdio` or `http` |
| `--addr` | `:8080` | Listen address (HTTP transport only) |

### Docker

Build and run as a container (serves HTTP on port 8080):

```bash
make docker-build
make docker-run
# or manually:
docker build -t calculator-mcp .
docker run --rm -p 8080:8080 calculator-mcp
```

### Common Make targets

```
make build         # compile the binary
make run           # build and run (stdio)
make docker-build  # build the Docker image
make docker-run    # run the server in Docker on port 8080
make fmt           # format source files
make vet           # run go vet
make tidy          # tidy go.mod / go.sum
make               # show all targets
```

### Connect via Claude Desktop

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "calculator": {
      "command": "/path/to/calculator-mcp"
    }
  }
}
```

## Dependencies

- [mcp-go v0.54](https://github.com/mark3labs/mcp-go) — MCP server SDK for Go
