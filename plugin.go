
package llmplugin

import "context"

type Plugin interface {
	Do(ctx context.Context, query string) (answer string, err error)

	GetName() string