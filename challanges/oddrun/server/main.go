package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const (
	DirName = "storage"
)

func main() {
	createDirIfMissed(DirName)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	serverUrl := os.Getenv("SERVER_URL")
	serverPort := os.Getenv("SERVER_PORT")

	http.HandleFunc("/", Controller)

	fullUrl := serverUrl + ":" + serverPort
	fmt.Printf("Starting HTTP server on %s\n", fullUrl)
	if err := http.ListenAndServe(fullUrl, nil); err != nil {
		log.Fatal(err)
	}
}
