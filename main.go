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
		
	}
