package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type session struct {
	start int
	end   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	maxIndex := 0

	// ====================================================================> 1

	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	counts1, _ := strconv.Atoi(line)

	sessions1 := make(map[int]session, counts1)

	for index := 0; index < counts1; index++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		splits := strings.Split(line, " ")
		index, _ := strconv.Atoi(splits[0])
		start, _ := strconv.Atoi(splits[1])
		end, _ := strconv.Atoi(splits[2])
		sessions1[index] = session{start: start, end: end}

		if index > maxIndex {
			maxIndex = index
		}
	}

	// ====================================================================> 2

	line, _ = reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	counts2, _ := strconv.Atoi(line)

	sessions2 := make(map[int]session, counts2)

	for index := 0; index < counts2; index++ {
		line, _ = reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		splits := strings.Split(line, " ")
		index, _ := strconv.Atoi(splits[0])
		start, _ := strconv.Atoi(splits[1])
		end, _ := strconv.Atoi(splits[2])
		sessions2[index] = session{start: start, end: end}

		if index > maxIndex {
			maxIndex = index
		}
	}

	// ====================================================================> SUMMATION

	sum := 0

	for index := 0; index < maxIndex; index++ {
		session1, found1 := sessions1[index]
		session2, found2 := sessions2[index]

		if !found1 || !found2 {
			continue
		}

		if session1.start < session2.end && session2.start < session1.end {
			sum += min(session1.end, session2.end) - max(session1.start, session2.start)
		}
	}

	fmt.Println(sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
