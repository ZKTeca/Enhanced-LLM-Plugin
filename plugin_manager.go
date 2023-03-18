
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