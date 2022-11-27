package llm

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Role string

const (
	RoleUser      Role = openai.ChatMessageRoleUser
	RoleAssistant Role = openai.ChatMessageRoleAssistant
	RoleSystem    Role = openai.ChatMessageRoleSystem
)

func (r Role) String() string {
	return string(r)
}