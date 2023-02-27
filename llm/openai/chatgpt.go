
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