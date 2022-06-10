package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element struct {
	row    int
	column int
}

func main() {
	var size int
	fmt.Scan(&size)

	matrix := make([][]bool, size)

	reader := bufio.NewReader(os.Stdin)
	for row := 0; row < size; row++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		line = strings.ReplaceAll(line, " ", "")

		matrix[row] = make([]bool, size)
		for column := 0; column < size; column++ {
			if line[column] == '1' {
				matrix[row][column] = true
			}
		}
	}

	maximum := 0

	for row := 0; row < size; row++ {
		for column := 0; column < size; column++ {
			if matrix[row][column] {
				value := findRecursive(matrix, Element{row: row, column: column}, size)
				if value > maximum {
					maximum = value
				}
			}
		}
	}

	fmt.Println(maximum)
}

func findRecursive(matrix [][]bool, e Element, size int) int {
	if !matrix[e.row][e.column] {
		return 0
	}

	matrix[e.row][e.column] = false
	value := 1

	adjs := getAdjacents(matrix, e, size)
	if len(adjs) == 0 {
		return value
	}

	for index := 0; index < len(adjs); index++ {
		value += findRecursive(matrix, adjs[index], size)
	}

	return value
}

func getAdjacents(matrix [][]bool, e Element, size int) []Element {
	res := make([]Element, 0)

	if up := e.row - 1; up >= 0 && matrix[up][e.column] {
		res = append(res, Element{row: up, column: e.column})
	}

	if left := e.column - 1; left >= 0 && matrix[e.row][left] {
		res = append(res, Element{row: e.row, column: left})
	}

	if down := e.row + 1; down < size && matrix[down][e.column] {
		res = append(res, Element{row: down, column: e.column})
	}

	if right := e.column + 1; right < size && matrix[e.row][right] {
		res = append(res, Element{row: e.row, column: right})
	}

	return res
}
