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

- Go 1.21+

## Usage

Build and run the server (communicates over stdio):

```bash
make run
```

Or without Make:

```bash
go build -o calculator-mcp .
./calculator-mcp
```

Common Make targets:

```
make build   # compile the binary
make run     # build and run
make fmt     # format source files
make vet     # run go vet
make tidy    # tidy go.mod / go.sum
make         # show all targets
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
