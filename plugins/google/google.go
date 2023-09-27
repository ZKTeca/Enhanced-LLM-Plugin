
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

	for _, o := range options {
		o(g)
	}

	return g
}

func (g Google) Do(ctx context.Context, query string) (answer string, err error) {

	results, err := g.doSearch(ctx, query)
	if err != nil {
		return "", err
	}

	return g.makeResult(ctx, query, results)
}

func (g Google) doSearch(ctx context.Context, query string) (*customsearch.Search, error) {
	client, err := customsearch.NewService(ctx, option.WithAPIKey(g.apiToken))
	if err != nil {
		return nil, errors.Wrap(err, "new google service failed")
	}

	results, err := client.Cse.List().Q(query).Cx(g.customSearchID).Do()
	if err != nil {
		return nil, errors.Wrap(err, "google search failed")
	}

	return results, nil
}

func (g Google) makeResult(ctx context.Context, query string, results *customsearch.Search) (string, error) {

	items := results.Items
	if len(items) == 0 {
		return "Google don't known", nil
	}

	if g.summarizer == nil {
		return g.makeRawResult(ctx, items)
	} else {
		return g.makeResultBySummary(ctx, query, items)
	}
}

func (g Google) makeRawResult(ctx context.Context, items []*customsearch.Result) (string, error) {