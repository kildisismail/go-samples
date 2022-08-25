package main

import "fmt"

func main() {
	c := make(chan string)
	go func() { c <- "hello" }()
	msg := <-c
	fmt.Println(msg)

	//one way channel
	ch := make(chan string)
	go sc(ch)
	fmt.Println(<-ch)
}

func sc(ch chan<- string) {
	ch <- "\nhello"
}
