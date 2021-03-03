package main

func maxSubSum4(vector []int) int {
	maxSum := 0

	currentSum := 0
	for index := 0; index < len(vector); index++ {
		currentSum += vector[index]
		currentSum = max(currentSum, 0)

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
