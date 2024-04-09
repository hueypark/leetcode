package findandreplacepattern

func findAndReplacePattern(words []string, pattern string) []string {
	var result []string
	for _, word := range words {
		if isMatch(word, pattern) {
			result = append(result, word)
		}
	}
	return result
}

func isMatch(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}
	wordToPattern := make(map[byte]byte)
	patternToWord := make(map[byte]byte)
	for i := 0; i < len(word); i++ {
		if _, ok := wordToPattern[word[i]]; !ok {
			wordToPattern[word[i]] = pattern[i]
		}
		if _, ok := patternToWord[pattern[i]]; !ok {
			patternToWord[pattern[i]] = word[i]
		}
		if wordToPattern[word[i]] != pattern[i] || patternToWord[pattern[i]] != word[i] {
			return false
		}
	}
	return true
}
