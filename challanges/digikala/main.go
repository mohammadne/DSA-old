package main

import "fmt"

func main() {
	input := "aa"

	length := len(input)
	diverge := 0

	center := length / 2
	for index := 0; index <= center-diverge/2; index++ {
		if input[index-diverge] != input[length-index-1] {
			diverge++
		}

		if diverge > 1 {
			break
		}
	}

	if diverge > 1 {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}
}
