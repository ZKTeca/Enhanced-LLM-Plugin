
package stablediffusion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	pluginName         = "StableDiffusion"
	pluginInputExample = "A beautiful girl"
	pluginDesc         = `Stable diffusion is text-to-image model capable of generating images given any text input`
)

type StableDiffusion struct {
	sdAddr string

	client *http.Client
}
