package main

import "io/ioutil"

func create_file(text ...string) {
	filename := "test.txt"
	var data string
	for _, value := range text {
		data += " " + value
	}
	ioutil.WriteFile(filename, []byte(data), 0644)
}

func main() {
	create_file("hola", "desde", "GOLANG")
}
