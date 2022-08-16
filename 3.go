package main

import (
	"fmt"
)

//Looping
func main() {
	fmt.Println("Starting...Hello, world!")
	toplam := 0
	for i := 0; i < 102; i++ {
		if i%2 == 0 {
			toplam += i
		}
	}
	fmt.Println("2 ile bölünebilenlerin toplamı=", toplam)
	i := 0
	toplam = 0
	for i < 100 {
		i++
		if i%3 == 0 {
			toplam += i
		}
	}
	fmt.Println("3 ile bölünebilenlerin toplamı=", toplam)
}
