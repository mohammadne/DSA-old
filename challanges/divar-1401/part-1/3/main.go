package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var line string
	fmt.Scan(&line)

	var sb strings.Builder
	for i := 0; i < len(line); i++ {
		b := line[i]
		if b == '0' || b == '1' || b == '2' || b == '3' || b == '4' ||
			b == '5' || b == '6' || b == '7' || b == '8' || b == '9' {
			sb.WriteByte(b)
		}
	}
	numberedLine := sb.String()

	for index := 1; index < len(numberedLine); index++ {
		number, _ := strconv.Atoi(numberedLine[index-1 : index+1])
		if isPrime(number) {
			fmt.Println(number)
		}
	}
}

func isPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	}

	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}
