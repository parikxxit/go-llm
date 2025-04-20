// Package generator provides interfaces and types for text generation.
package generator

import (
	"context"
)

// Message represents a message in a conversation
type Message struct {
	Role    string
	Content string
}

// TokenUsage represents token usage information
type TokenUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Choice represents a choice in a generation response
type Choice struct {
	Index        int
	Message      Message
	FinishReason string
}

// Request represents a text generation request
type Request struct {
	Model          string
	Messages       []Message
	MaxTokens      int
	Temperature    float64
	TopP           float64
	Stream         bool
	Stop           []string
	User           string
	ProviderParams map[string]interface{}
}

// Response represents a text generation response
type Response struct {
	ID      string
	Object  string
	Created int64
	Model   string
	Choices []Choice
	Usage   TokenUsage
}

// Generator defines the interface for text generation
type Generator interface {
	// Generate sends a text generation request
	Generate(ctx context.Context, req *Request) (*Response, error)

	// GenerateStream sends a streaming text generation request
	GenerateStream(ctx context.Context, req *Request) (<-chan *Response, error)

	// GetName returns the name of the implementation
	GetName() string
}
