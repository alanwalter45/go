package main02

import (
	"fmt"
	"math"
	"math/rand"
)

func f1() {
	fmt.Println("El valor de PI es :", math.Round(math.Pi))
}

func rn(n int) int {
	return rand.Intn(n)
}

func add(x, y float32) float32 {
	return x + y
}

func main() {
	/* f1()
	fmt.Println("Numero aleatorio: ", rn(100)) */
	num1 := 5.3
	var num2 float32 = 5.2
	fmt.Println(add(float32(num1), num2))
}
