package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) buildTree(value int, difference int, height int) {
	n.value = value

	if height > 1 {
		n.left.buildTree(value-difference, difference, height-1)
		n.right.buildTree(value+difference, difference, height-1)
	}
}

func (n *Node) isValid() bool {
	return n.value >= 0 && n.value <= 9
}

func (n *Node) validNumbers(height int) []int {
	if height == 1 {
		if n.isValid() {
			return []int{n.value}
		}

		return []int{}
	}

	fullSlice := make([]int, 0)

	if n.left.isValid() {
		lNums := n.left.validNumbers(height - 1)
		lNumsAppend := appendNumToStartOfSliceNumbers(n.left.value, lNums)
		fullSlice = append(fullSlice, lNumsAppend...)
	}

	if n.right.isValid() {
		rNums := n.right.validNumbers(height - 1)
		rNumsAppend := appendNumToStartOfSliceNumbers(n.right.value, rNums)
		fullSlice = append(fullSlice, rNumsAppend...)
	}

	return fullSlice
}

func main() {
	var length, difference int
	fmt.Scan(&length, &difference)

	result := make([]int, 0, 10)

	for index := 1; index < 10; index++ {
		tree := Node{}
		tree.buildTree(index, difference, length)
		result = append(result, tree.validNumbers(length)...)
	}

	var sb strings.Builder
	for index := 0; index < len(result); index++ {
		value := strconv.Itoa(result[index])
		sb.WriteString(value)
	}

	fmt.Println(sb.String())
}

func appendNumToStartOfSliceNumbers(num int, slice []int) []int {
	for index := 0; index < len(slice); index++ {
		value := slice[index]
		digits := CountDigits(value)
		slice[index] = (digits+1)*num + value
	}

	return slice
}

func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}
