package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	part := a / 2

	v := part + 1
	h := part + 1

	if isOdd(a) {
		v++
	}

	fmt.Printf("%d", h*v)
}

func isOdd(num int) bool {
	return num%2 != 0
}
