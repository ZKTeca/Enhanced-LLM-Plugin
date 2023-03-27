
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