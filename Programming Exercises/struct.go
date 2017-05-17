package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func main() {

	p1 := person{"Sam", "Jackson"}
	p2 := person{"Sophie", "Engleberg"}
	fmt.Println(p1)
	fmt.Println(p2)
}
