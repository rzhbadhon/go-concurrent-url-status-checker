package main

import (
	"fmt"
	"net/http"
	"sync"
)

func worker(wg *sync.WaitGroup, id int, jobs <-chan string, results chan<- string) {
	defer wg.Done()

	// recieve url from jobs channel
	// this repeats until the channel close
	for url := range jobs {
		fmt.Printf("Worker %d working job: %s\n", id, url)

		// we send get req
		resp, err := http.Get(url)

		// sending results into the results channel
		if err != nil {
			// if the url is down
			results <- fmt.Sprintf("[Failed] %s -> %s", url, err.Error())

		} else {
			//if successful
			results <- fmt.Sprintf("[Passed] %s -> Status: %s", url, resp.Status)
			resp.Body.Close()

		}

	}

}


func main(){
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
	for w:=1; w <= numWorkers; w++{
		go worker(&wg, w, jobs, results)
	}

	// send all job to jobs channel
	for _, url := range urlsToChecks{
		jobs <- url
	}

	// sending done now close jobs as no new task will be there...
	close(jobs)

	// another goroutine to close the results
	go func ()  {
		wg.Wait() // waits till workers finish the job
		close(results)// closes after all works done	
	}()

	// lets print all results
	fmt.Println("Results: ")
	for res := range results{
		fmt.Println(res)
	}

	fmt.Println("All done...")

}
