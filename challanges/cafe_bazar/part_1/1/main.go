package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size int
	fmt.Scan(&size)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	data := arranger(size, input)
	sum := sumArray(data)

	for index := 0; index < size; index++ {
		num := sum - data[index]
		fmt.Print(num)

		if index != size-1 {
			fmt.Print(" ")
		}
	}
}

func arranger(size int, input string) []int {
	out := make([]int, 0, size)

	inputArray := strings.Split(input, " ")

	for _, char := range inputArray {
		num, _ := strconv.Atoi(char)
		out = append(out, num)
	}

	return out
}

func sumArray(input []int) int {
	sum := 0

	for _, num := range input {
		sum += num
	}

	return sum
}
