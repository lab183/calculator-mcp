package main

import (
	"context"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func makeReq(args map[string]any) mcp.CallToolRequest {
	var r mcp.CallToolRequest
	r.Params.Name = "calculate"
	r.Params.Arguments = args
	return r
}

func TestCalculatorHandler(t *testing.T) {
	tests := []struct {
		name       string
		args       map[string]any
		wantResult string
		wantErr    string
	}{
		{name: "add", args: map[string]any{"operation": "add", "x": 3.0, "y": 4.0}, wantResult: "7.00"},
		{name: "subtract", args: map[string]any{"operation": "subtract", "x": 10.0, "y": 3.0}, wantResult: "7.00"},
		{name: "multiply", args: map[string]any{"operation": "multiply", "x": 6.0, "y": 7.0}, wantResult: "42.00"},
		{name: "divide", args: map[string]any{"operation": "divide", "x": 22.0, "y": 7.0}, wantResult: "3.14"},
		{name: "power", args: map[string]any{"operation": "power", "x": 2.0, "y": 8.0}, wantResult: "256.00"},
		{name: "sqrt", args: map[string]any{"operation": "sqrt", "x": 9.0}, wantResult: "3.00"},
		{name: "divide by zero", args: map[string]any{"operation": "divide", "x": 5.0, "y": 0.0}, wantErr: "cannot divide by zero"},
		{name: "sqrt negative", args: map[string]any{"operation": "sqrt", "x": -4.0}, wantErr: "cannot take square root"},
		{name: "missing y", args: map[string]any{"operation": "add", "x": 5.0}, wantErr: "y is required"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculatorHandler(context.Background(), makeReq(tc.args))

			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("expected error containing %q, got %q", tc.wantErr, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			text := result.Content[0].(mcp.TextContent).Text
			if text != tc.wantResult {
				t.Fatalf("got %q, want %q", text, tc.wantResult)
			}
		})
	}
}
