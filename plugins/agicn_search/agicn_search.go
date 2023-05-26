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
	pluginInputExample = "W