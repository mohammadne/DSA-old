package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	start int
	end   int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	p := pair{start: 0, end: len(input) - 1}
	answer := recursive(&p, input)
	fmt.Print(answer % 2021)
}

func recursive(p *pair, input string) int {
	// base case
	if p.end-p.start <= 1 {
		return 1
	}

	pairs := make([]pair, 0, 5)

	balance := 0
	start := p.start

	// we will count parallel pairs in the string
	length := 1 + p.end
	for index := p.start + 0; index < length; index++ {
		if input[index] == '(' {
			balance++
		} else if input[index] == ')' {
			balance--
		}

		if balance == 0 {
			p := pair{start: start, end: index}
			pairs = append(pairs, p)
			if length-1 != index {
				start = index + 1
			}
		}
	}

	// if it's an enclosing paranthese, then double it by 2
	if len(pairs) == 1 {
		return 2 * recursive(&pair{start: p.start + 1, end: p.end - 1}, input)
	}

	// return the summation of parallel paranthese as the result
	summation := 0
	for _, pair := range pairs {
		summation += recursive(&pair, input)
	}
	return summation
}
