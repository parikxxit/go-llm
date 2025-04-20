// Package embedder provides interfaces and types for embeddings.
package embedder

import (
	"context"
)

// TokenUsage represents token usage information
type TokenUsage struct {
	PromptTokens int
	TotalTokens  int
}

// EmbedData represents embedding data
type EmbedData struct {
	Object    string
	Embedding []float64
	Index     int
}

// Request represents an embedding request
type Request struct {
	Model          string
	Input          []string
	Dimensions     int
	User           string
	ProviderParams map[string]interface{}
}

// Response represents an embedding response
type Response struct {
	Object string
	Model  string
	Data   []EmbedData
	Usage  TokenUsage
}

// Embedder defines the interface for embeddings
type Embedder interface {
	// Embed sends an embedding request
	Embed(ctx context.Context, req *Request) (*Response, error)

	// GetName returns the name of the implementation
	GetEmbedderName() string
}
