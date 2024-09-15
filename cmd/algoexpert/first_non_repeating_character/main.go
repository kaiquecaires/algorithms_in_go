package firstnonrepeatingcharacter

func FirstNonRepeatingCharacter(str string) int {
	frequency := map[rune]int{}

	for _, c := range str {
		frequency[c]++
	}

	for i, c := range str {
		if frequency[c] == 1 {
			return i
		}
	}

	return -1
}
