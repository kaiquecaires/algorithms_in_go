package algoexpert

func IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIdx := 0
	arrayIdx := 0

	for sequenceIdx < len(sequence) && arrayIdx < len(array) {
		if sequence[sequenceIdx] == array[arrayIdx] {
			sequenceIdx++
		}
		arrayIdx++
	}

	return sequenceIdx == len(sequence)
}
