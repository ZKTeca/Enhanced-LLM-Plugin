
// Package google is Google Search Plugin
// Get a Google Serach API key according to the Instruction.
// https://stackoverflow.com/questions/37083058/programmatically-searching-google-in-python-using-custom-search

package google

import (
	"context"
	"fmt"
	"strings"

	"github.com/agi-cn/llmplugin/llm"
	"github.com/pkg/errors"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

const (
	pluginName = "Google"

	pluginInputExample = "Who is Google boss?"

	pluginDesc = `Search something on the Internet by query input. You can search online to get the information you need, especially to get valid real-time information.`
)

type Google struct {
	customSearchID string
	apiToken       string

	summarizer llm.Summarizer
}

func NewGoogle(customSearchID, apiToken string, options ...Option) *Google {
	g := &Google{
		customSearchID: customSearchID,
		apiToken:       apiToken,
	}
