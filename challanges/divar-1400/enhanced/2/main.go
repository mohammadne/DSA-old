package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	freeFloorType    = "free"
	specialFloorType = "special"
)

var (
	reader *bufio.Reader

	gotDesk      = "%s got desk %d-%d"
	gotDeskPrice = "%s got desk %d-%d for %d"
	noDesk       = "no desk available"
)

type Request struct {
	user      string
	start     int
	duration  int
	isSpecial bool
}

func (req1 *Request) crossEachOther(req2 *Request) bool {
	start1 := req1.start
	end1 := start1 + req1.duration

	start2 := req2.start
	end2 := start2 + req2.duration

	startCond := start2 > start1 && start2 < end1
	endCond := end2 > start1 && end2 < end1
	containsCond := start2 <= start1 && end2 >= end1

	if startCond || endCond || containsCond {
		return false
	}

	return true
}

type Desk struct {
	number   int
	requests []Request
}

func (d *Desk) request(req *Request) bool {
	canRequest := true

	for index := 0; index < len(d.requests); index++ {
		canRequest = d.requests[index].crossEachOther(req)

		if !canRequest {
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

func (floor *Floor) request(req *Request) bool {
	cond1 := !req.isSpecial && floor.price != 0
	cond2 := req.isSpecial && floor.price == 0

	if cond1 || cond2 {
		return false
	}

	desks := floor.desks

	for index := 0; index < len(desks); index++ {
		if desks[index].request(req) {
			var text string
			if price := floor.price; price != 0 {
				text = fmt.Sprintf(gotDeskPrice, req.user, floor.number+1, index+1, price)
			} else {
				text = fmt.Sprintf(gotDesk, req.user, floor.number+1, index+1)
			}

			fmt.Println(text)
			return true
		}
	}

	return false
}

type Floors []Floor

func (floors Floors) request(req *Request) {
	for index := 0; index < len(floors); index++ {
		floor := floors[index]

		if canRequest := floor.request(req); canRequest {
			return
		}
	}

	fmt.Println(noDesk)
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	processInput()
}

func readLine() string {
	line, _ := reader.ReadString('\n')
	return strings.Trim(line, "\n")
}

func processInput() {
	line := readLine()
	values := strings.Split(line, " ")
	floorsNum, _ := strconv.Atoi(values[0])
	specialPrice, _ := strconv.Atoi(values[1])

	floors := make(Floors, 0, floorsNum)

	for index := 0; index < floorsNum; index++ {
		floor := Floor{number: index, desks: make([]Desk, 0)}

		line := readLine()
		values := strings.Split(line, " ")
		desksNum, _ := strconv.Atoi(values[0])
		floorType := values[1]

		if floorType == specialFloorType {
			floor.price = specialPrice
		}

		for index := 0; index < desksNum; index++ {
			desk := Desk{number: index, requests: make([]Request, 0)}
			floor.desks = append(floor.desks, desk)
		}

		floors = append(floors, floor)
	}

	for {
		line := readLine()

		if line == "end" {
			break
		}

		values := strings.Split(line, " ")
		start, _ := strconv.Atoi(values[0])
		user := values[2]
		floorType := values[3]
		duration, _ := strconv.Atoi(values[4])

		floors.request(
			&Request{
				start:     start,
				user:      user,
				duration:  duration,
				isSpecial: floorType == specialFloorType,
			},
		)

	}
}
