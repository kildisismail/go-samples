package main

import (
	"fmt"
	"net/http"
)

//Error handling
func main() {
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(resp)
}
