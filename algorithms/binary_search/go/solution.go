package main

func binarySearch(inputs []int, x int) int {
	low, high := 0, len(inputs)

	for low <= high {
		middle := (low + high) / 2

		if inputs[middle] < x {
			low = middle + 1
		} else if inputs[middle] > x {
			high = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
