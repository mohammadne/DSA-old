package main

import (
	"fmt"
	"strings"
)

const (
	// operators
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

type Noun string

func (n Noun) process() int {
	values := make([]int, 0, len(n))

	repetition := 0
	current := byte(0)

	for index := 0; index < len(n); index++ {
		if current == 0 {
			current = n[index]
			repetition++
		} else if current != n[index] {
			worth := worthOfAlphabets[current]
			values = append(values, int(repetition)*int(worth))

			current = byte(n[index])
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

type Sentence struct {
	verb         string
	subSentences []Sentence

	noun Noun
}

func (s Sentence) process() float64 {
	if noun := s.noun; len(noun) != 0 {
		return float64(noun.process())
	}

	values := make([]float64, 0, len(s.subSentences))

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
	input := "abcd bcde ab ac abab a b"
	splitedInput := strings.Split(input, " ")
	sentence := buildGeneralizedList(splitedInput)

	value := sentence.process()
	fmt.Println(value)
}

func buildGeneralizedList(terms []string) Sentence {
	length := len(terms)

	if length == 1 {
		return Sentence{
			noun: Noun(terms[0]),
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

	for index := 1; index < lastVerbIndex; index++ {
		sentence := Sentence{noun: Noun(terms[index])}
		subSentences = append(subSentences, sentence)
	}

	for index := 0; index < verbsLength; index++ {
		if verbsLength-index <= 1 {
			sentence := buildGeneralizedList(terms[verbsIndex[index]:])
			subSentences = append(subSentences, sentence)
		} else {
			sentence := buildGeneralizedList(terms[verbsIndex[index]:verbsIndex[index+1]])
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

func sumIntSlice(slice []int) int {
	sum := 0
	for index := 0; index < len(slice); index++ {
		sum += slice[index]
	}

	return sum
}
