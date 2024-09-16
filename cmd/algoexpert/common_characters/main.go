package commoncharacters

func CommonCharacters(strings []string) []string {
	characters := map[rune]int{}

	for _, s := range strings {
		uniqueChars := map[rune]bool{}
		for _, c := range s {
			uniqueChars[c] = true
		}
		for k := range uniqueChars {
			characters[k]++
		}
	}

	common := []string{}

	for k, v := range characters {
		if v == len(strings) {
			common = append(common, string(k))
		}
	}

	return common
}
