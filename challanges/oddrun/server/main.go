package main

import (
	"log"
	"net/http"
	"os"
)

const (
	DirName = "storage"
)

func main() {
	createDirIfMissed(DirName)

	http.HandleFunc("/", Controller)
	url := "localhost" + ":" + os.Args[1]
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}
