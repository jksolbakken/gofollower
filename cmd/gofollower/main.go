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
	responseHandler := func(response linkfollower.VisitResponse) {
		msg := fmt.Sprintf("%v -> %d", response.VisitedURL, response.StatusCode)
		if response.AdditionalInfo != "" {
			msg += fmt.Sprintf(" + %s", response.AdditionalInfo)
		}
		fmt.Printf("%s\n", msg)
	}
	err = linkfollower.Follow(startUrl, responseHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
