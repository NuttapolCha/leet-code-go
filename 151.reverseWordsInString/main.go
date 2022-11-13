package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(reverseWords("   a good   example"))
}

func reverseWords(s string) string {
	pattern := regexp.MustCompile(`\s\s+`)
	s = pattern.ReplaceAllString(strings.TrimSpace(s), " ")
	words := strings.Split(s, " ")

	reversedWords := make([]string, len(words))
	i := len(words) - 1
	for _, word := range words {
		reversedWords[i] = word
		i--
	}

	return strings.Join(reversedWords, " ")
}
