package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func main() {
	fmt.Println("worker pools in Go")

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)
	errorChan := make(chan error)

	for w := 1; w <= 5; w++ {
		go crawl(w, jobs, results, errorChan)
	}
	urls := []string{
		"http://google.com",
		"http://bbc.co.uk",
		"http://reddit.com",
		"http://youtube.com",
		"http://twitch.tv",
		"http://leetcod",
		"http://variety.com",
		"http://golang.org",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		jobs <- Site{URL: url} // send Site to jobs channel
		wg.Done()
	}
	wg.Wait()
	close(jobs)

	// loops forever
	for {
		select {
		case err := <-errorChan:
			log.Println(err)
		case result := <-results:
			log.Println(result)
		}
	}

	/*
		number_of_urls := len(urls)
		for a := 1; a <= number_of_urls-counter; a++ {
			result := <-results
			log.Println(result)
		}
	*/
}

// sending Sites to the jobs channel that we want to crawl from
// sending a Result to the results channel
func crawl(wId int, jobs <-chan Site, results chan<- Result, ec chan<- error) {

	// for every site that gets passed to this jobs channel
	for site := range jobs {
		log.Printf("worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
			ec <- errors.New("status code not 200")
			return
		}
		if resp.StatusCode != 200 {
			ec <- errors.New("status code not 200")
			return
		}
		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}
