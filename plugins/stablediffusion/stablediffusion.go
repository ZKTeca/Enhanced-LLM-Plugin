
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

func NewStableDiffusion(sdAddr string) *StableDiffusion {
	if len(sdAddr) == 0 {
		panic("stable diffusion address is empty")
	}

	return &StableDiffusion{
		sdAddr: sdAddr,
		client: &http.Client{},
	}
}

func (s *StableDiffusion) Do(ctx context.Context, query string) (answer string, err error) {
	req, err := s.newRequest(ctx, query)