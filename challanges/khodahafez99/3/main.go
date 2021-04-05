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
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	splitedLine := append(strings.Split(input[0:len(input)-1], " "))

	// mobiles, _ := strconv.Atoi(splitedLine[0])
	migire, _ := strconv.Atoi(string(splitedLine[1]))
	mide, _ := strconv.Atoi(string(splitedLine[2]))

	input, _ = reader.ReadString('\n')
	splitedLine = strings.Split(input, " ")

	data := strings.Split(input, " ")

	minIndex := 0
	min, _ := strconv.Atoi(data[minIndex])
	for index, v := range data {
		newVal, _ := strconv.Atoi(v)
		if newVal < min {
			min = newVal
			minIndex = index
		}
	}

	reduceData := RemoveIndex(data, minIndex)

	chargePartsNeed := 0
	for _, value := range reduceData {
		intValue, _ := strconv.Atoi(value)
		divisions := math.Ceil(float64(100%intValue) / float64(mide))
		chargePartsNeed += int(divisions)
	}

	if divided := min / migire; divided >= chargePartsNeed {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
