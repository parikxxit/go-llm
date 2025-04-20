package openai

import (
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	*openai.Client
}

func NewOpenAI(apiKey string) *OpenAI {
	return &OpenAI{
		Client: openai.NewClient(apiKey),
	}
}
