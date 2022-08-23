package main

import (
	"fmt"
)

//Pointers
func main() {
	var ap *int
	a := 35
	ap = &a
	fmt.Println(*ap)
	i := 10
	example(&i)
	fmt.Println(i)
}

//example
func example(i *int) {
	*i++
}