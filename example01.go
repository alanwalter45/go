package main

import "fmt"

func main() {
	a := 5

	fmt.Println(a)

	var (
		b string
		c uint
		d float32
	)

	b = "aw45"
	c = 100
	d = 5.5

	fmt.Printf("%s -  %d - %.2f\n", b, c, d)
}
