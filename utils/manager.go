package utils

import (
	"fmt"
	"go-url-checker/utils"
	"sync"
)

func Manager() {
	urlsToChecks := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://golang.org",
		"https://doesntexistdomainever.xyz", // random
		"https://www.youtube.com",
	}

	// creating channel and waitgroup
	jobs := make(chan string, len(urlsToChecks))
	results := make(chan string, len(urlsToChecks))
	var wg sync.WaitGroup

	// set workers
	numWorkers := 3

	// set wg for workers
	wg.Add(numWorkers)
	for w := 1; w <= numWorkers; w++ {
		go utils.Worker(&wg, w, jobs, results)
	}

	// send all job to jobs channel
	for _, url := range urlsToChecks {
		jobs <- url
	}

	// sending done now close jobs as no new task will be there...
	close(jobs)

	// another goroutine to close the results
	go func() {
		wg.Wait()      // waits till workers finish the job
		close(results) // closes after all works done
	}()

	// lets print all results
	fmt.Println("Results: ")
	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("All done...")

}
