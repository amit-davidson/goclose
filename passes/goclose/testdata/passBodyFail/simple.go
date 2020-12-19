package main

import (
	"io"
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(body io.ReadCloser) error {
	_ = body
	return nil
}

func main() {
	resp := get()
	body := resp.Body // want `must be closed`
	_ = closeBody(body)
}
