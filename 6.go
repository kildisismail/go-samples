package main

import (
	"fmt"
)

//Functions
func add(a int, b int) int {
	c := a + b
	return c
}

func add2(a int, b int) (c int) {
	c = a + b
	return
}

func add3(a int, b int) (int, string) {
	c := a + b
	return c, "successfully added"
}

func main() {
	fmt.Println(add(2, 1))
	fmt.Println(add2(2, 1))
	sum, message := add3(2, 1)
	fmt.Println(message)
	fmt.Println(sum)
}
