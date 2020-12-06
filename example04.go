package main

import "fmt"

type Persona struct {
	name string
	age  uint8
}

func main() {
	/* data := Persona{
		name: "aw45",
		age:  15,
	} */

	data2 := Persona{
		"alanwalter45", 25,
	}

	fmt.Println(data2.name)
}
