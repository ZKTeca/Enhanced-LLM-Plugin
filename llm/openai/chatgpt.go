
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
		},
	}

	answer, err := c.Chat(ctx, messages)
	if err != nil {
		return "", err
	}

	return answer.Content, nil
}

func (c ChatGPT) Chat(ctx context.Context, messages []llm.LlmMessage) (*llm.LlmAnswer, error) {

	chatGPTMessages := c.makeChatGPTMessage(messages)

	return c.send(ctx, chatGPTMessages)

}

func (c ChatGPT) makeChatGPTMessage(messages []llm.LlmMessage) []openai.ChatCompletionMessage {

	chatGPTMessages := make([]openai.ChatCompletionMessage, 0, len(messages))
	for _, m := range messages {
		chatGPTMessages = append(chatGPTMessages, openai.ChatCompletionMessage{
			Role:    m.Role.String(),
			Content: m.Content,
		})