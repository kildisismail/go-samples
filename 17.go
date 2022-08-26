package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	ch <- "!" // extra message in buffer
	fmt.Println(<-ch)
}
