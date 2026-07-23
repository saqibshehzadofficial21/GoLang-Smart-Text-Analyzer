package main

import (
	"fmt"
	"text-analyzer/internal/analyzer"
	"text-analyzer/internal/reader"
)

func main() {

	text, err := reader.ReadFile("text/big-file.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	var workers int
	fmt.Print("Enter number of goroutines: ")
	fmt.Scan(&workers)

	if workers <= 0 {
		workers = 1
	}

	result := analyzer.AnalyzeParallel(text, workers)

	analyzer.PrintResult(result)
}