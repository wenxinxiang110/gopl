package ch8

import (
	"fmt"
	"github.com/adonovan/gopl.io/ch5/links"
	"log"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func Crawl(works []string) {
	worlList := make(chan []string)

	go func() {
		worlList <- works
	}()

	seen := make(map[string]bool)
	for list := range worlList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worlList <- crawl(link)
				}(link)
			}
		}

	}
}
