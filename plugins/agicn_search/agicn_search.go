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

	baseURL = "https://agi