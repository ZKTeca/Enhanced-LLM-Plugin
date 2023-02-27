
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