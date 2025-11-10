package rest

import (
	"encoding/json"
	"fmt"
	"go-url-checker/models"
	"go-url-checker/utils"
	"net/http"
	"sync"
)

func Verify(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody models.RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	urls := reqBody.Urls
	if len(urls) == 0 {
		http.Error(w, "No URLs provided", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup

	ResultsChan := make(chan models.Result, len(urls))
	finalResults := make(map[string]string)

	// run goroutine for each worker
	for _, url := range urls {
		wg.Add(1)
		go utils.CheckUrl(url, &wg, ResultsChan)
	}

	// this go routine will wait till the wait() is complete
	go func() {
		wg.Wait()          // wait until checkurl finishes all go routine
		close(ResultsChan) // wil close the result channel then
	}()

	// collect results from result channel
	for res := range ResultsChan {
		finalResults[res.URL] = res.Status
	}

	// send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResults)
	fmt.Println("Request processed successfully")

}
