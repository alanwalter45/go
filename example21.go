package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

func main() {
	message := "mensaje secreto"
	e := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println(e)
	d, _ := base64.StdEncoding.DecodeString(e)
	fmt.Println(string(d))

	dd, _ := hex.DecodeString(hex.EncodeToString([]byte(message)))
	fmt.Println("HEX>>>", string(dd))

	p := new(Car)
	p.Model = "z2018"
	p.Year = 2018

	data, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println(err)
	}
}

type Car struct {
	Model string `json:"model"`
	Year  int16  `json:"year"`
}
