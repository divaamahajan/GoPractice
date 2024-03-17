package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// List of websites to check
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
		"http://google.com",
	}
	// Create a channel to communicate between goroutines
	c := make(chan string)

	// Iterate over each link
	for _, link := range links {
		// Spawn a goroutine to check the link concurrently
		go checkLink(c, link)
	}

	// Alternative 1: Loop over the channel and print results
	// for i:=0; i<len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Alternative 2: Continuously check links in an infinite loop
	for l := range c {
		// Wait for the channel to return a value and assign it to l
		// Spawn a goroutine with an anonymous function to check the link after a delay
		go func(link string) {
			// Delay for 5 seconds before checking the link again
			time.Sleep(5 * time.Second)
			checkLink(c, link)
		}(l)
	}
}

// Function to check if a link is up or down
func checkLink(c chan string, link string) {
	// Perform an HTTP GET request to the link
	_, err := http.Get(link)

	// Check if there was an error during the GET request
	if err != nil {
		// Print the link and the error message indicating it might be down
		fmt.Println(link, err, "might be down")
		// Send the link through the channel to indicate it might be down
		c <- link
		// Alternative 1: Send a formatted string through the channel
		// c <- fmt.Sprintf("%s %v might be down", link, err)
		return
	}
	// If no error occurred, print that the link is up
	fmt.Println(link, "is up")
	// Send the link through the channel to indicate it is up
	c <- link
	// Alternative 1: Send a formatted string through the channel
	// c <- fmt.Sprintf("%s is up", link)
}
