package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("Other hello")
	}()
	fmt.Println("Hello World")
	var r stringf
	fmt.Scanln(&r)
	fmt.Print("...")
}
