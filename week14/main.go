package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Ltk bool
	Items []string
}

func main() {
	p1 := person{
		First: "Bill",
		Last: "Gates",
		Ltk: false,
		Items: []string{
			"Windows10",
			"Surface Pro 5",
			"IPhone 8",
		},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p1)
	fmt.Println("------------")
	fmt.Println(string(bs))
}
