package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://reddit.com",
	}

	c := make(chan string) // create a channel

	for _, link := range links {
		go checkLink(link, c) // go keyword creates a new go routine
	}
	for l := range c { // infinite loop, wait for a value to be sent into the channel
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
	// for {
	// 	go checkLink(<-c, c)  // wait for the channel to return a value and then call checkLink
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, "is up!")
	}
	c <- link
}
