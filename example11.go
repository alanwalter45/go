package main

import "fmt"

func print2() {
	fmt.Print("algo\n")
}

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {
	print2()
	r := fibo(6)
	fmt.Println(r)
}
