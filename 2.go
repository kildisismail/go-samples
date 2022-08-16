package main

import (
	"fmt"
)

//Variables
var b, c int = 2, 3

func main() {
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	//Type casting
	a := int(b)
	fmt.Println("a:", a)
}
