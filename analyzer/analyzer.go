package analyzer

import (
	"fmt"
	"time"
)

type Result struct {
	TotalChars        int
	TotalWords        int
	TotalLetters      int
	Vowels            int
	Consonants        int
	Uppercase         int
	Lowercase         int
	CharsWithSpace    int
	CharsWithoutSpace int
	Punctuations      int
	Integers          int
	Sentences         int
	Lines             int
	Spaces            int
	ExecutionTime     time.Duration
}

func Analyze(text string) Result {

	startTime := time.Now()

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

		// Word Count
		r.TotalWords += getWordCount(char, &inWord)

		// Letters, Vowels, Consonants, Case
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

		// Other Counts
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

	r.ExecutionTime = time.Since(startTime) // Time end 

	return r
}

// Simple Helper Functions

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

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}


func PrintResult(r Result) {
	fmt.Printf("\n \n TEXT ANYLIZER... \n\n")
	fmt.Printf("01 -Total characters: %d\n", r.TotalChars)
	fmt.Printf("02 -Total characters: %d\n", r.TotalChars)
	fmt.Printf("03 -Total words: %d\n", r.TotalWords)
	fmt.Printf("04 -Total letters: %d\n", r.TotalLetters)
	fmt.Printf("05 -VOWELS: %d\n", r.Vowels)
	fmt.Printf("06 -CONSONANTS: %d\n", r.Consonants)
	fmt.Printf("07 -Upper-Case Letters: %d\n", r.Uppercase)
	fmt.Printf("08 -Lower-Case Letters: %d\n", r.Lowercase)
	fmt.Printf("09 -Total Chars Without-Space : %d\n", r.CharsWithoutSpace)
	fmt.Printf("10 -Total Chars WithSpace : %d\n", r.CharsWithSpace)
	fmt.Printf("11 -Punctuations : %d\n", r.Punctuations)
	fmt.Printf("12 -Integers (Numbers) : %d\n", r.Integers)
	fmt.Printf("13 -Sentences : %d\n", r.Sentences)
	fmt.Printf("14 -Total Lines : %d\n", r.Lines)
	fmt.Printf("15 -Total Spaces : %d\n", r.Spaces)
	fmt.Printf("16 -Execution Time : %v\n", r.ExecutionTime)
}
