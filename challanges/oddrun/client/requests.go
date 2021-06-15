package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"unsafe"
)

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

func getRequest(url string, key string) {
	md5Key := md5.Sum([]byte(key))
	byteToInt := ByteArrayToInt(md5Key)
	strKey := strconv.FormatInt(byteToInt, 10)

	resp, err := http.Get(url + "/" + strKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		b := Body{}
		err = json.Unmarshal(body, &b)
		if err != nil {
			panic(err)
		}

		fmt.Printf("key:%s, value:%s\n", b.Key, b.Value)
	} else {
		fmt.Printf("status:%s\n", resp.Status)
	}
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
