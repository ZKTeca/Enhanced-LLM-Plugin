package llm

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Role string

const (
	RoleUser      Role = openai.ChatMessageRole