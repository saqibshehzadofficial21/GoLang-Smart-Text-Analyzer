package analyzer

import "fmt"

type Result struct {
	TotalChars       int
	TotalWords       int
	TotalLetters     int
	Vowels           int
	Consonants       int
	Uppercase        int
	Lowercase        int
	CharsWithSpace   int
	CharsWithoutSpace int
	Punctuations     int
	Integers         int
	Sentences        int
	Lines            int
	Spaces           int
}

func Analyze(text string) Result {
	var r Result

	inWord := false
	inSentence := false

	for _, char := range text {

		r.CharsWithSpace++
		r.TotalChars++

		if char == ' ' || char == '\t' {
			r.Spaces++
		}

		if char != ' ' && char != '\n' && char != '\t' {
			r.CharsWithoutSpace++
		}

		// Word Count
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if !inWord {
				r.TotalWords++
				inWord = true
			}
		} else {
			inWord = false
		}

		// Letters & Case
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			r.TotalLetters++

			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			   char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
				r.Vowels++
			} else {
				r.Consonants++
			}

			if char >= 'A' && char <= 'Z' {
				r.Uppercase++
			} else {
				r.Lowercase++
			}
		}

		// Punctuation
		if isPunctuation(char) {
			r.Punctuations++
		}

		// Numbers
		if char >= '0' && char <= '9' {
			r.Integers++
		}

		// Sentences
		if char == '.' || char == '!' || char == '?' {
			if inSentence {
				r.Sentences++
				inSentence = false
			}
		} else if char != ' ' && char != '\n' && char != '\t' {
			inSentence = true
		}

		// Lines
		if char == '\n' {
			r.Lines++
		}
	}

	if r.TotalChars > 0 {
		r.Lines++
	}

	return r
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
}