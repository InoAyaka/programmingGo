package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {

	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}

	var result IssuesSearchResult

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil

}
