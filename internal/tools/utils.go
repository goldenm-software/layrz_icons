package tools

import (
	"strings"
	"unicode"
)

// splitWords splits a camelCase/PascalCase/snake_case/kebab-case string into words.
// Handles acronyms like "HTMLParser" → ["HTML", "Parser"].
func splitWords(name string) []string {
	runes := []rune(name)
	var words []string
	var current []rune

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Separators: split and skip
		if r == '_' || r == '-' || r == ' ' {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			continue
		}

		if len(current) == 0 {
			current = append(current, r)
			continue
		}

		prevIsUpper := unicode.IsUpper(current[len(current)-1])
		prevIsLower := unicode.IsLower(current[len(current)-1])
		prevIsDigit := unicode.IsDigit(current[len(current)-1])
		curIsUpper := unicode.IsUpper(r)
		curIsLower := unicode.IsLower(r)
		curIsDigit := unicode.IsDigit(r)

		switch {
		// lowercase/digit → uppercase: start new word (e.g., "codePoint" → "code", "Point")
		case (prevIsLower || prevIsDigit) && curIsUpper:
			words = append(words, string(current))
			current = []rune{r}

		// uppercase → uppercase + lowercase: acronym boundary (e.g., "HTMLParser" at 'P', split before last upper)
		case prevIsUpper && curIsLower && len(current) > 1:
			words = append(words, string(current[:len(current)-1]))
			current = current[len(current)-1:]
			current = append(current, r)

		// letter → digit or digit → letter: start new word
		case (prevIsDigit && !curIsDigit) || (!prevIsDigit && curIsDigit && !prevIsUpper && !prevIsLower):
			words = append(words, string(current))
			current = []rune{r}

		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		words = append(words, string(current))
	}

	return words
}

func ToKebabCase(name string) string {
	words := splitWords(name)
	for i, w := range words {
		words[i] = strings.ToLower(w)
	}
	return strings.Join(words, "-")
}

func ToCamelCase(name string, capitalize bool) string {
	words := splitWords(name)
	for i, w := range words {
		if len(w) > 0 {
			runes := []rune(strings.ToLower(w))
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	result := strings.Join(words, "")
	if len(result) == 0 {
		return result
	}
	runes := []rune(result)
	if capitalize {
		runes[0] = unicode.ToUpper(runes[0])
	} else {
		runes[0] = unicode.ToLower(runes[0])
	}
	return string(runes)
}
