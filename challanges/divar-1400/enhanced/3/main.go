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

	reserveCommand = "reserve_desk"
	requestCommand = "request_desk"
)

var (
	reader *bufio.Reader

	requestDesk      = "%s got desk %d-%d"
	requestDeskPrice = "%s got desk %d-%d for %d"
	reserveDesk      = "%s reserved desk %d-%d for %d"
	noDesk           = "no desk available"
)

type Command struct {
	username  string
	timestamp int
	fromTime  int
	duration  int
}

func (com1 *Command) crossBoundaries(com2 *Command) bool {
	start1 := com1.timestamp
	end1 := com1.timestamp + com1.duration
	start2 := com2.timestamp
	end2 := com2.timestamp + com2.duration

	startCond := start2 > start1 && start2 < end1
	endCond := end2 > start1 && end2 < end1
	containsCond := start2 <= start1 && end2 >= end1

	if startCond || endCond || containsCond {
		return false
	}

	return true

}

type Reserve struct {
	command Command
}

type Request struct {
	command   Command
	isSpecial bool
}

type Desk struct {
	number   int
	commands []Command
}

func (d *Desk) command(command *Command) bool {
	for index := 0; index < len(d.commands); index++ {
		if cross := d.commands[index].crossBoundaries(command); !cross {
			return false
		}
	}

	d.commands = append(d.commands, *command)
	return true
}

type Floor struct {
	number int
	desks  []Desk
	price  int
}

func (floor *Floor) command(command *Command) int {
	for index := 0; index < len(floor.desks); index++ {

		if desk := floor.desks[index]; desk.command(command) {
			return desk.number
		}
	}

	return -1
}

func (floor *Floor) reserve(reserve *Reserve) int {
	if floor.price == 0 {
		return -1
	}

	return floor.command(&reserve.command)
}

func (floor *Floor) request(req *Request) int {
	cond1 := !req.isSpecial && floor.price != 0
	cond2 := req.isSpecial && floor.price == 0

	if cond1 || cond2 {
		return -1
	}

	return floor.command(&req.command)
}

type Floors []Floor

func (floors Floors) reserve(reserve *Reserve) {
	for floorIndex := 0; floorIndex < len(floors); floorIndex++ {
		floor := floors[floorIndex]

		if deskNumber := floor.reserve(reserve); deskNumber != -1 {
			var text string

			price := floor.price
			text = fmt.Sprintf(
				reserveDesk,
				reserve.command.username,
				floor.number,
				deskNumber,
				price,
			)

			fmt.Println(text)
			return
		}
	}

	fmt.Println(noDesk)
}

func (floors Floors) request(req *Request) {
	for floorIndex := 0; floorIndex < len(floors); floorIndex++ {
		floor := floors[floorIndex]

		if deskNumber := floor.request(req); deskNumber != -1 {
			var text string

			if price := floor.price; price != 0 {
				text = fmt.Sprintf(
					requestDeskPrice,
					req.command.username,
					floor.number,
					deskNumber,
					price,
				)
			} else {
				text = fmt.Sprintf(
					requestDesk,
					req.command.username,
					floor.number,
					deskNumber,
				)
			}

			fmt.Println(text)
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
		floor := Floor{number: index + 1, desks: make([]Desk, 0)}

		line := readLine()
		values := strings.Split(line, " ")
		desksNum, _ := strconv.Atoi(values[0])
		floorType := values[1]

		if floorType == specialFloorType {
			floor.price = specialPrice
		}

		for index := 0; index < desksNum; index++ {
			desk := Desk{number: index + 1, commands: make([]Command, 0)}
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
		timestamp, _ := strconv.Atoi(values[0])
		username := values[2]
		duration, _ := strconv.Atoi(values[4])

		command := Command{
			username:  username,
			timestamp: timestamp,
			fromTime:  timestamp,
			duration:  duration,
		}

		if values[1] == requestCommand {
			request := Request{
				command:   command,
				isSpecial: values[3] == specialFloorType,
			}

			floors.request(&request)
		} else {
			fromTime, _ := strconv.Atoi(values[3])
			command.fromTime = fromTime
			reserve := Reserve{command: command}

			floors.reserve(&reserve)
		}
	}
}
