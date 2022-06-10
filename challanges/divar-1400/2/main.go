package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	user     string
	start    int
	duration int
	isFree   bool
}

type Desk struct {
	number   int
	requests []Request
}

func (d *Desk) request(req *Request) bool {
	canRequest := true

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
	floors []Floor
)

func main() {
	inputReciever()
}

func getInput(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}

func inputReciever() {
	reader := bufio.NewReader(os.Stdin)

	line := getInput(reader)
	values := strings.Split(line, " ")
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

		for index := 0; index < desksNum; index++ {
			desk := Desk{number: index, requests: make([]Request, 0)}
			floor.desks = append(floor.desks, desk)
		}

		floors = append(floors, floor)
	}

	for {
		line := getInput(reader)

		if line == "end" {
			break
		}

		values := strings.Split(line, " ")
		start, _ := strconv.Atoi(values[0])
		user := values[2]
		isFree := values[3] == "free"
		duration, _ := strconv.Atoi(values[4])

		requestDesk(
			&Request{
				start:    start,
				user:     user,
				duration: duration,
				isFree:   isFree,
			},
		)

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
				price := floors[index1].price
				if price == 0 {
					fmt.Printf("%s got desk %d-%d\n", req.user, index1+1, index2+1)
				} else {
					fmt.Printf("%s got desk %d-%d for %d\n", req.user, index1+1, index2+1, price)
				}

				return
			}
		}

	}

	fmt.Printf("no desk available\n")
}
