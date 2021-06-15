package main

import (
	"log"
	"net/http"
)

const (
	DirName = "cache"
)

var (
	servers = []http.Server{
		{Addr: "http://localhost:8091"},
		{Addr: "http://localhost:8092"},
		{Addr: "http://localhost:8093"},
		{Addr: "http://localhost:8094"},
	}
)

func main() {
	createDirIfMissed(DirName)

	http.HandleFunc("/", Controller)
	url := "localhost" + ":" + "8088"
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}
