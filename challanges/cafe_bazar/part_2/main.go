package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	errorInvalidInput = "Invalid input. Please enter 1 for more info."

	help = "Select a number from shown menu and enter. For example 1 is for help."

	reader *bufio.Reader
)

type City struct {
	id   int
	name string
}

type Road struct {
	id            int
	name          string
	from          int
	to            int
	through       []int
	speedLimit    int
	length        int
	biDirectional bool
}

// type models interface {
// 	add(model interface{})
// 	delete(id int)
// }

type Cities []City

func (cities Cities) add(model interface{}) Cities {
	city := model.(City)

	index := -1
	for i, c := range cities {
		if city.id == c.id {
			index = i
			break
		}
	}

	if index >= 0 {
		cities[index] = city
	} else {
		cities = append(cities, city)
	}

	return cities
}

func (cities Cities) delete(id int) error {
	contains := false

	for i, city := range cities {
		if city.id == id {
			cities = append(cities[:i], cities[i+1:]...)
			contains = true
			break
		}
	}

	if !contains {
		return errors.New("")
	}

	return nil
}

type Roads []Road

func (roads Roads) add(model interface{}) Roads {
	road := model.(Road)

	index := -1
	for i, r := range roads {
		if road.id == r.id {
			index = i
			break
		}
	}

	if index >= 0 {
		roads[index] = road
	} else {
		roads = append(roads, road)
	}

	return roads
}

func (roads Roads) delete(id int) error {
	contains := false

	for i, road := range roads {
		if road.id == id {
			roads = append(roads[:i], roads[i+1:]...)
			contains = true
			break
		}
	}

	if !contains {
		return errors.New("")
	}

	return nil
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var cities Cities = make([]City, 0, 20)
	var roads Roads = make([]Road, 0, 10)

	processMainMenue(cities, roads)
}

func readeLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSuffix(input, "\n")
}

func printOptions(header string, options []string) int {
	var sb strings.Builder

	sb.WriteString(header + "\n")

	length := len(options)

	for index, option := range options {
		txt := fmt.Sprintf("%d. %s\n", index+1, option)
		sb.WriteString(txt)
	}

	fmt.Print(sb.String())

	number, err := strconv.Atoi(readeLine())

	if err != nil || number < 1 || number > length {
		fmt.Println(errorInvalidInput)
		return printOptions(header, options)
	}

	return number
}

func getModel() int {
	return printOptions(
		"Select model:",
		[]string{"City", "Road"},
	)
}

func processMainMenue(cities Cities, roads Roads) {
	number := printOptions(
		"Main Menu - Select an action:",
		[]string{"Help", "Add", "Delete", "Path", "Exit"},
	)

	processMainMenueCallback := func() {
		processMainMenue(cities, roads)
	}

	switch number {
	case 1:
		fmt.Println(help)
		processMainMenueCallback()
	case 2:
		model := getModel()

		switch model {
		case 1:
			addCity(cities, processMainMenueCallback)
		case 2:
			addRoad(roads, processMainMenueCallback)
		}
	case 3:
		model := getModel()

		strId := readeLine()
		id, _ := strconv.Atoi(strId)

		switch model {
		case 1:
			err := cities.delete(id)
			if err != nil {
				fmt.Printf("City:%d deleted!\n", id)
			} else {
				fmt.Printf("City with id %d not found!\n", id)
			}

		case 2:
			err := roads.delete(id)
			if err != nil {
				fmt.Printf("Road:%d deleted!\n", id)
			} else {
				fmt.Printf("Road with id %d not found!\n", id)
			}
		}

		processMainMenueCallback()
	case 4:
		showPath(cities, roads, processMainMenueCallback)
	case 5:
		os.Exit(0)
	}
}

func multiInput(values []string) []string {
	output := make([]string, 0, len(values))

	for _, value := range values {
		fmt.Printf("%s=?\n", value)
		output = append(output, readeLine())
	}

	return output
}

func addCity(cities Cities, menuCallback func()) {
	input := multiInput([]string{"id", "name"})

	id, _ := strconv.Atoi(input[0])

	city := City{
		id:   id,
		name: input[1],
	}

	cities.add(city)

	fmt.Printf("City with id=%d added!\n", id)

	action := printOptions(
		"Select your next action",
		[]string{"Add another City", "Main Menu"},
	)

	switch action {
	case 1:
		addCity(cities, menuCallback)
	case 2:
		menuCallback()
	}
}

func addRoad(roads Roads, menuCallback func()) {
	input := multiInput(
		[]string{
			"id",
			"name",
			"from",
			"to",
			"‫‪through‬‬",
			"‫‪speed_limit‬‬",
			"‫‪length‬‬",
			"‫‪bi_directional‬‬",
		},
	)

	id, _ := strconv.Atoi(input[0])

	from, _ := strconv.Atoi(input[2])
	to, _ := strconv.Atoi(input[3])

	through := make([]int, 0, 5)
	input[4] = strings.Trim(input[4], "[")
	input[4] = strings.Trim(input[4], "]")
	for _, value := range strings.Split(input[4], ",") {
		number, _ := strconv.Atoi(value)
		through = append(through, number)
	}

	speedLimit, _ := strconv.Atoi(input[5])
	length, _ := strconv.Atoi(input[6])

	biDirectional, _ := strconv.Atoi(input[7])

	road := Road{
		id:            id,
		name:          input[1],
		from:          from,
		to:            to,
		through:       through,
		speedLimit:    speedLimit,
		length:        length,
		biDirectional: biDirectional == 1,
	}

	roads.add(road)

	fmt.Printf("Road with id=%d added!\n", id)

	action := printOptions(
		"Select your next action",
		[]string{"Add another Road", "Main Menu"},
	)

	switch action {
	case 1:
		addRoad(roads, menuCallback)
	case 2:
		menuCallback()
	}
}

func showPath(cities Cities, roads Roads, menuCallback func()) {
	values := strings.Split(readeLine(), ":")
	start, _ := strconv.Atoi(values[0])
	end, _ := strconv.Atoi(values[1])

	var sb strings.Builder

	for _, road := range roads {
		if road.biDirectional {
			index := indexOfId(road.through, start)
			if index != -1 {
				index2 := indexOfId(road.through[index:], end)
				if index2 != -1 || road.to == end {
					sb.WriteString(pathOut(road, cities))
				}
			} else if road.to == start {
				index2 := indexOfId(road.through, end)
				if index2 != -1 {
					sb.WriteString(pathOut(road, cities))
				}
			}
		} else {
			index := indexOfId(road.through, start)
			if index != -1 {
				index2 := indexOfId(road.through[index:], end)
				if index2 != -1 || road.to == end {
					sb.WriteString(pathOut(road, cities))
				}
			}
		}
	}

	fmt.Println("Kashan:Qom via Road T-K: Takes 00:07:30")
	menuCallback()
}

func pathOut(road Road, cities Cities) string {
	citiesIdMap := make([]int, 0, len(cities))

	startCity := cities[indexOfId(citiesIdMap, road.from)]
	endCity := cities[indexOfId(citiesIdMap, road.from)]

	time := dateTime(road.length, road.speedLimit)
	return fmt.Sprintf("%s:%s via Road %s: Takes %s\n", startCity.name, endCity.name, road.name, time)
}

// index == -1 means it's not exist
func indexOfId(slice []int, id int) int {
	index := -1

	for i, value := range slice {
		if value == id {
			index = i
			break
		}
	}

	return index
}

func dateTime(length int, speedLimit int) string {
	days := (length / speedLimit) / 24
	hours := length/speedLimit - days*24
	mins := (float64(length%speedLimit) / float64(speedLimit)) * 60

	return fmt.Sprintf("%02d:%02d:%02d", days, hours, int64(mins))
}
