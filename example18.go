package main

import "fmt"

func main() {
	var peoples = make(map[string]float32)
	peoples["Tonny"] = 18.3
	peoples["Luisa"] = 51

	fmt.Println(peoples)

	peoples["Ana"] = 32

	delete(peoples, "Luisa")

	for k, v := range peoples {
		fmt.Println(k, " : ", v)
	}
}
