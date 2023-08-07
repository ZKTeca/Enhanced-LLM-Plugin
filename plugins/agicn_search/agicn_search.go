package agicn_search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	pluginName         = "AgicnSearch"
	pluginInputExample = "Who is Google boss?"
	pluginDesc         = `Search something by query input.`

	baseURL = "https://agicn-ducksearch.vercel.app/search"
)

type searchResponse struct {
	Title string `json:"title"`
	Href  string `json:"href"`
	Body  string `json:"body"`
}

type AgicnSearch struct {
	client *http.Client
}

func NewAgicnSearch() *AgicnSearch {
	c := &http.Client{}
	return &AgicnSearch{c}
}

func (s AgicnSearch) Do(ctx context.Context, query string) (answer string, err error) {
	searchResults, err := s.doHTTPRequest(ctx, query)
	if err != nil {
		return "", err
	}

	answer = s.makeAnswer(searchResults)
	return answer, nil
}

func (s AgicnSearch) doHTTPRequest(ctx context.Context, query string) ([]searchResponse, error) {
	params := url.Values{}
	params.Add("q", query)

	url := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp