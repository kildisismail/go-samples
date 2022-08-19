package main

import (
	"errors"
	"fmt"
)

//Custom Error handling

func Increment(n int) (int, error) {
	if n < 0 {
		//return obj err
		return nil, errors.New("math: cannot process negative number")
	}
	return (n + 1), nil
}

func main() {
	num := 5

	if inc, err := Increment(num); err != nil {
		fmt.Printf("Failed Number: %v, err msg: %v", num, err)
	} else {
		fmt.Printf("Increment num: %v", inc)
	}
}
