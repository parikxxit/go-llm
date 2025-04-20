// Package reranker provides interfaces and types for reranking.
package reranker

import (
	"context"
)

// Document represents a document for reranking
type Document struct {
	ID   string
	Text string
}

// TokenUsage represents token usage information
type TokenUsage struct {
	PromptTokens int
	TotalTokens  int
}

// Result represents a reranking result
type Result struct {
	Document       Document
	Index          int
	RelevanceScore float64
}

// Request represents a reranking request
type Request struct {
	Model           string
	Query           string
	Documents       []Document
	TopN            int
	ReturnDocuments bool
	User            string
	ProviderParams  map[string]interface{}
}

// Response represents a reranking response
type Response struct {
	Object  string
	Model   string
	Results []Result
	Usage   TokenUsage
}

// Reranker defines the interface for reranking
type Reranker interface {
	// Rerank sends a reranking request
	Rerank(ctx context.Context, req *Request) (*Response, error)

	// GetName returns the name of the implementation
	GetRerankerName() string
}
