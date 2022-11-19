package llm

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Role string

const (
	RoleUser      Role = openai.ChatMessageRoleUser
	RoleAssistant Role = openai.ChatMessageRoleAssistant
	R