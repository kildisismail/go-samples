package main

import (
	"fmt"
)

//Interfaces
type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}

type dog struct {
	Type      string
	Dangerous bool
}

func (d dog) description() string {
	return fmt.Sprintf("Dangerous %v", d.Dangerous)
}

func (c cat) description() string {
	return fmt.Sprintf("Sound %v", c.Sound)
}

func main() {
	var a animal
	a = dog{Dangerous: true}
	fmt.Println(a.description())
	a = cat{Sound: "meow..!"}
	fmt.Println(a.description())
}
