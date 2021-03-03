package main

// maxSubSum1 is CUBIT maximum contiguous subsequence sum algorithm.
func maxSubSum1(vector []int) int {
	length := len(vector)
	maxSum := 0

	for firstIndex := 0; firstIndex < length; firstIndex++ {
		for secondIndex := firstIndex; secondIndex < length; secondIndex++ {
			innerSum := 0

			for index := firstIndex; index < secondIndex; index++ {
				innerSum += vector[index]
			}

			if innerSum > maxSum {
				maxSum = innerSum
			}
		}
	}

	return maxSum
}
