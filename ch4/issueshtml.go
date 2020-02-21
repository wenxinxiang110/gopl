package ch4

import (
	"html/template"
	"io"
	"log"
)

const (
	TemplList = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`
)

var issueList = template.Must(template.New("issuelist").Parse(TemplList))

func PrintIssuesHtml(is *IssuesSearchRes, w io.Writer) {
	if err := issueList.Execute(w, is); err != nil {
		log.Print(err)
	}
	log.Print("print issues into html success")
}
