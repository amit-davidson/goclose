package main

import (
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(close func() error) error {
	return close()
}

func main() {
	resp := get()
	closeFunc := resp.Body.Close
	_ = closeBody(closeFunc)
}
