package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://localhost:8000/about")
	bytes, _ := ioutil.ReadAll(response.Body)
	rs := string(bytes)
	response.Body.Close()
	fmt.Println(rs)
}
