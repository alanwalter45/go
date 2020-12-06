package main

import "fmt"

type operation struct {
	a, b int
}

func Modify(o *operation) (resultado int) {
	o.a *= 3
	return o.a + o.b
}

func main() {
	var (
		a = 5
		b = 10
	)

	o := &operation{
		a,
		b,
	}

	result := Modify(o)
	println(result)
	fmt.Println(o)
}
