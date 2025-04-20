// Package types provides shared types used across the go-llm packages.
package types

// TokenUsage represents token usage information
type TokenUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}
