package main

import (
	"fmt"
	"io"
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
				closeResponse(response.Body)
				continue
			}
			if response.StatusCode == http.StatusOK {
				fmt.Println("[Consumer] got 200 OK: application started")
				closeResponse(response.Body)
				return true
			}

			closeResponse(response.Body)
			fmt.Println("[Consumer] app is not ready yet to consume events, continue to wait")
		}
	}
}

func closeResponse(responseBody io.Closer) {
	if err := responseBody; err != nil {
		fmt.Println("[Consumer] can not close the response")
	}
}
