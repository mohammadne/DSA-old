package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var specialNum, maxUsage int
	fmt.Scan(&specialNum, &maxUsage)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	trimmedInput := strings.Trim(input, "\n")

	specialArray := make([]int, 0, specialNum)
	for _, char := range strings.Split(trimmedInput, " ") {
		num, _ := strconv.Atoi(char)
		specialArray = append(specialArray, num)
	}

	sortedSpecialArray := RadixSort(specialArray)

	firstArrival := sortedSpecialArray[0]
	diffArray := createDiffrenceArray(specialNum, sortedSpecialArray)
	sorteddiffArray := RadixSort(diffArray)
	res := sorteddiffArray[:specialNum-1-maxUsage]

	fmt.Print(firstArrival + sumArray(res))
}

func RadixSort(intArr []int) []int {
	tmp := make([]int, len(intArr))
	copy(tmp, intArr)
	places := BigNumPlaceCount(tmp)

	for index := range make([]int, places) {

		place := int(math.Pow(float64(10), float64(index)))

		count := [10]int{}

		Loop(place, intArr, &count)
		intArr = AuxArray(place, intArr, &count)

	}

	return intArr
}

//AuxArray generate a new array for a place value divisor
// The new array will be sorted according to the place value in
//numbers
func AuxArray(divisor int, intArr []int, count *[10]int) []int {
	//Start from the end
	aux := make([]int, len(intArr))
	for i := len(intArr) - 1; i >= 0; i-- {
		//find the target significant digit, if divisor is 10,
		//find the 10th place digit in the number.
		k := (intArr[i] / divisor) % 10
		//find the value corrsponding to this index in the count array
		index := (*count)[k]
		//Now in aux array, put this number at the index
		aux[index] = intArr[i]
		//Decrement the count array at the kth index.
		(*count)[k]--
		//fmt.Printf("Count %v -- Aux %v --- IntArr %v\n", *count, intArr, aux)
	}
	return aux
}

func createDiffrenceArray(size int, sortedSpecialArray []int) []int {
	arr := make([]int, 0, size-1)

	for index := 0; index < size-1; index++ {
		distance := sortedSpecialArray[index+1] - sortedSpecialArray[index]
		arr = append(arr, distance)
	}

	return arr
}

func BigNumPlaceCount(arr []int) int {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			rplc := arr[i-1]
			arr[i-1] = arr[i]
			arr[i] = rplc
		}
	}
	biggest := arr[len(arr)-1]
	return len(strconv.Itoa(biggest))
}

func Loop(divisor int, intArr []int, count *[10]int) {
	for _, value := range intArr {
		rem := (value / divisor) % 10
		(*count)[rem] += 1
	}
	(*count)[0] = (*count)[0] - 1

	for i := 1; i < len(*count); i++ {
		(*count)[i] = (*count)[i] + (*count)[i-1]
	}
}

func sumArray(input []int) int {
	sum := 0

	for _, num := range input {
		sum += num
	}

	return sum
}
