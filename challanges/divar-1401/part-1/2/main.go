package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line)

	char := line[0]
	counts := 1

	var sb strings.Builder
	for index := 1; index < len(line); index++ {
		newChar := line[index]

		if newChar != char {
			txt := strconv.Itoa(counts) + string(char)
			sb.WriteString(txt)
			char = newChar
			counts = 1
		} else {
			counts++
		}
	}

	// write last character
	txt := strconv.Itoa(counts) + string(char)
	sb.WriteString(txt)

	fmt.Println(sb.String())
}
