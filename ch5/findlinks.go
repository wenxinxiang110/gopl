package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

func FindLinks(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		return
	}

	//for _, link := range Visit(nil, doc) {
	//	fmt.Println(link)
	//}

	Outline(nil, doc)
}

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}

	return links

}

func Outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Outline(stack, c)
	}

}
