package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
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

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		postRequest(
			"http://localhost:8090",
			&Body{
				Key:   RandomString(10),
				Value: RandomString(10),
			},
		)

		time.Sleep(time.Second)
	}
}

func postRequest(url string, body interface{}) {
	byteBody, _ := json.Marshal(&body)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteBody))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("url:%s, status:%s\n", url, resp.Status)

	defer resp.Body.Close()
}

func RandomString(count int) string {
	var sb strings.Builder

	for len(sb.String()) != count {
		randomNumber := rand.Intn(len(letters))
		sb.WriteByte(letters[randomNumber])
	}

	return sb.String()
}
