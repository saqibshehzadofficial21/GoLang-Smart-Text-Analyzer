package analyzer

import "time"

func Analyze(text string) Result {

	start := time.Now()

	var r Result
	inWord := false
	inSentence := false

	for _, char := range text {

		r.TotalChars++
		r.CharsWithSpace++

		if char == ' ' || char == '\t' {
			r.Spaces++
		}

		if char != ' ' && char != '\n' && char != '\t' {
			r.CharsWithoutSpace++
		}

		r.TotalWords += getWordCount(char, &inWord)

		if isLetter(char) {
			r.TotalLetters++

			if isVowel(char) {
				r.Vowels++
			} else {
				r.Consonants++
			}

			if isUpper(char) {
				r.Uppercase++
			} else {
				r.Lowercase++
			}
		}

		if isPunctuation(char) {
			r.Punctuations++
		}

		if isDigit(char) {
			r.Integers++
		}

		r.Sentences += getSentenceCount(char, &inSentence)

		if char == '\n' {
			r.Lines++
		}
	}

	if r.TotalChars > 0 {
		r.Lines++
	}

	r.ExecutionTime = time.Since(start)

	return r
}