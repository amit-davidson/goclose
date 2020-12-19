package main

import (
	"io"
	_ "io"
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(body io.ReadCloser) error {
	err := body.Close()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	resp := get()
	body := resp.Body
	_ = closeBody(body)
}
