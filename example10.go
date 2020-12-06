package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func suma_multipe(values ...int) int {
	var suma int = 0
	for _, value := range values {
		suma += value
	}
	return suma
}

func change_string(cadena string, value, value2 int) string {
	return cadena + fmt.Sprintf("=> %d , %d", value, value2)
}

func main() {
	resultado := suma(5, 6)
	fmt.Println(resultado)
	fmt.Println(suma_multipe(1, 2, 3, 4, 5))

	message := "frase"
	value, value2 := 10, 5
	fmt.Println(change_string(message, value, value2))
}
