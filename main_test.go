package main

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		input    AddInput
		expected float64
	}{
		{
			name:     "positive numbers",
			input:    AddInput{A: 2, B: 3},
			expected: 5,
		},
		{
			name:     "negative numbers",
			input:    AddInput{A: -5, B: -3},
			expected: -8,
		},
		{
			name:     "mixed signs",
			input:    AddInput{A: 10, B: -4},
			expected: 6,
		},
		{
			name:     "zero values",
			input:    AddInput{A: 0, B: 0},
			expected: 0,
		},
		{
			name:     "decimal numbers",
			input:    AddInput{A: 1.5, B: 2.5},
			expected: 4.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			req := &mcp.CallToolRequest{}

			_, output, err := Add(ctx, req, tt.input)
			if err != nil {
				t.Errorf("Add() error = %v", err)
				return
			}

			if output.Result != tt.expected {
				t.Errorf("Add() = %v, want %v", output.Result, tt.expected)
			}
		})
	}
}
