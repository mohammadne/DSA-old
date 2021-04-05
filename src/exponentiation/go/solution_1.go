package exponentiation

func exponentiation1(num int, power int) int {
	if power == 0 {
		return 1
	} else if power == 1 {
		return num
	}

	return exponentiation1(num, power-1) * num
}
