package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := inputReciever()

	for _, line := range values {
		isRound := false

		isRound = rule_1(line) || rule_2(line) || rule_3(line)

		if isRound {
			fmt.Println("Ronde!")
		} else {
			fmt.Println("Rond Nist")
		}
	}
}

func rule_1(value string) bool {
	res := false

	for index1 := 0; index1 < 8; index1++ {
		tekrar := 0
		for index2 := index1; index2 < 8; index2++ {
			if value[index1] == value[index2] {
				tekrar++
			}
		}

		if tekrar >= 4 {
			res = true
			break
		}
	}

	return res
}

func rule_2(value string) bool {
	res := false

	for index := 0; index < 6; index++ {
		if value[index] == value[index+1] && value[index+1] == value[index+2] {
			res = true
			break
		}
	}

	return res
}

func rule_3(value string) bool {
	res := true

	for index := 0; index < 4; index++ {
		if value[index] != value[7-index] {
			res = false
			break
		}
	}

	return res
}

// inputReciever gives unformatted lines
func inputReciever() []string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	linesNumber, err := strconv.Atoi(scanner.Text())

	if err != nil {
		// panic("wrong input")
		// fmt.Println("wrong formatted text given")
		os.Exit(1)
	}

	// to increase performance we will
	// alocate a slice with cap=10 and length=0
	// to easily append to it
	lines := make([]string, 0, 50)

	for lineIndex := 0; lineIndex < linesNumber; lineIndex++ {
		scanner.Scan()
		lines = append(lines, scanner.Text())
	}

	return lines
}
