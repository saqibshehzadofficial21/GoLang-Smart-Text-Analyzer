package analyzer

import "time"

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