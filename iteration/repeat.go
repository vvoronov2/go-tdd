package iteration

import "strings"

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}

func StringRepeat(word string, count int) string {
	return strings.Repeat(word, count)
}

func StringEqual(first string, second string) bool {
	return strings.EqualFold(first, second)
}
