package main

import "fmt"

func main() {
	a := []int{212, 12, 3001, 14, 501, 7800, 9932, 33, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(InsertionSort(a))
}

//Difference with bubble sort, Here at any iteration of outerloop,
//the array to the left of the element will be sorted
func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i+1; j++ {
			//compare element present at index i with every element present
			//  left of it place it in right place so that array on the
			//left remains   sorted
			if numbers[j] > numbers[i] {
				intermediate := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = intermediate
			}
		}
		fmt.Println(numbers)
	}
	return numbers
}
