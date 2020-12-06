package main

import "fmt"

type Dream struct {
	name string
	age  uint8
}

func (d Dream) get_information() string {
	return fmt.Sprintf("name: %s age: %d", d.name, d.age)
}

func (d *Dream) set_name(name string) {
	d.name = name
}

func main() {
	d := Dream{
		name: "aw45",
		age:  15,
	}
	d.set_name("alanwalter45")
	fmt.Println(d.get_information())
}
