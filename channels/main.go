package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //create a channel

	for _, link := range links {
		go checkStatus(link, c) //start go routine and give it a channel
		fmt.Println(<-c) //wait for response from channel
	}

}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link + ": Might be down" //response through the channel
		return
	}
	c <- link + ": UP" //response through the channel
}
