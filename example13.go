package main

import (
	"fmt"
	"strings"
	"time"
)

type People struct {
	id      int
	name    string
	created time.Time
}

func main() {
	var p People
	p.id = 1
	p.name = "alanwalter45"
	p.created = time.Now()

	fmt.Println(p)

	var names string = "martin,maria,juana"

	p2 := new(People)
	p2.id = 2
	p2.name = strings.Split(names, ",")[1]
	p2.created = time.Now()

	fmt.Println(p2, len(p2.name))

}
