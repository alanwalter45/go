package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text := "alanwalter"
	fmt.Println(text[4:])
	fmt.Println(strings.ToUpper(text))
	num := "10"
	v, _ := strconv.ParseInt(num, 10, 0)
	fmt.Println(v + 1)
}
