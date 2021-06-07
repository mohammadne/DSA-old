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

var worthOfAlphabets = map[byte]byte{
	byte('a'): 1,
	byte('b'): 2,
	byte('c'): 3,
	byte('d'): 4,
	byte('e'): 5,
}

// Sentence has union type charactristics
type Sentence struct {
	verb         string
	subSentences []Sentence

	noun string
}

func (s Sentence) process() int {
	if noun := s.noun; len(noun) != 0 {
		return processNoun(noun)
	}

	values := make([]int, 0, len(s.subSentences))

	for index := 0; index < len(s.subSentences); index++ {
		values = append(values, s.subSentences[index].process())
	}

	result := values[0]
	for index := 1; index < len(values); index++ {
		if s.verb == addition {
			result += values[index]
		} else if s.verb == subtraction {
			result -= values[index]
		} else if s.verb == multiplication {
			result *= values[index]
		} else if s.verb == division {
			result /= values[index]
		}
	}

	return result
}

func main() {
	input := "abcd abcd aabbc ab a c ccd dede cccd cd  "
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
		} else {
			sentence := processTerms(terms[verbsIndex[index]:verbsIndex[index+1]])
			subSentences = append(subSentences, sentence)
		}
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

func processNoun(str string) int {
	values := make([]int, 0, len(str))

	repetition := 0
	current := byte(0)

	for index := 0; index < len(str); index++ {
		if current == 0 {
			current = str[index]
			repetition++
		} else if current != str[index] {
			worth := worthOfAlphabets[current]
			values = append(values, int(repetition)*int(worth))

			current = byte(str[index])
			repetition = 1
		} else {
			repetition++
		}
	}

	worth := worthOfAlphabets[current]
	values = append(values, int(repetition)*int(worth))

	for index := 0; index < len(values); index++ {
		newVal := values[index] % 5
		values[index] = newVal * newVal
	}

	return sumIntSlice(values)
}

func sumIntSlice(slice []int) int {
	sum := 0
	for index := 0; index < len(slice); index++ {
		sum += slice[index]
	}

	return sum
}
