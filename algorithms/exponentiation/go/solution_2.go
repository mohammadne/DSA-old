package exponentiation

func exponentiation2(num int, power int) int {
	if power == 0 {
		return 1
	} else if power == 1 {
		return num
	}

	if reminder := power % 2; reminder == 0 {
		return exponentiation2(num*num, power/2)
	}

	return exponentiation2(num*num, power/2) * num
}
