package main

import (
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}

func main() {
	resp := get()
	_ = resp.Body.Close // want `must be closed`
}
