package main

import (
	"fmt"
	"time"
)

func main() {
	go c()
	fmt.Println("I am main")
	time.Sleep(time.Second * 2)
}
func c() {
	time.Sleep(time.Second * 2)
	fmt.Println("I am concurrent")
}
