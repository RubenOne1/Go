package main

import (
	"fmt"
)

type person struct {
	fname string
	age   int
	prefs []string
}

func main() {
	p1 := person{
		"Bond",
		32,
		[]string{"one", "another", "last", "lied"},
	}
	fmt.Println(p1.prefs)
}
