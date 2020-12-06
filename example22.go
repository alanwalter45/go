package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	message := "hola a todos"
	data := md5.New()
	data.Write([]byte(message))
	data_md5 := hex.EncodeToString(data.Sum(nil))
	fmt.Println(data_md5)
}
