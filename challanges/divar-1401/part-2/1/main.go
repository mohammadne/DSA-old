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
	var counts int
	fmt.Scan(&counts)

	sessions := make([]session, counts)

	reader := bufio.NewReader(os.Stdin)
	for index := 0; index < counts; index++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		splits := strings.Split(line, " ")
		start, _ := strconv.Atoi(splits[1])
		end, _ := strconv.Atoi(splits[2])
		sessions[index] = session{start: start, end: end}
	}

	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	splits := strings.Split(line, " ")
	duration, _ := strconv.Atoi(splits[1])

	if duration < sessions[0].start {
		fmt.Println(0, duration)
		return
	}

	for index := 0; index < counts; index++ {
		if diff := sessions[index+1].start - sessions[index].end; diff > duration {
			fmt.Println(sessions[index].end+1, sessions[index].end+1+duration)
			return
		}
	}

	fmt.Println(sessions[counts-1].end+1, sessions[counts-1].end+1+duration)
}
