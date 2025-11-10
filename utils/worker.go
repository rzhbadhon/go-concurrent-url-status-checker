package utils

import (
	"fmt"
	"go-url-checker/models"
	"net/http"
	"sync"
	"time"
)

func CheckUrl(url string, wg *sync.WaitGroup, ResultsChan chan<- models.Result) {
	defer wg.Done()

	// a client to check url for 5 sec
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil{
		// if the site is down of any issues
		ResultsChan <- models.Result{
			URL: url,
			Status: fmt.Sprintf("Down 🔴 (Error: %s)", err.Error()),
		}
		return
	}

	defer resp.Body.Close()
	// if success
	ResultsChan <- models.Result{
		URL: url,
		Status: fmt.Sprintf("Up ✅ (Status: %s)", resp.Status),
	}
}
