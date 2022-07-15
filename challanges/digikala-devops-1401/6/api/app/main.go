package main

import "log"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var Users []User

func main() {
	if err := NewServer(); err != nil {
		log.Fatal(err.Error())
	}

	blocker := make(chan struct{})
	<-blocker
}
