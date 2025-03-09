package main

import (
	"fmt"
	"gofollower/pkg/linkfollower"
	"log"
	"net/url"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		log.Fatal("usage: follow <START_URL>")
	}
	startUrl, err := url.Parse(args[1])
	if err != nil {
		log.Fatalln(err)
	}
	responseHandler := func(visitedURL *url.URL, response linkfollower.VisitResponse) {
		fmt.Printf("%v -> %d %s\n", visitedURL, response.StatusCode, response.Additional)
	}
	err = linkfollower.Follow(startUrl, responseHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
