package main

import (
	"fmt"
	_ "io"
	"net/http"
)

type Ben struct {
	name map[string]int
}

func get() *http.Response {
	resp, _ := http.Get("https://example.com")
	return resp
}
func lose(m map[string]int) {
	//_ = body.Close()
	fmt.Println(m)
}

func main() {
	//resp := get()
	var ben = &Ben{map[string]int{"one":1}}
	lose(ben.name)
}
