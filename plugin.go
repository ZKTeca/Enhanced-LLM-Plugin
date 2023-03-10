
package llmplugin

import "context"

type Plugin interface {
	Do(ctx context.Context, query string) (answer string, err error)

	GetName() string
	GetInputExample() string
	GetDesc() string
}

var _ Plugin = (*SimplePlugin)(nil)

type SimplePlugin struct {
	// Name of plugin.
	Name string

	// InputExample is the example of input.
	InputExample string

	// Desc is the description of plugin for LLM to understand what is the plugin and what for.
	Desc string