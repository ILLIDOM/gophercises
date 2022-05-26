package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(wordsInCamelCaseString("halloTestCase"))
}

func wordsInCamelCaseString(word string) int {
	number_of_words := 1

	for _, char := range word {
		str := string(char)
		if strings.ToUpper(str) == str {
			number_of_words++
		}
	}

	return number_of_words
}
