package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"unsafe"
)

type Body struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Cache bool   `json:"cache"`
}

func Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetObjectController(w, r)
	case http.MethodPost:
		PutObjectController(w, r)
	}
}

func GetObjectController(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	id, parseErr := strconv.ParseInt(segments[len(segments)-1], 10, 64)
	if parseErr != nil {
		code := http.StatusBadRequest
		http.Error(w, http.StatusText(code), code)
		return
	}

	strId := strconv.FormatInt(id, 10)
	path := DirName + "/" + strId

	if !isFileExists(path) {
		remain := id % int64(len(servers))
		getRequest(servers[remain].Addr + "/" + strId)
		return
	}

	bytes, err := readFile(path)
	if err != nil {
		panic(err)
	}

	values := strings.Split(string(bytes), "\n")

	body := Body{
		Key:   values[0],
		Value: values[1],
	}

	bytesBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytesBody)
}

func getRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func PutObjectController(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	b := Body{}
	err := json.Unmarshal(body, &b)
	if err != nil {
		panic(err)
	}

	if b.Cache {
		md5Key := md5.Sum([]byte(b.Key))
		byteToInt := ByteArrayToInt(md5Key)

		path := DirName + "/" + strconv.FormatInt(byteToInt, 10)
		if !isFileExists(path) {
			file, err := createFile(path)
			if err != nil {
				panic("")
			}
			defer file.Close()

			file.WriteString(b.Key)
			file.WriteString("\n")
			file.WriteString(b.Value)
		}
	}

	md5Key := md5.Sum([]byte(b.Key))
	byteToInt := ByteArrayToInt(md5Key)

	remain := byteToInt % int64(len(servers))
	if remain < 0 {
		remain *= -1
	}

	postRequest(servers[remain].Addr, b)
}

func postRequest(url string, body Body) {
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
	defer resp.Body.Close()
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
