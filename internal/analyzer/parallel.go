package analyzer

import (
	"sync"
	"text-analyzer/internal/splitter"
	"time"
)

func AnalyzeParallel(text string, workers int) Result {

	start := time.Now() 

	if workers <= 0 {
		workers = 1
	}

	chunkSize := len(text) / workers
	chunks := splitter.SplitText(text, chunkSize)

	var wg sync.WaitGroup
	resultChan := make(chan Result, len(chunks))

	for _, chunk := range chunks {

		wg.Add(1)

		go func(c string) {
			defer wg.Done()

			resultChan <- Analyze(c)

		}(chunk)
	}

	wg.Wait()
	close(resultChan)

	var final Result

	for r := range resultChan {
		final.TotalChars += r.TotalChars
		final.TotalWords += r.TotalWords
		final.TotalLetters += r.TotalLetters
		final.Vowels += r.Vowels
		final.Consonants += r.Consonants
		final.Uppercase += r.Uppercase
		final.Lowercase += r.Lowercase
		final.CharsWithSpace += r.CharsWithSpace
		final.CharsWithoutSpace += r.CharsWithoutSpace
		final.Punctuations += r.Punctuations
		final.Integers += r.Integers
		final.Sentences += r.Sentences
		final.Lines += r.Lines
		final.Spaces += r.Spaces
	}


	final.ExecutionTime = time.Since(start) 

	return final
}