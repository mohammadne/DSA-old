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
	floorsNum, _ := strconv.Atoi(line)

	for index := 0; index < floorsNum; index++ {
		floor := Floor{number: index, desks: make([]Desk, 0)}

		line := getInput(reader)
		desksNum, _ := strconv.Atoi(line)
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
		// requestType := values[1]
		user := values[2]
		duration, _ := strconv.Atoi(values[3])

		requestDesk(
			&Request{
				start:    start,
				user:     user,
				duration: duration,
			},
		)

	}
}

func requestDesk(req *Request) {

	for index1 := 0; index1 < len(floors); index1++ {
		desks := floors[index1].desks

		for index2 := 0; index2 < len(desks); index2++ {
			if desks[index2].request(req) {
				fmt.Printf("%s got desk %d-%d\n", req.user, index1+1, index2+1)
				return
			}
		}

	}

	fmt.Printf("no desk available\n")

}
