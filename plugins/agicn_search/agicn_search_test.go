
package agicn_search

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgicnSearch(t *testing.T) {

	// TODO(zy): fix agi.cn search
	t.Skip("agi.cn search not valid NOW")

	ts := []struct {
		testname string
		query    string
	}{
		{
			"Search in english",
			"NBA schedule today",
		},
		{
			"Search in chinese",
			"今天nba有哪些比赛",
		},
	}

	s := NewAgicnSearch()