package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"Todd": 45, "Nina": 25, "Patrick": 27}
	for k, _ := range x {
		fmt.Println(k, "-", v)
	}
}
