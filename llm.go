// Package gollm provides a simple, unified interface for working with various LLMs.
package gollm

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/parikxxit/go-llm/embedder"
	"github.com/parikxxit/go-llm/generator"
	"github.com/parikxxit/go-llm/reranker"
	"github.com/rs/zerolog"
)

// Client represents a gollm client for interacting with LLMs
type Client struct {
	llm               generator.Generator
	embedder          embedder.Embedder
	reranker          reranker.Reranker
	retryCount        int
	fallbackGenerator []generator.Generator
	fallbackEmbedder  []embedder.Embedder
	fallbackReranker  []reranker.Reranker
	timeout           time.Duration
	debug             bool
	logger            zerolog.Logger
}

// NewClient creates a new gollm client with the specified LLM implementation
func NewClient(llm generator.Generator, opts ...Option) *Client {
	if llm == nil {
		panic("llm cannot be nil")
	}

	client := &Client{
		llm:        llm,
		retryCount: 3,
		timeout:    30 * time.Second,
		debug:      false,
	}

	// Check if the LLM implements additional capabilities
	if e, ok := llm.(embedder.Embedder); ok {
		client.embedder = e
	}

	if r, ok := llm.(reranker.Reranker); ok {
		client.reranker = r
	}

	for _, opt := range opts {
		opt(client)
	}
	// Initialize logger with default settings and generator name
	client.logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Str("generator", client.llm.GetName()).
		Logger()

	return client
}

// WithEmbedder creates a new client with an additional embedder
func WithEmbedder(emb embedder.Embedder) Option {
	return func(c *Client) {
		c.embedder = emb
	}
}

// WithReranker creates a new client with an additional reranker
func WithReranker(rer reranker.Reranker) Option {
	return func(c *Client) {
		c.reranker = rer
	}
}

// HasGenerator returns true if the client has a generator
func (c *Client) HasGenerator() bool {
	return c.llm != nil
}

// HasEmbedder returns true if the client has an embedder
func (c *Client) HasEmbedder() bool {
	return c.embedder != nil
}

// HasReranker returns true if the client has a reranker
func (c *Client) HasReranker() bool {
	return c.reranker != nil
}

// Generate sends a text generation request to the LLM
func (c *Client) Generate(ctx context.Context, request *generator.Request) (*generator.Response, error) {
	if c.llm == nil {
		return nil, fmt.Errorf("generator capability not available")
	}

	if c.debug {
		c.logger.Info().Msgf("Generating Response for req:%s", request.Messages[0].Content)
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.llm.Generate(ctx, request)
	if err != nil {
		// TODO: Add retry logic with fallback generators
		return nil, err
	}

	return resp, nil
}

// GenerateStream sends a streaming text generation request to the LLM
func (c *Client) GenerateStream(ctx context.Context, request *generator.Request) (<-chan *generator.Response, error) {
	if c.llm == nil {
		return nil, fmt.Errorf("generator capability not available")
	}

	if c.debug {
		c.logger.Info().Msgf("started streaming req with msg:%s", request.Messages[0].Content)
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	stream, err := c.llm.GenerateStream(ctx, request)
	if err != nil {
		// TODO: Add retry logic with fallback generators
		return nil, err
	}

	return stream, nil
}

// Embed sends an embedding request to the LLM
func (c *Client) Embed(ctx context.Context, request *embedder.Request) (*embedder.Response, error) {
	if c.embedder == nil {
		return nil, fmt.Errorf("embedder capability not available")
	}

	if c.debug {
		c.logger.Info().Msgf("embedding: %s with embedder: %s", request.Model, request.Input[0])
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.embedder.Embed(ctx, request)
	if err != nil {
		// TODO: Add retry logic with fallback embedders
		return nil, err
	}

	return resp, nil
}

// Rerank sends a reranking request to the LLM
func (c *Client) Rerank(ctx context.Context, request *reranker.Request) (*reranker.Response, error) {
	if c.reranker == nil {
		return nil, fmt.Errorf("reranker capability not available")
	}

	if c.debug {
		c.logger.Info().Msgf("reranking matches")
	}

	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	resp, err := c.reranker.Rerank(ctx, request)
	if err != nil {
		// TODO: Add retry logic with fallback rerankers
		return nil, err
	}

	return resp, nil
}

// RetryCount returns the number of retries configured for the client
func (c *Client) RetryCount() int {
	return c.retryCount
}

// FallbackGenerators returns the list of fallback generators configured for the client
func (c *Client) FallbackGenerators() []generator.Generator {
	return c.fallbackGenerator
}

// FallbackEmbedders returns the list of fallback embedders configured for the client
func (c *Client) FallbackEmbedders() []embedder.Embedder {
	return c.fallbackEmbedder
}

// FallbackRerankers returns the list of fallback rerankers configured for the client
func (c *Client) FallbackRerankers() []reranker.Reranker {
	return c.fallbackReranker
}

// Timeout returns the timeout configured for the client
func (c *Client) Timeout() time.Duration {
	return c.timeout
}

// Debug returns whether debug mode is enabled for the client
func (c *Client) Debug() bool {
	return c.debug
}

// Option is a function that configures a Client
type Option func(*Client)

// WithRetryCount sets the number of retries for the client
func WithRetryCount(count int) Option {
	return func(c *Client) {
		c.retryCount = count
	}
}

// WithFallbackGenerators sets the fallback generators for the client
func WithFallbackGenerators(generators []generator.Generator) Option {
	return func(c *Client) {
		c.fallbackGenerator = generators
	}
}

// WithFallbackEmbedders sets the fallback embedders for the client
func WithFallbackEmbedders(embedders []embedder.Embedder) Option {
	return func(c *Client) {
		c.fallbackEmbedder = embedders
	}
}

// WithFallbackRerankers sets the fallback rerankers for the client
func WithFallbackRerankers(rerankers []reranker.Reranker) Option {
	return func(c *Client) {
		c.fallbackReranker = rerankers
	}
}

// WithTimeout sets the timeout for the client
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithDebug enables debug mode for the client
func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}
