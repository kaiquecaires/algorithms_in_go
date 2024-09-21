package semordnilap

func Semordnilap(words []string) [][]string {
	storage := make(map[string]string)
	pairs := [][]string{}

	for _, word := range words {
		if _, ok := storage[word]; ok {
			pairs = append(pairs, []string{storage[word], word})
			continue
		}

		reversedWord := ""

		for i := len(word) - 1; i >= 0; i-- {
			reversedWord += string(word[i])
		}

		storage[reversedWord] = word
	}

	return pairs
}
