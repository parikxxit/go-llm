package openai

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/parikxxit/go-llm/generator"
)

const (
	errNoModelResponse = "no response received from model"
	errOpenAIInternal  = "internal server error from OpenAI"
)

type OpenAI struct {
	Client openai.Client
	Model  string
}

func NewOpenAI(cfg generator.Config) *OpenAI {
	return &OpenAI{
		Client: openai.NewClient(
			option.WithAPIKey(cfg.ApiKey),
		),
		Model: cfg.Model,
	}
}

func (o *OpenAI) Generate(ctx context.Context, req *generator.Request) (*generator.Response, error) {
	messages := make([]openai.ChatCompletionMessageParamUnion, 0, len(req.Messages))
	for _, m := range req.Messages {
		switch m.Role {
		case generator.USER:
			messages = append(messages, openai.UserMessage(m.Content))
		case generator.ASSISTANT:
			messages = append(messages, openai.AssistantMessage(m.Content))
		}
	}

	chat, err := o.Client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    o.Model,
	})
	if err != nil {
		return nil, err
	}
	return getResponse(chat)
}

func (o *OpenAI) Chat(ctx context.Context, messages []generator.Message) (*generator.Response, error) {

	return nil, nil
}

func (o *OpenAI) GenerateStream(ctx context.Context, req *generator.Request) (<-chan *generator.Response, error) {
	//TODO: implement before merge
	return nil, nil
}

func (o *OpenAI) GetName() string {
	return o.Model
}

func getResponse(r *openai.ChatCompletion) (*generator.Response, error) {
	if len(r.Choices) == 0 {
		return nil, fmt.Errorf("%s: %s", errNoModelResponse, r.Model)
	}
	choice := r.Choices[0]
	return &generator.Response{
		ID:      uuid.New().String(),
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   r.Model,
		Content: choice.Message.Content,
		Usage: generator.TokenUsage{
			PromptTokens:     int(r.Usage.PromptTokens),
			CompletionTokens: int(r.Usage.CompletionTokens),
			TotalTokens:      int(r.Usage.TotalTokens),
		},
	}, nil
}
