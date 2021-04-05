package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	nums := aizuArray(input)

	if len(nums) != 2 {
		os.Exit(1)
	}

	if nums[0] > nums[1] {
		lNums := nums[0] - nums[1]

		for index := 0; index < lNums; index++ {
			fmt.Printf("L ")
		}

	} else if nums[0] < nums[1] {
		rNums := nums[1] - nums[0]

		for index := 0; index < rNums; index++ {
			fmt.Printf("R ")
		}
	} else {
		fmt.Printf("Saal Noo Mobarak!")

	}
}

func aizuArray(input string) []int {
	inputStringSlice := strings.Split(input, " ")
	outputIntSlice := make([]int, len(inputStringSlice))

	for i, v := range inputStringSlice {
		temp, _ := strconv.Atoi(v)
		outputIntSlice[i] = temp
	}

	return outputIntSlice
}
