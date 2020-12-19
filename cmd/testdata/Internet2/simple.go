package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	isAppReady()
}

func isAppReady() bool {
	shutdownChannel := make(chan bool)
	appCheckFrequency := 5
	appHost := ""
	appHealthCheckURI := ""
	ticker := time.NewTicker(time.Duration(appCheckFrequency) * time.Second)
	for {
		select {
		case <-shutdownChannel:
			fmt.Println("[Consumer] got Shutdown signal, terminating app health check")
			return false
		case <-ticker.C:
			response, err := http.Get(appHost + appHealthCheckURI)
			if err != nil {
				fmt.Println("[Consumer] error after the app health check request")
				closeResponse(response)
				continue
			}
			if response.StatusCode == http.StatusOK {
				fmt.Println("[Consumer] got 200 OK: application started")
				closeResponse(response)
				return true
			}

			closeResponse(response)
			fmt.Println("[Consumer] app is not ready yet to consume events, continue to wait")
		}
	}
}

func closeResponse(resp *http.Response) {
	if err := resp.Body; err != nil {
		fmt.Println("[Consumer] can not close the response")
	}
}
