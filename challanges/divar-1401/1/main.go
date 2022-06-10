package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var size int
	fmt.Scan(&size)

	sum := 0

	reader := bufio.NewReader(os.Stdin)

	for index := 0; index < size; index++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")

		splits := strings.Split(line, " ")

		if main, _ := strconv.Atoi(splits[index]); main%3 == 1 {
			sum += main
		}

		if notMain, _ := strconv.Atoi(splits[size-index-1]); notMain%3 == 1 {
			sum += notMain
		}
	}

	fmt.Println(sum)
}
