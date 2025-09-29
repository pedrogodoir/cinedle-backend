package utils

import "strings"

//capitalize the first letter of each word in a string
func ToTitle(s string) string {
	words := strings.Fields(strings.ToLower(s))
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, " ")
}
