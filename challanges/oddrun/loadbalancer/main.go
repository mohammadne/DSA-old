package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"unsafe"
)

var (
	serversPorts = []int{
		8081,
		8092,
		8083,
		8084,
	}
)

type BodyScheme struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	b := BodyScheme{}
	err := json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}

	md5Key := md5.Sum([]byte(b.Key))
	byteToInt := ByteArrayToInt(md5Key)
	remain := byteToInt % int64(len(serversPorts))

	request(fmt.Sprintf("http://localhost:%d", serversPorts[remain]), r.Body)
}

func ByteArrayToInt(arr [16]byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}

	if val < 0 {
		val = val * -1
	}

	return val

}

func request(url string, body io.Reader) {
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
