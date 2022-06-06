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
	input = strings.TrimSuffix(input, "\n")

	array := convertStringToArrayOfIntSlice(size, input)
	count := countStrictlyDesceasing(size, array)
	fmt.Print(count)

}

func convertStringToArrayOfIntSlice(size int, input string) []int {
	out := make([]int, 0, size)

	inputArray := strings.Split(input, " ")

	for _, char := range inputArray {
		num, _ := strconv.Atoi(char)
		out = append(out, num)
	}

	return out
}

// the idea is that a sorted subarray of length 'l' adds l*(l-1)/2 to result
// means all the combinations between them will be counted
func countStrictlyDesceasing(size int, array []int) int {
	count := 0

	length := 1
	for index := 0; index < size-1; index++ {
		if array[index+1] < array[index] {
			length++
		} else {
			count += (((length - 1) * length) / 2)
			length = 1
		}
	}

	if length > 1 {
		count += (((length - 1) * length) / 2)
	}

	return count
}
