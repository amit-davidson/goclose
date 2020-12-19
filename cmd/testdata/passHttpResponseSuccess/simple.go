package main

import (
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(resp *http.Response) error {
	err := resp.Body.Close()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	resp := get()
	_ = closeBody(resp)
}
