
package llmplugin

import (
	"context"
	"os"
	"testing"

	"github.com/agi-cn/llmplugin/llm"
	"github.com/agi-cn/llmplugin/llm/openai"
	"github.com/agi-cn/llmplugin/plugins/calculator"
	"github.com/agi-cn/llmplugin/plugins/google"
	"github.com/agi-cn/llmplugin/plugins/stablediffusion"
	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManagerSelectPlugin(t *testing.T) {
	manager := newChatGPTManager()

	t.Run("Digital Computing", func(t *testing.T) {
		pluginCtxs, err := manager.Select(context.Background(), "10 add 20 equals ?")
		require.NoError(t, err)

		require.Equal(t, 1, len(pluginCtxs))
		require.True(t, includePlugin(pluginCtxs, "Calculator"))

		choices := pluginCtxs[0]

		answer, err := choices.Plugin.Do(context.Background(), choices.Input)
		require.NoError(t, err)

		assert.Equal(t, "30", answer)
	})

	t.Run("Query Weather", func(t *testing.T) {
		choices, err := manager.Select(context.Background(), "How is the weather today?")
		assert.NoError(t, err)

		assert.NotEmpty(t, choices)
		assert.True(t, includePlugin(choices, "Weather"))
	})

	t.Run("Stable Diffusion", func(t *testing.T) {
		choices, err := manager.Select(context.Background(), "Draw a girl image")
		assert.NoError(t, err)

		assert.NotEmpty(t, choices)
		assert.True(t, includePlugin(choices, "StableDiffusion"))
	})

	t.Run("Google", func(t *testing.T) {

		choices, err := manager.Select(context.Background(), "NBA 总决赛现在如何？")
		assert.NoError(t, err)

		assert.NotEmpty(t, choices)
		assert.True(t, includePlugin(choices, "Google"))
	})
}

func TestManagerSelectPlugin_WithoutChoice(t *testing.T) {
	manager := newChatGPTManager()

	t.Run("Quick Sort Source Code", func(t *testing.T) {
		choices, err := manager.Select(context.Background(), "quick sort source code in python")
		assert.NoError(t, err)

		assert.Empty(t, choices)
	})

}

func includePlugin(pluginCtxs []PluginContext, target string) bool {