package maxsubsum

// maxSubSum3 is Driver for divide-and-conquer maximum contiguous
// subsequence sum algorithm.
func maxSubSum3(vector []int) int {
	return maxSumRec(vector, 0, len(vector)-1)
}

// maxSumRec is Recursive maximum contiguous subsequence sum algorithm.
// Finds maximum sum in subarray spanning a[left..right].
// Does not attempt to maintain actual best sequence.
func maxSumRec(vector []int, left int, right int) int {
	// base case
	if left == right {
		return max(vector[left], 0)
	}

	center := (left + right) / 2
	maxLeftSum := maxSumRec(vector, left, center)
	maxRightSum := maxSumRec(vector, center+1, right)

	maxLeftBorderSum, leftBorderSum := 0, 0
	for index := center; index < left; index-- {
		leftBorderSum += vector[index]
		if leftBorderSum > maxLeftBorderSum {
			maxLeftBorderSum = leftBorderSum
		}
	}

	maxRightBorderSum, rightBorderSum := 0, 0
	for index := center; index < right; index++ {
		rightBorderSum += vector[index]
		if rightBorderSum > maxRightBorderSum {
			maxRightBorderSum = rightBorderSum
		}
	}

	maxCrossingSum := maxLeftBorderSum + maxRightBorderSum

	return max(maxLeftSum, maxRightSum, maxCrossingSum)
}
