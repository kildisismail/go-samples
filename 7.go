package main

import (
	"fmt"
)

//Struct

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) describe() {
	fmt.Printf("%v is %v years old.", p.name, p.age)
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p person) setName(name string) {
	p.name = name
}

func main() {
	pp := &person{name: "John", age: 1, gender: "Male"}
	pp.describe()
	pp.setAge(11)
	fmt.Println("\n", pp.age)
	pp.setName("Johnson")
	fmt.Println(pp.name)
}
