package google

import (
	"context"
	"os"
	"testing"

	"github.com/agi-cn/llmplugin/llm/openai"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGoogleWithoutSummary(t *testing.T) {

	_ = godo