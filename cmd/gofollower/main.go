package main

import (
	"fmt"
	"gofollower/pkg/linkfollower"
	"log"
	"net/url"
)

func main() {
	startUrl, err := url.Parse("https://tinyurl.com/m3q2xt")
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
