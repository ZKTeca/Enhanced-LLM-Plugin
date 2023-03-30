
package llmplugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/agi-cn/llmplugin/llm"

	"github.com/sirupsen/logrus"
)

type PluginContext struct {
	Plugin

	// Input for handle function of plugin.
	Input string
}

type PluginManager struct {
	llmer llm.LLMer

	// plugins <key:name, value:Plugin>
	plugins map[string]Plugin
}

type PluginManagerOpt func(manager *PluginManager)

// WithPlugin enable one plugin.
func WithPlugin(p Plugin) PluginManagerOpt {

	return func(manager *PluginManager) {
		name := strings.ToLower(p.GetName())
		if _, ok := manager.plugins[name]; !ok {
			manager.plugins[name] = p
		}
	}
}

// WithPlugins enable multiple plugins.
func WithPlugins(plugins []Plugin) PluginManagerOpt {

	return func(manager *PluginManager) {

		for _, p := range plugins {
			opt := WithPlugin(p)
			opt(manager)
		}
	}
}

// NewPluginManager create plugin manager.
func NewPluginManager(llmer llm.LLMer, opts ...PluginManagerOpt) *PluginManager {

	manager := &PluginManager{
		llmer:   llmer,
		plugins: make(map[string]Plugin, 4),
	}

	for _, opt := range opts {
		opt(manager)
	}

	return manager
}

// Select to choice some plugin to finish the task.
func (m *PluginManager) Select(ctx context.Context, query string) ([]PluginContext, error) {

	answer, err := m.chatWithLlm(ctx, query)
	if err != nil {
		logrus.Errorf("chat with llm error: %v", err)
		return nil, err
	}

	pluginCtxs := m.choicePlugins(answer)

	// for debug
	for _, c := range pluginCtxs {
		logrus.Debugf("query: %s choice plugins: %s input: %s", query, c.GetName(), c.Input)
	}

	return pluginCtxs, nil
}

func (m *PluginManager) makePrompt(query string) string {

	tools := m.makeTaskList()

	prompt := fmt.Sprintf(`You will performs one task based on the following object:
	%s

	You can call one or multiple of the following functions in triple backticks:
	'''
	%s
	'''

	In each response, you must start with a function call like Tool name and args, split by ':',like:
	Google: query
	Weather:

	Don't explain why you use a tool. If you cannot figure out the answer, you say 'I don’t know'.

	Select only the corresponding tool and do not return any results.`,

		query,
		tools,
	)

	return prompt
}

func (m *PluginManager) makeTaskList() string {
