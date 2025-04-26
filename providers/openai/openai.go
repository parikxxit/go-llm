package openai

import (
	"context"

	"github.com/parikxxit/go-llm/generator"
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	*openai.Client
	Model string
}

func NewOpenAI(cfg generator.Config) *OpenAI {
	return &OpenAI{
		Client: openai.NewClient(cfg.ApiKey),
		Model:  cfg.Model,
	}
}

func (o *OpenAI) Generate(ctx context.Context, req generator.Request) (*generator.Response, error) {
	//TODO: implement before merge

	return nil, nil
}

func (o *OpenAI) GenerateStream(ctx context.Context, req *generator.Request) (<-chan *generator.Response, error) {
	//TODO: implement before merge
	return nil, nil
}

func (o *OpenAI) GetName() (generator.Model, error) {
}
