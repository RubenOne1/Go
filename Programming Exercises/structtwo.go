package main

import (
	"fmt"
)

type person struct {
	fname string
	age   int
}

func main() {
	p1 := person{
		"Bond",
		32,
	}
	fmt.Println(p1.fname)
}
