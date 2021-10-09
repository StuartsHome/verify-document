package http_client

import (
	"net/http"
	"time"
)

type Client interface {
	Do(request *http.Request) (*http.Response, error)
	Get(url string) (*http.Response, error)
}

func makeHttpClient(idleConnectionsPerHost int) *http.Client {
	transport := &http.Transport{
		IdleConnTimeout:     10 * time.Second,
		MaxIdleConnsPerHost: idleConnectionsPerHost,
	}
	return &http.Client{
		Transport: transport,
	}
}

type input struct {
	request *http.Request
	url     string
	method  int
	out     chan output
}

type output struct {
	response *http.Response
	error    error
}

type workerClient struct {
	client   Client
	requests chan input
}

func (wc *workerClient) Get(url string) (response *http.Response, err error) {
	// push a Get onto the queue
	in := input{
		url: url,
		// Get method
		method: 2,
		out:    make(chan output),
	}

	wc.requests <- in
	// return once it it done
	out := <-in.out
	return out.response, out.error
}
