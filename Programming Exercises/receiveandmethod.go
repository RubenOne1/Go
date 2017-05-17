package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Moneypenny", 24}
	fmt.Println(p1)
	p1.foo()
	
}

func (p person) foo() {
	fmt.Println("Hello from foo")
}
