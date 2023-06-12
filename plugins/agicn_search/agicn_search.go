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
	c := &http.Cli