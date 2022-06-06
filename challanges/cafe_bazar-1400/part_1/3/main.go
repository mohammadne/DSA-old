package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input1, _ := reader.ReadString('\n')
	trimmedInput1 := strings.Trim(input1, "\n")
	createdAssignee1 := assignee(trimmedInput1)

	input2, _ := reader.ReadString('\n')
	trimmedInput2 := strings.Trim(input2, "\n")
	createdAssignee2 := assignee(trimmedInput2)

	if createdAssignee1 == createdAssignee2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}

}

func assignee(input string) string {
	assign := 0
	assignMap := map[rune]int{}

	var sb strings.Builder
	for _, runeChar := range input {
		if val, ok := assignMap[runeChar]; ok {
			txt := fmt.Sprintf("%d", val)
			sb.WriteString(txt)
		} else {
			assign++
			assignMap[runeChar] = assign

			txt := fmt.Sprintf("%d", assign)
			sb.WriteString(txt)
		}
	}

	return sb.String()
}
