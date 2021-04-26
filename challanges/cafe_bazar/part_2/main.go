package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	fatalError        = "it should not happen, log."
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

type models interface {
	add(model interface{})
	delete(id int)
}

type Cities []City

func (cities Cities) add(model interface{}) {
	cities = append(cities, model.(City))
}

func (cities Cities) delete(id int) {
}

type Roads []Road

func (roads Roads) add(model interface{}) {
	roads = append(roads, model.(Road))
}

func (roads Roads) delete(id int) {
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
		fmt.Println(help)
	case 4:
		fmt.Println(help)
	case 5:
		os.Exit(0)
	}

}

func multiInput(values []string) []string {
	output := make([]string, 0, len(values))

	for _, value := range values {
		fmt.Printf("%s=?\n", value)
		values = append(values, readeLine())
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

	fmt.Printf("City with id=%d added!", id)

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
	‫‪from‬‬, _ := strconv.Atoi(input[2])
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

	‫‪biDirectional‬‬, _ := strconv.Atoi(input[7])

	road := Road{
		id:      id,
		name:    input[1],
		from: from,
		to:      to,
		through: through,
		speedLimit: speedLimit,
		‫‪length‬‬: ‫‪length‬‬,
		biDirectional: ‫‪biDirectional‬‬ == 1,
		// from: ‫‪from‬‬,
	}

	roads.add(road)

	fmt.Printf("Road with id=%d added!", id)

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
