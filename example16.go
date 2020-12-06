package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func mlog(message string) {
	file, error := os.OpenFile("./mylog.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if error != nil {
		fmt.Println("no se creo el archivo de log")
	}
	log2 := log.New(file, "alan:", log.Ldate)
	log2.Println(message)
}

func main() {
	content, _ := ioutil.ReadFile("./test.txt")
	fmt.Println(string(content))
	log.Println("iniciando log...")
	mlog("utilizando log en la aplicacion")
	mlog("simplemente para comprobar")
}
