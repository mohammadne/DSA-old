package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	barchasb := make([]rune, 5)

	for _, char := range input {
		barchasb = append(barchasb, char)
	}

	if ban1(barchasb) || ban2(barchasb) || ban3(barchasb) {
		fmt.Print("nakhor lite")
	} else {
		fmt.Print("rahat baash")
	}

}

// at most 2 red
func ban1(barchasb []rune) bool {
	redNum := 0

	for _, char := range barchasb {
		if char == 'R' {
			redNum++
		}
	}

	return redNum >= 3
}

// pass 2 red, 2 yellow
func ban2(barchasb []rune) bool {
	redNum := 0
	yellowNum := 0

	for _, char := range barchasb {
		if char == 'R' {
			redNum++
		} else if char == 'Y' {
			yellowNum++
		}
	}

	return (redNum >= 2 && yellowNum >= 2)
}

// all red and yellow
func ban3(barchasb []rune) bool {
	greenNum := 0

	for _, char := range barchasb {
		if char == 'G' {
			greenNum++
		}
	}

	return greenNum == 0
}
