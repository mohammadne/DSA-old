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

	maxDiff := data[1] - data[0]

	minNumBeforeIndex := data[0]
	for index := 1; index < size; index++ {
		num := data[index-1]
		if num < minNumBeforeIndex {
			minNumBeforeIndex = num
		}

		dif := data[index] - minNumBeforeIndex
		if dif > maxDiff {
			maxDiff = dif
		}
	}

	fmt.Print(maxDiff)

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
