package maxsubsum

func max(vars ...int) int {
	max := vars[0]

	for _, index := range vars {
		if index > max {
			max = index
		}
	}

	return max
}
