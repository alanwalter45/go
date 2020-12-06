package main

import "fmt"

func main() {
	a := make([][]int, 2)

	a[0] = make([]int, 2)
	a[0][0] = 1
	a[0][1] = 2
	a[1] = make([]int, 2)
	a[1][0] = 3
	a[1][1] = 4

	for index, value := range a {
		fmt.Println(index)
		fmt.Println(value)
	}

}
