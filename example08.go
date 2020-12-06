package main

import (
	"fmt"
	"math/rand"
)

func change_item(s [4]int) {
	s[0] = 10
}

func main() {
	var a [4]int = [4]int{0, 0, 0, 0}

	change_item(a)
	fmt.Println(a)

	var matrix [][]int
	matrix = make([][]int, 3*3)

	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}

}
