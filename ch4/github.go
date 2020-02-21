package ch4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

var IssuesParam = []string{
	"repo:golang/go",
	"is:open",
	"json",
	"decoder",
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssuesSearchRes struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchRes, error) {
	q := url.QueryEscape(strings.Join(terms, " "))

	resp, err := http.Get(IssuesURL + "?q=" + q)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var r IssuesSearchRes
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}
