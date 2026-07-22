package main

import (
	"fmt"
	"os"
	"text-analyzer/analyzer"

)

func main() {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	text := string(data)

	result := analyzer.Analyze(text)
	analyzer.PrintResult(result)
}