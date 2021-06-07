package main

import (
	"fmt"
	"strings"
)

const (
	addition       = "abcd"
	subtraction    = "bcde"
	multiplication = "dede"
	division       = "abab"
)

// Sentence has union type charactristics
type Sentence struct {
	verb         string
	subSentences []Sentence

	noun string
}

func (s Sentence) process() int {
	return 0
}

func main() {
	input := ""
	processedInput := processString(input)
	sentence := processTerms(processedInput)

	value := sentence.process()
	fmt.Println(value)
}

func processString(input string) []string {
	return strings.Split(input, " ")
}

func processTerms(terms []string) Sentence {
	length := len(terms)

	if length == 1 {
		return Sentence{
			noun: terms[0],
		}
	}

	verb := terms[0]
	subSentences := make([]Sentence, 0, length/2)

	verbsIndex := make([]int, 0, length/2)
	for index := 1; index < length; index++ {
		if isVerb(terms[index]) {
			verbsIndex = append(verbsIndex, index)
		}
	}

	verbsLength := len(verbsIndex)
	lastVerbIndex := length
	if verbsLength != 0 {
		lastVerbIndex = verbsIndex[0]
	}

	for index := 0; index < lastVerbIndex; index++ {
		sentence := Sentence{noun: terms[index]}
		subSentences = append(subSentences, sentence)
	}

	for index := 0; index < verbsLength; index++ {
		if verbsLength-index <= 1 {
			sentence := processTerms(terms[verbsIndex[index]:])
			subSentences = append(subSentences, sentence)
		}

		sentence := processTerms(terms[verbsIndex[index]:verbsIndex[index+1]])
		subSentences = append(subSentences, sentence)
	}

	return Sentence{
		subSentences: subSentences,
		verb:         verb,
	}

}

func isVerb(input string) bool {
	return input == addition ||
		input == subtraction ||
		input == multiplication ||
		input == division
}
