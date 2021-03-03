package main

func gcd(num1 int, num2 int) int {
	for num2 != 0 {
		reminder := num1 % num2
		num1 = num2
		num2 = reminder
	}

	return num1
}
