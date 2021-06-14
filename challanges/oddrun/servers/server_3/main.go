package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	filename := "out.txt"

	file, err := getFile(filename)
	if err != nil {
		panic(err.Error())
	}

	file.Write(append(body, byte('\n')))
}

func getFile(filename string) (*os.File, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return os.Create(filename)
	}

	if info.IsDir() {
		return nil, fmt.Errorf("IT'S A DIRECTORY, PROVIDE A FILE")
	}

	return os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
}
