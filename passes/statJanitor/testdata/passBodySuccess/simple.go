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
	return body.Close()
}

func main() {
	resp := get()
	body := resp.Body
	_ = closeBody(body)
}
