package main

// import (
// 	"bufio"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// var (
// 	errorInvalidInput = "Invalid input. Please enter 1 for more info."

// 	help = "Select a number from shown menu and enter. For example 1 is for help."

// 	reader *bufio.Reader

// 	countries Countries
// 	cities    Cities
// 	roads     Roads
// )

// type Country struct {
// 	id   int
// 	name string
// }

// type City struct {
// 	id         int
// 	name       string
// 	country_id int
// }

// type Road struct {
// 	id            int
// 	name          string
// 	from          int
// 	to            int
// 	through       []int
// 	speedLimit    int
// 	length        int
// 	biDirectional bool
// }

// // type models interface {
// // 	add(model interface{})
// // 	delete(id int)
// // }

// type Cities []City

// func (cities Cities) add(model interface{}) Cities {
// 	city := model.(City)

// 	index := -1
// 	for i, c := range cities {
// 		if city.id == c.id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index >= 0 {
// 		cities[index] = city
// 	} else {
// 		cities = append(cities, city)
// 	}

// 	return cities
// }

// func (cities Cities) delete(id int) error {
// 	contains := false

// 	for i, city := range cities {
// 		if city.id == id {
// 			cities = append(cities[:i], cities[i+1:]...)
// 			contains = true
// 			break
// 		}
// 	}

// 	if !contains {
// 		return errors.New("")
// 	}

// 	return nil
// }

// type Roads []Road

// func (roads Roads) add(model interface{}) Roads {
// 	road := model.(Road)

// 	index := -1
// 	for i, r := range roads {
// 		if road.id == r.id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		return append(roads, road)
// 	}

// 	roads[index] = road
// 	return roads
// }

// func (roads Roads) delete(id int) error {
// 	contains := false

// 	for i, road := range roads {
// 		if road.id == id {
// 			roads = append(roads[:i], roads[i+1:]...)
// 			contains = true
// 			break
// 		}
// 	}

// 	if !contains {
// 		return errors.New("")
// 	}

// 	return nil
// }

// type Countries []Country

// func (countries Countries) export() string {
// 	var sb strings.Builder

// 	for _, country := range countries {
// 		str := fmt.Sprintf("%d\n%d\n%s\n%d", 2, country.id, country.name, 2)
// 		sb.WriteString(str)
// 	}

// 	return sb.String()
// }

// func (countries Countries) add(country Country) Countries {
// 	index := -1
// 	for i, c := range countries {
// 		if country.id == c.id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		return append(countries, country)
// 	}

// 	countries[index] = country
// 	return countries
// }

// func (countries Countries) delete(id int) error {
// 	contains := false

// 	for i, country := range countries {
// 		if country.id == id {
// 			countries = append(countries[:i], countries[i+1:]...)
// 			contains = true
// 			break
// 		}
// 	}

// 	if !contains {
// 		return errors.New("")
// 	}

// 	return nil
// }

// func (countries Countries) giveCountry(id int) *Country {
// 	var country Country

// 	for _, c := range countries {
// 		if c.id == id {
// 			country = c
// 			break
// 		}
// 	}

// 	return &country
// }

// func main() {
// 	reader = bufio.NewReader(os.Stdin)

// 	countries = make([]Country, 0, 20)
// 	cities = make([]City, 0, 20)
// 	roads = make([]Road, 0, 10)

// 	processMainMenue()
// }

// func readeLine() string {
// 	input, _ := reader.ReadString('\n')
// 	return strings.Trim(input, "\n")
// }

// func printOptions(header string, options []string) int {
// 	var sb strings.Builder

// 	sb.WriteString(header + "\n")

// 	length := len(options)

// 	for index, option := range options {
// 		txt := fmt.Sprintf("%d. %s\n", index+1, option)
// 		sb.WriteString(txt)
// 	}

// 	fmt.Print(sb.String())

// 	line := readeLine()
// 	fmt.Println(len(line))
// 	number, err := strconv.Atoi(line)
// 	// fmt.Println(number)

// 	if err != nil || number < 1 || number > length {
// 		// TODO
// 		// log.Fatal(err)
// 		fmt.Println(errorInvalidInput)
// 		return printOptions(header, options)
// 	}

// 	return number
// }

// func getModel() int {
// 	return printOptions(
// 		"Select model:",
// 		[]string{"City", "Road", "Country"},
// 	)
// }

// func processMainMenue() {
// 	number := printOptions(
// 		"Main Menu - Select an action:",
// 		[]string{"Help", "Add", "Delete", "Path", "Exit"},
// 	)

// 	switch number {
// 	case 1:
// 		fmt.Println(help)
// 		processMainMenue()
// 	case 2:
// 		model := getModel()

// 		switch model {
// 		case 1:
// 			addCity()
// 		case 2:
// 			addRoad()
// 		case 3:
// 			addCountry()
// 		}
// 	case 3:
// 		model := getModel()

// 		strId := readeLine()
// 		id, _ := strconv.Atoi(strId)

// 		switch model {
// 		case 1:
// 			err := cities.delete(id)
// 			if err != nil {
// 				fmt.Printf("City with id %d not found!\n", id)
// 			} else {
// 				fmt.Printf("City:%d deleted!\n", id)
// 			}

// 		case 2:
// 			err := roads.delete(id)
// 			if err != nil {
// 				fmt.Printf("Road with id %d not found!\n", id)
// 			} else {
// 				fmt.Printf("Road:%d deleted!\n", id)
// 			}

// 		case 3:
// 			err := countries.delete(id)
// 			if err != nil {
// 				fmt.Printf("Country with id %d not found!\n", id)
// 			} else {
// 				fmt.Printf("Country:%d deleted!\n", id)
// 			}
// 		}

// 		processMainMenue()
// 	case 4:
// 		showPath()
// 	case 5:
// 		os.Exit(0)
// 	}
// }

// func multiInput(values []string) []string {
// 	output := make([]string, 0, len(values))

// 	for _, value := range values {
// 		fmt.Printf("%s=?\n", value)
// 		output = append(output, readeLine())
// 	}

// 	return output
// }

// func addCity() {
// 	input := multiInput([]string{"id", "name", "country_id"})

// 	id, _ := strconv.Atoi(input[0])
// 	country_id, _ := strconv.Atoi(input[2])

// 	city := City{
// 		id:         id,
// 		name:       input[1],
// 		country_id: country_id,
// 	}

// 	cities = cities.add(city)

// 	fmt.Printf("City with id=%d added!\n", id)

// 	action := printOptions(
// 		"Select your next action",
// 		[]string{"Add another City", "Main Menu"},
// 	)

// 	switch action {
// 	case 1:
// 		addCity()
// 	case 2:
// 		processMainMenue()
// 	}
// }

// func addRoad() {
// 	input := multiInput(
// 		[]string{
// 			"id",
// 			"name",
// 			"from",
// 			"to",
// 			"‫‪through‬‬",
// 			"‫‪speed_limit‬‬",
// 			"‫‪length‬‬",
// 			"‫‪bi_directional‬‬",
// 		},
// 	)

// 	id, _ := strconv.Atoi(input[0])

// 	from, _ := strconv.Atoi(input[2])
// 	to, _ := strconv.Atoi(input[3])

// 	through := make([]int, 0, 5)
// 	input[4] = strings.Trim(input[4], "[")
// 	input[4] = strings.Trim(input[4], "]")
// 	for _, value := range strings.Split(input[4], ",") {
// 		number, _ := strconv.Atoi(value)
// 		through = append(through, number)
// 	}

// 	speedLimit, _ := strconv.Atoi(input[5])
// 	length, _ := strconv.Atoi(input[6])

// 	biDirectional, _ := strconv.Atoi(input[7])

// 	road := Road{
// 		id:            id,
// 		name:          input[1],
// 		from:          from,
// 		to:            to,
// 		through:       through,
// 		speedLimit:    speedLimit,
// 		length:        length,
// 		biDirectional: biDirectional == 1,
// 	}

// 	roads = roads.add(road)

// 	fmt.Printf("Road with id=%d added!\n", id)

// 	action := printOptions(
// 		"Select your next action",
// 		[]string{"Add another Road", "Main Menu"},
// 	)

// 	switch action {
// 	case 1:
// 		addRoad()
// 	case 2:
// 		processMainMenue()
// 	}
// }

// func addCountry() {
// 	input := multiInput([]string{"id", "name"})

// 	id, _ := strconv.Atoi(input[0])

// 	country := Country{
// 		id:   id,
// 		name: input[1],
// 	}

// 	countries = countries.add(country)

// 	fmt.Printf("Country with id=%d added!\n", id)

// 	action := printOptions(
// 		"Select your next action",
// 		[]string{"Add another Country", "Main Menu"},
// 	)

// 	switch action {
// 	case 1:
// 		addCountry()
// 	case 2:
// 		processMainMenue()
// 	}
// }

// func showPath() {
// 	values := strings.Split(readeLine(), ":")
// 	start, _ := strconv.Atoi(values[0])
// 	end, _ := strconv.Atoi(values[1])

// 	var sb strings.Builder

// 	for _, road := range roads {
// 		if road.biDirectional {
// 			index := indexOfId(road.through, start)
// 			if index != -1 {
// 				index2 := indexOfId(road.through[index:], end)
// 				if index2 != -1 {
// 					sb.WriteString(pathOut(road, index, index2+index))
// 				} else if road.to == end {
// 					sb.WriteString(pathOut(road, index, len(cities)-1))
// 				}
// 			} else if road.to == start {
// 				index2 := indexOfId(road.through, end)
// 				if index2 != -1 {
// 					sb.WriteString(pathOut(road, len(cities)-1, index2))
// 				}
// 			}
// 		} else {
// 			index := indexOfId(road.through, start)
// 			if index != -1 {
// 				index2 := indexOfId(road.through[index:], end)
// 				if index2 != -1 {
// 					sb.WriteString(pathOut(road, index, index2+index))
// 				} else if road.to == end {
// 					sb.WriteString(pathOut(road, index, len(cities)-1))
// 				}
// 			}
// 		}
// 	}

// 	fmt.Print(sb.String())
// 	processMainMenue()
// }

// func pathOut(road Road, from int, to int) string {
// 	startCity := cities[from]
// 	startCountry := countries.giveCountry(startCity.country_id)

// 	endCity := cities[to]
// 	endCountry := countries.giveCountry(endCity.country_id)

// 	time := dateTime(road.length, road.speedLimit)
// 	return fmt.Sprintf("%s/%s:%s/%s via Road %s: Takes %s\n", startCountry.name, startCity.name, endCountry.name, endCity.name, road.name, time)
// }

// // index == -1 means it's not exist
// func indexOfId(slice []int, id int) int {
// 	index := -1

// 	for i, value := range slice {
// 		if value == id {
// 			index = i
// 			break
// 		}
// 	}

// 	return index
// }

// func dateTime(length int, speedLimit int) string {
// 	days := (length / speedLimit) / 24
// 	hours := length/speedLimit - days*24
// 	mins := (float64(length%speedLimit) / float64(speedLimit)) * 60

// 	return fmt.Sprintf("%02d:%02d:%02d", days, hours, int64(mins))
// }
