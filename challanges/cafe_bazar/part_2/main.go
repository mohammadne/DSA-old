package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// https://stackoverflow.com/questions/26076165/go-find-scan-for-structs-functions

var (
	fatalError        = "it should not happen, log."
	errorInvalidInput = "Invalid input. Please enter 1 for more info."

	help = "Select a number from shown menu and enter. For example 1 is for help."
)

type citiesSlice []city

type roadsSlice []road

func (cities citiesSlice) add(c city) {
	_ = append(cities, c)
}

var cities citiesSlice

var roads roadsSlice

func main() {
	reader := bufio.NewReader(os.Stdin)

	action := processMainMenue(reader)
	switch action {
	case 2:
		processAdd(reader)
	case 3:
		processAdd(reader)
	case 4:
		processAdd(reader)
	default:
		log.Fatal(fatalError)
		os.Exit(1)
	}
}

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

// type model interface {
// 	add()
// 	delete()
// }

type city struct {
	id   int
	name string
}

type road struct {
	id            int
	name          string
	from          int
	to            int
	through       []int
	speedLimit    int
	length        int
	biDirectional bool
}

func printOptions(reader *bufio.Reader, header string, options []string) int {
	var sb strings.Builder

	sb.WriteString(header + "\n")

	length := len(options)

	for index, option := range options {
		txt := fmt.Sprintf("%d. %s\n", index+1, option)
		sb.WriteString(txt)
	}

	fmt.Print(sb.String())

	input, _ := reader.ReadString('\n')
	number, err := strconv.Atoi(strings.TrimSuffix(input, "\n"))

	if err != nil || number < 1 || number > length {
		fmt.Println(errorInvalidInput)
		return printOptions(reader, header, options)
	}

	return number
}

func getModel(reader *bufio.Reader) int {
	return printOptions(
		reader,
		"Select model:",
		[]string{"City", "Road"},
	)
}

func processMainMenue(reader *bufio.Reader) int {
	number := printOptions(
		reader,
		"Main Menu - Select an action:",
		[]string{"Help", "Add", "Delete", "Path", "Exit"},
	)

	// need for help
	if number == 1 {
		fmt.Println(help)
		return processMainMenue(reader)
	}

	// press Exit
	if number == 5 {
		os.Exit(0)
	}

	return number
}

func processAdd(reader *bufio.Reader, model int) {
	if model == 0 {
		model = getModel(reader)
	}

	switch model {
	case 1:
		var id int
		fmt.Scan(&id)

		var name string
		fmt.Scan(&name)

		city := city{
			id:   id,
			name: name,
		}

		cities.add(city)

		fmt.Printf("City with id=%d added!\n", id)

		number := printOptions(
			reader,
			"Select your next action",
			[]string{"Add another City", "Main manu"},
		)

		switch number {
		case 1:
			return processAdd(reader, 1)
		}
	}
}
