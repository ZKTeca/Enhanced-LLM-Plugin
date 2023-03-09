
package openai

import (
	"context"

	"github.com/agi-cn/llmplugin/llm"

	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	model  string
	client *openai.Client
}

type Option func(c *ChatGPT)

func WithModel(model string) Option {
	return func(c *ChatGPT) {
		c.model = model
	}
}

func NewChatGPT(token string, opts ...Option) *ChatGPT {

	client := openai.NewClient(token)

	chatgpt := &ChatGPT{
		model:  openai.GPT3Dot5Turbo,
		client: client,
	}

	for _, opt := range opts {
		opt(chatgpt)
	}

	return chatgpt
}

func (c ChatGPT) Summary(ctx context.Context, content string) (string, error) {

	messages := []llm.LlmMessage{
		{
			Role:    llm.RoleUser,
			Content: content,