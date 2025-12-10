package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// AddInput defines the input parameters for the add tool
type AddInput struct {
	A float64 `json:"a" jsonschema:"first number"`
	B float64 `json:"b" jsonschema:"second number"`
}

// AddOutput defines the output for the add tool
type AddOutput struct {
	Result float64 `json:"result" jsonschema:"sum of a and b"`
}

// Add adds two numbers together
func Add(ctx context.Context, req *mcp.CallToolRequest, input AddInput) (*mcp.CallToolResult, AddOutput, error) {
	return nil, AddOutput{Result: input.A + input.B}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "test-mcp-server",
		Version: "v1.0.0",
	}, nil)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "add",
		Description: "adds two numbers together",
	}, Add)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
