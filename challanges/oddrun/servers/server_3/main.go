package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	preBytes, _ := ioutil.ReadFile("out.txt")

	preBytes = append(preBytes, byte('\n'))
	preBytes = append(preBytes, body...)
	ioutil.WriteFile("out.txt", preBytes, 0644)
}
