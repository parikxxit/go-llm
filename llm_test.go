package gollm

import (
	"context"
	"testing"
	"time"

	"github.com/parikxxit/go-llm/embedder"
	"github.com/parikxxit/go-llm/generator"
	"github.com/parikxxit/go-llm/providers/mock"
	"github.com/parikxxit/go-llm/reranker"
)

func TestClient_Generate(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM
	client := NewClient(mockLLM)

	// Check if the client has the generator capability
	if !client.HasGenerator() {
		t.Fatal("Client should have generator capability")
	}

	// Create a test request
	req := &generator.Request{
		Model: "test-model",
		Messages: []generator.Message{
			{
				Role:    "user",
				Content: "Hello, world!",
			},
		},
		MaxTokens:   100,
		Temperature: 0.7,
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Generate
	resp, err := client.Generate(ctx, req)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	// Verify the response
	if resp == nil {
		t.Fatal("Response is nil")
	}
	if len(resp.Choices) == 0 {
		t.Fatal("No choices in response")
	}
	if resp.Choices[0].Message.Content == "" {
		t.Fatal("Empty content in response")
	}
	if resp.Usage.TotalTokens == 0 {
		t.Fatal("No token usage in response")
	}
}

func TestClient_GenerateStream(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM
	client := NewClient(mockLLM)

	// Check if the client has the generator capability
	if !client.HasGenerator() {
		t.Fatal("Client should have generator capability")
	}

	// Create a test request
	req := &generator.Request{
		Model: "test-model",
		Messages: []generator.Message{
			{
				Role:    "user",
				Content: "Hello, world!",
			},
		},
		MaxTokens:   100,
		Temperature: 0.7,
		Stream:      true,
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call GenerateStream
	stream, err := client.GenerateStream(ctx, req)
	if err != nil {
		t.Fatalf("GenerateStream failed: %v", err)
	}

	// Read from the stream
	var responses []*generator.Response
	for resp := range stream {
		responses = append(responses, resp)
	}

	// Verify the responses
	if len(responses) == 0 {
		t.Fatal("No responses in stream")
	}
	for _, resp := range responses {
		if resp == nil {
			t.Fatal("Response is nil")
		}
		if len(resp.Choices) == 0 {
			t.Fatal("No choices in response")
		}
		if resp.Choices[0].Message.Content == "" {
			t.Fatal("Empty content in response")
		}
	}
}

func TestClient_Embed(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM
	client := NewClient(mockLLM)

	// Check if the client has the embedder capability
	if !client.HasEmbedder() {
		t.Fatal("Client should have embedder capability")
	}

	// Create a test request
	req := &embedder.Request{
		Model:      "test-model",
		Input:      []string{"Hello, world!"},
		Dimensions: 1536,
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Embed
	resp, err := client.Embed(ctx, req)
	if err != nil {
		t.Fatalf("Embed failed: %v", err)
	}

	// Verify the response
	if resp == nil {
		t.Fatal("Response is nil")
	}
	if len(resp.Data) == 0 {
		t.Fatal("No data in response")
	}
	if len(resp.Data[0].Embedding) == 0 {
		t.Fatal("Empty embedding in response")
	}
	if resp.Usage.TotalTokens == 0 {
		t.Fatal("No token usage in response")
	}
}

func TestClient_Rerank(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM
	client := NewClient(mockLLM)

	// Check if the client has the reranker capability
	if !client.HasReranker() {
		t.Fatal("Client should have reranker capability")
	}

	// Create a test request
	req := &reranker.Request{
		Model: "test-model",
		Query: "What is the capital of France?",
		Documents: []reranker.Document{
			{
				ID:   "1",
				Text: "Paris is the capital of France.",
			},
			{
				ID:   "2",
				Text: "London is the capital of England.",
			},
		},
		TopN:            2,
		ReturnDocuments: true,
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call Rerank
	resp, err := client.Rerank(ctx, req)
	if err != nil {
		t.Fatalf("Rerank failed: %v", err)
	}

	// Verify the response
	if resp == nil {
		t.Fatal("Response is nil")
	}
	if len(resp.Results) == 0 {
		t.Fatal("No results in response")
	}
	if resp.Results[0].RelevanceScore == 0 {
		t.Fatal("Zero relevance score in response")
	}
	if resp.Usage.TotalTokens == 0 {
		t.Fatal("No token usage in response")
	}
}

func TestClient_WithRetryCount(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM and retry count
	client := NewClient(mockLLM, WithRetryCount(3))

	// Verify the retry count
	if client.RetryCount() != 3 {
		t.Errorf("Expected retry count to be 3, got %d", client.RetryCount())
	}
}

func TestClient_WithFallbackGenerators(t *testing.T) {
	// Create mock LLMs
	mockLLM1 := mock.New()
	mockLLM2 := mock.New()

	// Create a client with the mock LLM and fallback generators
	client := NewClient(mockLLM1, WithFallbackGenerators([]generator.Generator{mockLLM2}))

	// Verify the fallback generators
	if len(client.FallbackGenerators()) != 1 {
		t.Errorf("Expected 1 fallback generator, got %d", len(client.FallbackGenerators()))
	}
	if client.FallbackGenerators()[0] != mockLLM2 {
		t.Error("Fallback generator does not match")
	}
}

func TestClient_WithTimeout(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM and timeout
	timeout := 10 * time.Second
	client := NewClient(mockLLM, WithTimeout(timeout))

	// Verify the timeout
	if client.Timeout() != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, client.Timeout())
	}
}

func TestClient_WithDebug(t *testing.T) {
	// Create a mock LLM
	mockLLM := mock.New()

	// Create a client with the mock LLM and debug mode
	client := NewClient(mockLLM, WithDebug(true))

	// Verify the debug mode
	if !client.Debug() {
		t.Error("Expected debug mode to be true")
	}
}
