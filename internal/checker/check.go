package checker

import (
	"net/http"
	"time"
)

type CheckResult struct {
	Target string
	Status string
	Err    error
}

func CheckUrl(url string, results chan<- CheckResult) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		results <- CheckResult{Target: url, Status: "error", Err: err}
		return
	}
	defer resp.Body.Close()
	results <- CheckResult{Target: url, Status: resp.Status}
}
