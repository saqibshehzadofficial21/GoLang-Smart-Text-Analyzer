package analyzer

func isLetter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isUpper(char rune) bool {
	return char >= 'A' && char <= 'Z'
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
		char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'
}

func isPunctuation(char rune) bool {
	punct := ".,!?-'\":;()[]{}"
	for _, p := range punct {
		if char == p {
			return true
		}
	}
	return false
}

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func getWordCount(char rune, inWord *bool) int {
	if isLetter(char) || isDigit(char) {
		if !*inWord {
			*inWord = true
			return 1
		}
		return 0
	}
	*inWord = false
	return 0
}

func getSentenceCount(char rune, inSentence *bool) int {
	if char == '.' || char == '!' || char == '?' {
		if *inSentence {
			*inSentence = false
			return 1
		}
	} else if !isSpace(char) {
		*inSentence = true
	}
	return 0
}