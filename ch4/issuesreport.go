package ch4

import (
	"html/template"
	"io"
	"log"
	"time"
)

const (
	TemplRange = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var (
	report = template.Must(
		template.New("issuelist").
			Funcs(template.FuncMap{"daysAgo": daysAgo}).
			Parse(TemplRange))
)

func PrintIssues(is *IssuesSearchRes, w io.Writer) {
	if err := report.Execute(w, is); err != nil {
		log.Print(err)
	}
}
