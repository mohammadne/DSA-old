package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	length, input := inputReciever()
	data := arrangement(length, input)

	output := 0

	for index := 0; index < len(data)-1; index++ {
		min := MinOf(data[index]...)
		max := MaxOf(data[index]...)

		minus := (max - min)

		nextMin := MinOf(data[index+1]...)

		if value := nextMin - 1; value > 0 {
			minus += value
		}

		if minus > output {
			output = minus
		}

	}

	fmt.Println(output)

}

func arrangement(length int, data [][2]int) map[int][]int {
	output := map[int][]int{}

	for zel := 0; zel < length; zel++ {
		spotsInRow := make([]int, 0, 10)

		for i := 0; i < len(data); i++ {
			if data[i][0]-1 == zel {
				spotsInRow = append(spotsInRow, data[i][1])
			}
		}

		output[zel] = spotsInRow
	}

	return output

}

func inputReciever() (int, [][2]int) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	inputSlice := strings.Split(scanner.Text(), " ")
	length, _ := strconv.Atoi(inputSlice[0])
	spotsNumber, _ := strconv.Atoi(inputSlice[1])

	lines := make([][2]int, 0, 50)

	for lineIndex := 0; lineIndex < spotsNumber; lineIndex++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		lines = append(lines, [2]int{x, y})
	}

	return length, lines
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
