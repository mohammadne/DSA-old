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
	Cache bool   `json:"cache"`
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
		"http://localhost:8088",
		"amir",
	)

	// postRequest(
	// 	"http://localhost:8088",
	// 	&Body{
	// 		Key:   "amir",
	// 		Value: "noori",
	// 		Cache: true,
	// 	},
	// )

	// for {
	// 	postRequest(
	// 		"http://localhost:8088",
	// 		&Body{
	// 			Key:   RandomString(10),
	// 			Value: RandomString(10),
	// 			Cache: true,
	// 		},
	// 	)

	// 	time.Sleep(time.Second)
	// }
}
