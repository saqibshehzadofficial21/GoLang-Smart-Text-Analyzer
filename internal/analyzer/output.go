package analyzer

import "fmt"

func PrintResult(r Result) {

	fmt.Println("\n TEXT ANALYZER \n")

	fmt.Println("Total Characters:", r.TotalChars)
	fmt.Println("Total Words:", r.TotalWords)
	fmt.Println("Total Letters:", r.TotalLetters)
	fmt.Println("Vowels:", r.Vowels)
	fmt.Println("Consonants:", r.Consonants)
	fmt.Println("Uppercase:", r.Uppercase)
	fmt.Println("Lowercase:", r.Lowercase)
	fmt.Println("Spaces:", r.Spaces)
	fmt.Println("Lines:", r.Lines)
	fmt.Println("Execution Time:", r.ExecutionTime.Milliseconds(), "ms")
}