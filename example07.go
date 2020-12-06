package main

import "fmt"

func main() {
	var i uint = 0 // escribiendo un comentario

	for true {

		i++

		if i == 3 {
			continue
		}

		fmt.Println(i)

		if i > 10 {
			break
		}
	}
}
