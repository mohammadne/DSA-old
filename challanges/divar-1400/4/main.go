package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reserve struct {
	user     string
	start    int
	duration int
	features []int
}

type Request struct {
	user      string
	start     int
	duration  int
	isFree    bool
	isRequest bool
	features  []int
}

type Desk struct {
	number   int
	requests []Request
	features []int
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (d *Desk) request(req *Request) bool {
	canRequest := true

	if req.features != nil && !testEq(d.features, req.features) {
		canRequest = false
		return false
	}

	for index := 0; index < len(d.requests); index++ {
		start := d.requests[index].start
		end := start + d.requests[index].duration

		startCond := req.start > start && req.start < end
		endCond := req.start+req.duration > start && req.start+req.duration < end
		containsCond := req.start <= start && req.start+req.duration >= end

		if startCond || endCond || containsCond {
			canRequest = false
			break
		}
	}

	if canRequest {
		d.requests = append(d.requests, *req)
	}

	return canRequest
}

type Floor struct {
	number int
	desks  []Desk
	price  int
}

var (
	floors   []Floor
	features map[int]int
)

func main() {
	features = map[int]int{}
	inputReciever()
}

func getInput(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}

func inputReciever() {
	reader := bufio.NewReader(os.Stdin)

	line := getInput(reader)
	featuresNum, _ := strconv.Atoi(line)

	line = getInput(reader)
	values := strings.Split(line, " ")
	for index := 0; index < featuresNum; index++ {
		price, _ := strconv.Atoi(values[index])
		features[index] = price
	}

	line = getInput(reader)
	values = strings.Split(line, " ")
	floorsNum, _ := strconv.Atoi(values[0])
	price, _ := strconv.Atoi(values[1])

	for index := 0; index < floorsNum; index++ {
		floor := Floor{number: index, desks: make([]Desk, 0)}

		line := getInput(reader)
		values := strings.Split(line, " ")
		desksNum, _ := strconv.Atoi(values[0])
		isFree := values[1] == "free"

		if !isFree {
			floor.price = price
		}

		line = getInput(reader)
		values = strings.Split(line, " ")

		for index := 0; index < desksNum; index++ {
			features := stringToFeatures(values[index])
			desk := Desk{number: index, requests: make([]Request, 0), features: features}
			floor.desks = append(floor.desks, desk)
		}

		floors = append(floors, floor)
	}

	for {
		line := getInput(reader)

		if line == "end" {
			break
		}

		request := Request{}

		values := strings.Split(line, " ")

		if values[1] == "request_desk" {
			request.start, _ = strconv.Atoi(values[0])
			request.isRequest = true
			request.user = values[2]
			request.isFree = values[3] == "free"
			request.duration, _ = strconv.Atoi(values[4])
		} else {
			request.isRequest = false
			request.user = values[2]
			request.start, _ = strconv.Atoi(values[3])
			request.isFree = false
			request.duration, _ = strconv.Atoi(values[4])
			request.features = stringToFeatures(values[5])
		}

		requestDesk(&request)

	}
}

func requestDesk(req *Request) {
	for index1 := 0; index1 < len(floors); index1++ {
		desks := floors[index1].desks

		contra1 := req.isFree && floors[index1].price != 0
		contra2 := !req.isFree && floors[index1].price == 0
		if contra1 || contra2 {
			continue
		}

		for index2 := 0; index2 < len(desks); index2++ {
			canRequest := desks[index2].request(req)
			if canRequest {
				floorPrice := floors[index1].price

				featurePrice := 0
				for index := 0; index < len(desks[index2].features); index++ {
					featurePrice += features[desks[index2].features[index]]
				}

				price := floorPrice + req.duration*featurePrice

				if req.isRequest {
					fmt.Printf("%s got desk %d-%d for %d\n", req.user, index1+1, index2+1, price)
				} else {
					fmt.Printf("%s reserved desk %d-%d for %d\n", req.user, index1+1, index2+1, price)
				}

				return
			}
		}

	}

	fmt.Printf("no desk available\n")
}

func stringToFeatures(input string) []int {
	result := make([]int, 0)

	for index := 0; index < len(input); index++ {
		if string(input[index]) == "1" {
			result = append(result, index)
		}
	}

	return result
}
