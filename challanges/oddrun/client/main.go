package main

import (
	"math/rand"
	"strings"
	"time"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Body struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func RandomString(count int) string {
	var sb strings.Builder

	for len(sb.String()) != count {
		randomNumber := rand.Intn(len(letters))
		sb.WriteByte(letters[randomNumber])
	}

	return sb.String()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	getRequest(
		"http://localhost:8090",
		"mohammad",
	)

	// postRequest(
	// 	"http://localhost:8090",
	// 	&Body{
	// 		Key:   "mohammad",
	// 		Value: "nasr",
	// 	},
	// )

	// for {
	// 	postRequest(
	// 		"http://localhost:8090",
	// 		&Body{
	// 			Key:   RandomString(10),
	// 			Value: RandomString(10),
	// 		},
	// 	)

	// 	time.Sleep(time.Second)
	// }
}
