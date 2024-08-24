package controllers

import "strings"

func StructureString(str string) string {
	var result string

	for _, char := range str {
		if char == '_' {
			result += " "
		} else if char == '-' {
			result += ", "
		} else {
			result += string(char)
		}
	}
	return result
}

func FormatString(str string) string {
	words := strings.Fields(StructureString(str))

	for i, word := range words {
		switch strings.ToLower(word) {
		case "usa", "uk":
			words[i] = strings.ToUpper(word)
		default:
			words[i] = strings.ToUpper(string(word[0])) + string(word[1:])
		}
	}
	return strings.Join(words, " ")
}
