package main

import (
	"fmt"
)

type person struct {
	first string
}

func (p person) speak(){
	fmt.Println(p.first,"I am a person")
}

func main(){
	fmt.Println("Hello, all")
	p := person{"Davis"}
	p.speak()
	p1 := person{"Sam"}
	p1.speak()
	x := []int{7, 9, 42}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}
	m := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	fmt.Println(m)
}
