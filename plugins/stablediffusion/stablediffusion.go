
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