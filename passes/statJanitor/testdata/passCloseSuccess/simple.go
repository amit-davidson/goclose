package main

import (
	"fmt"
	"net/http"
)

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func closeBody(close func() error) error {
	fmt.Println("asda")
	fmt.Println("asda")
	fmt.Println("asda")
	fmt.Println("asda")
	return nil
	//return close()
}

func main() {
	resp := get()
	closeFunc := resp.Body.Close
	_ = closeBody(closeFunc)
}
