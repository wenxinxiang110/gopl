package main

import (
	"gopl/chapter1"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/gif", chapter1.GifHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
