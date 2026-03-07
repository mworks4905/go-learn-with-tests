package iteration

import "strings"

func Repeat(character string, iterations int) string {
	var repeated strings.Builder
	for range iterations {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func Reverse(word string) string {
	var reversed string

	for _, letter := range word {
		reversed = string(letter) + reversed
	}
	return reversed
}