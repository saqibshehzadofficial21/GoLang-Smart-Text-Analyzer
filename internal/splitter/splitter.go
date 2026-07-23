package splitter

func SplitText(text string, chunkSize int) []string {

	if chunkSize <= 0 {
		return []string{text}
	}

	var chunks []string

	for i := 0; i < len(text); i += chunkSize {
		end := i + chunkSize

		if end > len(text) {
			end = len(text)
		}

		chunks = append(chunks, text[i:end])
	}

	return chunks
}