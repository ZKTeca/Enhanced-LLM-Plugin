
package stablediffusion

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStableDiffusion(t *testing.T) {

	sd := NewStableDiffusion("127.0.0.1:19000")

	answer, err := sd.Do(context.Background(), "a girl play golf")
	require.NoError(t, err)

	assert.NotEmpty(t, answer)