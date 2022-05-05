package hfgolib

func ForUnicode(word string) map[string]rune {
	wordMap := make(map[string]rune)
	for _, r := range word {
		wordMap[string(r)] = r
	}

	return wordMap
}
