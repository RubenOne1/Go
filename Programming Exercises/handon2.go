package main

import "fmt"

type person struct {
	fName string
	lName string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("uptown, func you up, uptown func you up")
}

func (sa secretAgent) speak() {
	fmt.Println("shaken, not stirred")
}

func main() {
	p1 := person{"Nina", "Simone", 25}
	sa1 := secretAgent{person{"Ian", "Fleming", 42}, false}
	fmt.Println(p1)
	fmt.Println(sa1)
	p1.speak()
	sa1.speak()
	sa1.person.speak()
}
