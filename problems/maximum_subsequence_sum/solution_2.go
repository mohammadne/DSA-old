package main

// maxSubSum2 is Quadratic maximum contiguous subsequence sum algorithm.
func maxSubSum2(vector []int) int {
	length := len(vector)
	maxSum := 0

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		rowSum := 0

		for secondIndex := firstIndex; secondIndex < length; secondIndex++ {
			rowSum += vector[secondIndex]

			if rowSum > maxSum {
				maxSum = rowSum
			}
		}
	}

	return maxSum
}
