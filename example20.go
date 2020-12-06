package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func main() {
	go m1()
	go m2()

	var n string
	fmt.Scanln(&n)
}

func m1() {
	for i := 0; i < 5; i++ {
		lock.Lock()
		fmt.Println(i + 1)
		lock.Unlock()
		time.Sleep(3 * time.Second)
	}
}

func m2() {
	for i := 0; i < 5; i++ {
		lock.Lock()
		fmt.Println(i + 100)
		lock.Unlock()
		time.Sleep(1 * time.Second)
	}
}
