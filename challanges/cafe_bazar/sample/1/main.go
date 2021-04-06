package main

import (
	"fmt"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)

	var sb strings.Builder
	for index := 0; index < a; index++ {
		sb.WriteString("man khoshghlab hastam\n")
	}

	o := strings.Trim(sb.String(), "\n")
	fmt.Print(o)
}
