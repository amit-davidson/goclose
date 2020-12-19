package main

import (
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(resp *http.Response) error {
	return resp.Body.Close()
}

func main() {
	resp := get()
	_ = closeBody(resp)
}
