package main

import "fmt"

func main() {
	a := 50

	fmt.Println(a)
	fmt.Println(&a)

	b := &a
	*b = 10

	fmt.Println(*b, a)

	var c *int
	c = &a
	*c = 11
	fmt.Println(*c, "-", a, "-", *b)

	d := new(int)
	/* bla bla bla */
}
