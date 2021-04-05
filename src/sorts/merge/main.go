package main

import (
	"fmt"
)

func main() {
	a := []int{2, 1, 3, 4, 50, 78, 32, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(MergeSort(a))
}

func merge(a, b []int) []int {
	size, i, j := len(a)+len(b), 0, 0
	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			//all the elements of a already been judged
			//assuming a and b both are sorted, this happens because
			//some cases will have not equal a and b, so it might
			// be a possibility that one array got finished earlier.
			result[k] = b[j]
			j += 1
		case j == len(b):
			//alll the elements of a already been judged
			//assuming a nd b both are sorted
			result[k] = a[i]
			i += 1
		case a[i] > b[j]:
			result[k] = b[j]
			//increment the index of b array because its present index
			//is already appended to the result array
			j += 1
		case a[i] < b[j]:
			//increment the index of a array because its present index
			//element is already appended to the result array
			result[k] = a[i]
			i += 1
		case a[i] == b[j]:
			//in case of equality
			result[k] = b[j]
			j += 1
		}
	}
	return result
}

func MergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	middle := int(len(numbers) / 2)
	return merge(MergeSort(numbers[middle:]), MergeSort(numbers[:middle]))
}
