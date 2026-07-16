
package main

import (
	"fmt"
	"os"
	"text-analyzer/analyzer"
)

func main() {
	data, err := os.ReadFile("02-PARTICE.txt")
	if err != nil {
		fmt.Println("SHOWS ERROS : ", err)
		return
	}

	text := string(data)

	result := analyzer.Analyze(text)
	analyzer.PrintResult(result)
}








































// // package main

// // // Buffered (The Store-and-Forward Method)
// // // Unbuffered (The Direct Hand-off)
// // //"bufio" // buffered input/output (io.Reader and io.Writer) ==> to read it line-by-line or word-by-word

// // import (
// // 	"fmt"
// // 	"unicode"

// // 	"os"
// // )

// // func main() {

// // 	fmt.Printf("\n \n TEXT ANYLIZER... \n ")

// // 	data, err := os.ReadFile("02-PARTICE.txt")

// // 	if err != nil {

// // 		fmt.Println("SHOWS ERROS : ", err)
// // 		return
// // 	}

// // 	text := string(data)

// // 	var words, charcters, letters, vowels, constants, uppercase, lowercase int
// // 	var totalCharsWithoutSpace, totalCharsWithSpace int

// // 	inword := false

// // 	for _, char := range text {

// // 		totalCharsWithSpace++

// // 		if !unicode.IsSpace(char) {
// // 			totalCharsWithoutSpace++
// // 		}

// // 		charcters++

// // 		if char == ' ' || char == '\n' || char == '\t' {

// // 			inword = false

// // 		} else {

// // 			if !inword {

// // 				words++

// // 				inword = true
// // 			}
// // 		}

// // 		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {

// // 			letters++ // lettters , withoutspace

// // 			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {

// // 				vowels++

// // 			} else {

// // 				constants++
// // 			}
// // 		}

// // 		if unicode.IsLetter(char) {

// // 			if unicode.IsUpper(char) {

// // 				uppercase++

// // 			} else {

// // 				lowercase++
// // 			}
// // 		}

// // 	}

// // 	fmt.Printf("01 -Total characters: %d\n", charcters)
// // 	fmt.Printf("02 -Total characters: %d\n", len(text))

// // 	fmt.Printf("03 -Total words: %d\n", words)

// // 	fmt.Printf("04 -Total letters: %d\n", letters)

// // 	fmt.Printf("05 -VOWELS WORDS: %d\n", vowels)
// // 	fmt.Printf("06 -CONSATNTS WORDS: %d\n", constants)

// // 	fmt.Printf("07 -Upper-Case Letters: %d\n", uppercase)
// // 	fmt.Printf("08 -Lower-Case Letters: %d\n", lowercase)

// // 	fmt.Printf("09 -Total Chars Without-Space : %d\n", totalCharsWithoutSpace)
// // 	fmt.Printf("10 -Total Chars WithSpace : %d\n", totalCharsWithSpace)

// // }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	fmt.Printf("\n \n TEXT ANYLIZER... \n ")

// 	data, err := os.ReadFile("02-PARTICE.txt")

// 	if err != nil {

// 		fmt.Println("SHOWS ERROS : ", err)
// 		return
// 	}

// 	text := string(data)

// 	var words, charcters, letters, vowels, constants, uppercase, lowercase int
// 	var totalCharsWithoutSpace, totalCharsWithSpace int
// 	var punctuations, integers, sentences, lines, spaces int

// 	inword := false
// 	insentence := false

// 	for _, char := range text {

// 		totalCharsWithSpace++
// 		charcters++

// 		if char == ' ' || char == '\t' {
// 			spaces++                    // Count spaces and tabs
// 		}

// 		if char != ' ' && char != '\n' && char != '\t' {
// 			totalCharsWithoutSpace++
// 		}

// 		// Word Count (Letters + Numbers)
// 		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
// 			if !inword {
// 				words++
// 				inword = true
// 			}
// 		} else {
// 			inword = false
// 		}

// 		// Letters, Vowels, Consonants, Upper Lower
// 		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {

// 			letters++

// 			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
// 			   char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
// 				vowels++
// 			} else {
// 				constants++
// 			}

// 			if char >= 'A' && char <= 'Z' {
// 				uppercase++
// 			} else {
// 				lowercase++
// 			}
// 		}

// 		// Punctuations Count
// 		if char == '.' || char == ',' || char == '!' || char == '?' || char == '-' ||
// 		   char == '"' || char == '\'' || char == ':' || char == ';' || char == '(' ||
// 		   char == ')' || char == '[' || char == ']' || char == '{' || char == '}' {
// 			punctuations++
// 		}

// 		// Integers (Numbers) Count
// 		if char >= '0' && char <= '9' {
// 			integers++
// 		}

// 		// Sentences Count
// 		if char == '.' || char == '!' || char == '?' {
// 			if insentence {
// 				sentences++
// 				insentence = false
// 			}
// 		} else if char != ' ' && char != '\n' && char != '\t' {
// 			insentence = true
// 		}

// 		// Lines Count
// 		if char == '\n' {
// 			lines++
// 		}
// 	}

// 	if charcters > 0 {
// 		lines++
// 	}

// 	fmt.Printf("01 -Total characters: %d\n", charcters)
// 	fmt.Printf("02 -Total characters: %d\n", len(text))


// 	fmt.Printf("03 -Total words: %d\n", words)

// 	fmt.Printf("04 -Total letters: %d\n", letters)

// 	fmt.Printf("05 -VOWELS WORDS: %d\n", vowels)
// 	fmt.Printf("06 -CONSATNTS WORDS: %d\n", constants)


// 	fmt.Printf("07 -Upper-Case Letters: %d\n", uppercase)
// 	fmt.Printf("08 -Lower-Case Letters: %d\n", lowercase)


// 	fmt.Printf("09 -Total Chars Without-Space : %d\n", totalCharsWithoutSpace)
// 	fmt.Printf("10 -Total Chars WithSpace : %d\n", totalCharsWithSpace)


// 	fmt.Printf("11 -Punctuations : %d\n", punctuations)


// 	fmt.Printf("12 -Integers (Numbers) : %d\n", integers)

// 	fmt.Printf("13 -Sentences : %d\n", sentences)


// 	fmt.Printf("14 -Total Lines : %d\n", lines)


// 	fmt.Printf("15 -Total Spaces : %d\n", spaces)

// }



