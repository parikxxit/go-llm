# go-llm

A minimalist Go library for interacting with LLM APIs. One-line integration with OpenAI, Anthropic, and other providers.

## Features

- 🚀 **Simple API** - Consistent interface across all providers
- 🔌 **Multiple Providers** - Support for OpenAI, Anthropic, and more
- 🧩 **Modular Design** - Easy to extend with new providers
- 🔄 **Streaming Support** - Efficient handling of streaming responses
- 🧮 **Embeddings** - Generate embeddings for semantic search
- 🔍 **Reranking** - Improve search result relevance (for supported providers)
- ⚙️ **Configurable** - Customize behavior with options
- 📊 **Observability** - Monitor usage and performance

## Installation

```bash
go get github.com/parikxxit/go-llm
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/parikxxit/go-llm"
	"github.com/parikxxit/go-llm/providers/openai"
)

func main() {
	// Create a provider with your API key
	provider := openai.New(os.Getenv("OPENAI_API_KEY"))

	// Create a client with the provider
	client := gollm.New(provider)

	// Create a request
	request := &gollm.GenerateRequest{
		Model: "gpt-3.5-turbo",
		Messages: []gollm.Message{
			{Role: "user", Content: "Hello, how are you?"},
		},
	}

	// Generate a response
	response, err := client.Generate(context.Background(), request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the response
	fmt.Println(response.Choices[0].Message.Content)
}
```

## Supported Providers

- [x] OpenAI
- [ ] Anthropic
- [ ] Google (Gemini)
- [ ] Azure OpenAI
- [ ] Cohere
- [ ] Mistral
- [ ] Ollama

## Usage

### Text Generation

```go
response, err := client.Generate(ctx, &gollm.GenerateRequest{
	Model: "gpt-3.5-turbo",
	Messages: []gollm.Message{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "What's the capital of France?"},
	},
	Temperature: 0.7,
	MaxTokens:   100,
})
```

### Streaming

```go
stream, err := client.GenerateStream(ctx, &gollm.GenerateRequest{
	Model: "gpt-3.5-turbo",
	Messages: []gollm.Message{
		{Role: "user", Content: "Write a poem about coding."},
	},
	Stream: true,
})

for response := range stream {
	fmt.Print(response.Choices[0].Message.Content)
}
```

### Embeddings

```go
embeddings, err := client.Embed(ctx, &gollm.EmbedRequest{
	Model: "text-embedding-3-small",
	Input: []string{"What is the meaning of life?"},
})
```

### Document Reranking

```go
results, err := client.Rerank(ctx, &gollm.RerankRequest{
	Model: "cohere-rerank-english-v2.0",
	Query: "quantum computing applications",
	Documents: []gollm.Document{
		{Text: "Quantum computing in cryptography and security"},
		{Text: "Classical computing architecture and design"},
		{Text: "Quantum algorithms for optimization problems"},
	},
	TopN: 2,
})
```

## Configuration Options

You can configure the client with additional options:

```go
client := gollm.New(provider).
	WithRetryCount(3).
	WithTimeout(30).
	WithDebug(true)
```

Or with fallback providers:

```go
primaryProvider := openai.New(openaiKey)
fallbackProvider := anthropic.New(anthropicKey)

client := gollm.New(primaryProvider).
	WithFallbackProviders(fallbackProvider)
```

## Roadmap

### Short-term (0-3 months)

- [ ] Complete OpenAI provider implementation
- [ ] Add Anthropic provider
- [ ] Add Google Gemini provider
- [ ] Implement proper streaming support
- [ ] Add retry logic with backoff
- [ ] Implement fallback provider logic
- [ ] Add basic observability metrics
- [ ] Add simple logging middleware
- [ ] Documentation improvements

### Medium-term (3-6 months)

- [ ] Add more providers (Mistral, Cohere, etc.)
- [ ] Add fine-tuning support
- [ ] Implement function calling capabilities
- [ ] Add rate limiting support
- [ ] Add provider selection middleware
- [ ] Improve error handling
- [ ] Add caching middleware
- [ ] Create integration test suite
- [ ] Add examples for types use cases

### Long-term (6+ months)

- [ ] Add intelligent routing between providers
- [ ] Implement cost optimization strategies
- [ ] Add provider performance analytics
- [ ] Implement advanced observability features
- [ ] Add support for custom models
- [ ] Create a CLI tool for interacting with models
- [ ] Add support for local models (Ollama, etc.)
- [ ] Implement context management for long conversations

## Detailed Feature Roadmap

### Core Features

- **Provider Interface**: types interface for all LLM providers
- **Text Generation**: Generate text completions from prompts
- **Streaming**: Support for streaming responses
- **Embeddings**: Generate vector embeddings for text
- **Reranking**: Rerank search results for better relevance
- **Function Calling**: Support for function/tool calling capabilities
- **Batching**: Process multiple requests in batches
- **Fine-tuning**: Customize models with your data

### Observability

- **Metrics**: Track request counts, latencies, token usage, etc.
- **Logging**: Structured logging of requests and responses
- **Tracing**: Distributed tracing support for complex workflows
- **Cost Tracking**: Monitor API costs across providers
- **Usage Analytics**: Visualize usage patterns and trends

### Reliability

- **Retries**: Automatic retries with configurable backoff
- **Fallbacks**: Graceful degradation with fallback providers
- **Rate Limiting**: Client-side rate limiting to prevent quota issues
- **Circuit Breaking**: Prevent cascading failures
- **Caching**: Cache responses for improved performance
- **Connection Pooling**: Efficient management of HTTP connections

### Extensibility

- **Middleware Support**: Add custom processing logic
- **Custom Providers**: Easy integration of new providers
- **Provider Selection**: Intelligent routing between providers
- **Plugins**: Extend functionality with plugins

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 