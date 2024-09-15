package generatedocument

func GenerateDocument(characters string, document string) bool {
	frequency := map[rune]int{}

	for _, c := range characters {
		frequency[c] += 1
	}

	for _, c := range document {
		f, ok := frequency[c]

		if !ok {
			return false
		}

		if f == 0 {
			return false
		}

		frequency[c]--
	}

	return true
}
