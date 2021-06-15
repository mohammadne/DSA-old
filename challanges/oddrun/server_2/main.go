package main

import (
	"log"
	"net/http"
)

const (
	DirName = "storage"
)

func main() {
	createDirIfMissed(DirName)

	http.HandleFunc("/", Controller)
	url := "localhost" + ":" + "9092"
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}
