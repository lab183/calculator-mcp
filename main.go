package main

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer(
		"Calculator Demo",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)
	// Add a calculator tool
	calculatorTool := mcp.NewTool("calculate",
		mcp.WithDescription("Perform basic arithmetic operations"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
			mcp.Enum("add", "subtract", "multiply", "divide", "power", "sqrt"),
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("First number"),
		),
		mcp.WithNumber("y",
			mcp.Description("Second number (not required for sqrt)"),
		),
	)
	// Add the calculator handler
	s.AddTool(calculatorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		op, err := request.RequireString("operation")
		if err != nil {
			return nil, err
		}
		x, err := request.RequireFloat("x")
		if err != nil {
			return nil, err
		}
		var result float64
		switch op {
		case "sqrt":
			if x < 0 {
				return nil, errors.New("cannot take square root of a negative number")
			}
			result = math.Sqrt(x)
		default:
			y, yErr := request.RequireFloat("y")
			if yErr != nil {
				return nil, errors.New("y is required for this operation")
			}
			switch op {
			case "add":
				result = x + y
			case "subtract":
				result = x - y
			case "multiply":
				result = x * y
			case "divide":
				if y == 0 {
					return nil, errors.New("cannot divide by zero")
				}
				result = x / y
			case "power":
				result = math.Pow(x, y)
			}
		}
		return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
	})
	// Start the server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
