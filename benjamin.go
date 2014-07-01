package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultRequests = 5

func benchmark(url string, requests int) {
	start := time.Now()
	successfulRequests := 0
	failedRequests := 0
	totalRead := 0

	for i := 0; i < requests; i++ {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
			failedRequests++
			continue
		} else {
			successfulRequests++
			body, _ := ioutil.ReadAll(resp.Body)
			totalRead += len(body)
		}
	}

	elapsedTime := time.Since(start).Seconds()

  fmt.Printf("Total Requests: %d\n", requests)
  fmt.Printf("Sucessful Requests: %d\n", successfulRequests)

  if failedRequests > 0 {
    fmt.Printf("Failed Requests: %d\n", failedRequests)
  }

  fmt.Printf("Total transferred: %d [Bytes] received\n", totalRead)
  fmt.Printf("Transfer rate: %f [KBytes per second] received\n", float64(totalRead) / float64(1024) / float64(elapsedTime))
	fmt.Printf("Total time taken: %f seconds\n", elapsedTime)
	fmt.Printf("Requests per second: %f\n", float64(requests)/elapsedTime)
}

func main() {
	var url string
	var requests int

	flag.StringVar(&url, "url", "", "The URL to benchmark")
	flag.StringVar(&url, "u", "", "The URL to benchmark")
	flag.IntVar(&requests, "requests", defaultRequests, "The number of requests to perform on URL")
	flag.IntVar(&requests, "r", defaultRequests, "The number of requests to perform on URL")
	flag.Parse()

	fmt.Printf("Benchmarking requests to %s. Please wait...\n\n", url)
	benchmark(url, requests)
}
